package transcoder

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/Autherz/goffmpeg/ffmpeg"
	"github.com/Autherz/goffmpeg/models"
	"github.com/Autherz/goffmpeg/utils"
)

// Transcoder Main struct
type Transcoder struct {
	stdErrPipe    io.ReadCloser
	stdStdinPipe  io.WriteCloser
	process       *exec.Cmd
	mediafile     *models.Mediafile
	configuration ffmpeg.Configuration
}

// SetProcessStderrPipe Set the STDERR pipe
func (t *Transcoder) SetProcessStderrPipe(v io.ReadCloser) {
	t.stdErrPipe = v
}

// SetProcessStdinPipe Set the STDIN pipe
func (t *Transcoder) SetProcessStdinPipe(v io.WriteCloser) {
	t.stdStdinPipe = v
}

// SetProcess Set the transcoding process
func (t *Transcoder) SetProcess(cmd *exec.Cmd) {
	t.process = cmd
}

// SetMediaFile Set the media file
func (t *Transcoder) SetMediaFile(v *models.Mediafile) {
	t.mediafile = v
}

// SetConfiguration Set the transcoding configuration
func (t *Transcoder) SetConfiguration(v ffmpeg.Configuration) {
	t.configuration = v
}

// Process Get transcoding process
func (t Transcoder) Process() *exec.Cmd {
	return t.process
}

// MediaFile Get the ttranscoding media file.
func (t Transcoder) MediaFile() *models.Mediafile {
	return t.mediafile
}

// FFmpegExec Get FFmpeg Bin path
func (t Transcoder) FFmpegExec() string {
	return t.configuration.FfmpegBin
}

// FFprobeExec Get FFprobe Bin path
func (t Transcoder) FFprobeExec() string {
	return t.configuration.FfprobeBin
}

// GetCommand Build and get command
func (t Transcoder) GetCommand() []string {
	media := t.mediafile
	rcommand := append([]string{"-y"}, media.ToStrCommand()...)
	return rcommand
}

// Initialize Init the transcoding process
func (t *Transcoder) Initialize(inputPath string, outputPath string) error {
	var err error
	var out bytes.Buffer
	var Metadata models.Metadata

	cfg := t.configuration

	if len(cfg.FfmpegBin) == 0 || len(cfg.FfprobeBin) == 0 {
		cfg, err = ffmpeg.Configure()
		if err != nil {
			return err
		}
	}

	if inputPath == "" {
		return errors.New("error on transcoder.Initialize: inputPath missing")
	}

	command := []string{"-i", inputPath, "-print_format", "json", "-show_format", "-show_streams", "-show_error"}

	cmd := exec.Command(cfg.FfprobeBin, command...)
	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error executing (%s) | error: %s", command, err)
	}

	if err = json.Unmarshal([]byte(out.String()), &Metadata); err != nil {
		return err
	}

	// Set new Mediafile
	MediaFile := new(models.Mediafile)
	MediaFile.SetMetadata(Metadata)
	MediaFile.SetInputPath(inputPath)
	MediaFile.SetOutputPath(outputPath)

	// Set transcoder configuration
	t.SetMediaFile(MediaFile)
	t.SetConfiguration(cfg)

	return nil

}

// Run Starts the transcoding process
func (t *Transcoder) Run(progress bool) <-chan error {
	done := make(chan error)
	command := t.GetCommand()

	if !progress {
		command = append([]string{"-nostats", "-loglevel", "0"}, command...)
	}

	proc := exec.Command(t.configuration.FfmpegBin, command...)
	if progress {
		errStream, err := proc.StderrPipe()
		if err != nil {
			fmt.Println("Progress not available: " + err.Error())
		} else {
			t.stdErrPipe = errStream
		}
	}

	stdin, err := proc.StdinPipe()
	if nil != err {
		fmt.Println("Stdin not available: " + err.Error())
	}

	t.stdStdinPipe = stdin

	out := &bytes.Buffer{}
	if progress {
		proc.Stdout = out
	}

	err = proc.Start()

	t.SetProcess(proc)
	go func(err error, out *bytes.Buffer) {
		if err != nil {
			done <- fmt.Errorf("Failed Start FFMPEG (%s) with %s, message %s", command, err, out.String())
			close(done)
			return
		}
		err = proc.Wait()
		if err != nil {
			err = fmt.Errorf("Failed Finish FFMPEG (%s) with %s message %s", command, err, out.String())
		}
		done <- err
		close(done)
	}(err, out)

	return done
}

// Stop Ends the transcoding process
func (t *Transcoder) Stop() error {
	if t.process != nil {

		stdin := t.stdStdinPipe
		if stdin != nil {
			stdin.Write([]byte("q\n"))
		}
	}
	return nil
}

// Output Returns the transcoding progress channel
func (t Transcoder) Output() <-chan models.Progress {
	out := make(chan models.Progress)

	go func() {
		defer close(out)
		if t.stdErrPipe == nil {
			out <- models.Progress{}
			return
		}

		defer t.stdErrPipe.Close()

		scanner := bufio.NewScanner(t.stdErrPipe)

		split := func(data []byte, atEOF bool) (advance int, token []byte, spliterror error) {
			if atEOF && len(data) == 0 {
				return 0, nil, nil
			}
			if i := bytes.IndexByte(data, '\n'); i >= 0 {
				// We have a full newline-terminated line.
				return i + 1, data[0:i], nil
			}
			if i := bytes.IndexByte(data, '\r'); i >= 0 {
				// We have a cr terminated line
				return i + 1, data[0:i], nil
			}
			if atEOF {
				return len(data), data, nil
			}

			return 0, nil, nil
		}

		scanner.Split(split)
		buf := make([]byte, 2)
		scanner.Buffer(buf, bufio.MaxScanTokenSize)

		for scanner.Scan() {
			Progress := new(models.Progress)
			line := scanner.Text()
			if strings.Contains(line, "frame=") && strings.Contains(line, "time=") && strings.Contains(line, "bitrate=") {
				var re = regexp.MustCompile(`=\s+`)
				st := re.ReplaceAllString(line, `=`)

				f := strings.Fields(st)
				var framesProcessed string
				var currentTime string
				var currentBitrate string
				var currentSpeed string

				for j := 0; j < len(f); j++ {
					field := f[j]
					fieldSplit := strings.Split(field, "=")

					if len(fieldSplit) > 1 {
						fieldname := strings.Split(field, "=")[0]
						fieldvalue := strings.Split(field, "=")[1]

						if fieldname == "frame" {
							framesProcessed = fieldvalue
						}

						if fieldname == "time" {
							currentTime = fieldvalue
						}

						if fieldname == "bitrate" {
							currentBitrate = fieldvalue
						}
						if fieldname == "speed" {
							currentSpeed = fieldvalue
						}
					}
				}

				timesec := utils.DurToSec(currentTime)
				dursec, _ := strconv.ParseFloat(t.MediaFile().Metadata().Format.Duration, 64)
				//live stream check
				if dursec != 0 {
					// Progress calculation
					progress := (timesec * 100) / dursec
					Progress.Progress = progress
				}
				Progress.CurrentBitrate = currentBitrate
				Progress.FramesProcessed = framesProcessed
				Progress.CurrentTime = currentTime
				Progress.Speed = currentSpeed
				out <- *Progress
			}
		}
	}()

	return out
}
