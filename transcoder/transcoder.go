package transcoder

import (
	"errors"
	"os"
	"goffmpeg/models"
    "os/exec"
	"fmt"
	"goffmpeg/ffmpeg"
	"goffmpeg/utils"
	"bytes"
	"encoding/json"
	"bufio"
	"strings"
	"regexp"
	"strconv"
)

type Transcoder struct {
	process               *exec.Cmd
	inputPath             string
	outputPath            string
	mediafile			  *models.Mediafile
	configuration         ffmpeg.Configuration
}

func (t *Transcoder) SetProccess(v *exec.Cmd) {
	t.process = v
}

func (t *Transcoder) SetInputPath(v string) {
	t.inputPath = v
}

func (t *Transcoder) SetOutputPath(v string) {
	t.outputPath = v
}

func (t *Transcoder) SetMediaFile(v *models.Mediafile) {
	t.mediafile = v
}

func (t *Transcoder) SetConfiguration(v ffmpeg.Configuration) {
	t.configuration = v
}

/*** GETTERS ***/

func (t Transcoder) Process() *exec.Cmd {
	return t.process
}

func (t Transcoder) InputPath() string {
	return t.inputPath
}

func (t Transcoder) OutputPath() string {
	return t.outputPath
}

func (t Transcoder) MediaFile() *models.Mediafile {
	return t.mediafile
}

func (t Transcoder) FFmpegExec() string {
	return t.configuration.FfmpegBin
}

func (t Transcoder) FFprobeExec() string {
	return t.configuration.FfprobeBin
}

func (t Transcoder) GetCommand() string {
	var rcommand string

	rcommand = fmt.Sprintf("%s -y -i %s ", t.configuration.FfmpegBin, t.inputPath)

	media := t.mediafile

	rcommand += media.ToStrCommand()

	rcommand += " " + t.outputPath

	return rcommand

}

/*** FUNCTIONS ***/

func (t *Transcoder) Initialize(inputPath string, outputPath string) (error) {

	configuration, err := ffmpeg.Configure()
	if err != nil {
		fmt.Println(err)
		return err
	}

	if inputPath == "" {
    	return errors.New("error: transcoder.Initialize -> inputPath missing")
	}

	_, err = os.Stat(inputPath)
	if os.IsNotExist(err) {
		return errors.New("error: transcoder.Initialize -> input file not found")
	}

	command := fmt.Sprintf("%s -i %s -print_format json -show_format -show_streams -show_error", configuration.FfprobeBin, inputPath)

	cmd := exec.Command(configuration.ExecCmd, configuration.ExecArgs, command)

	fmt.Println("FFprobe command: " + command)

	var out bytes.Buffer

	cmd.Stdout = &out

	err = cmd.Start()

	if err != nil {
		return err
	}

	_, err = cmd.Process.Wait()
	if err != nil {
		return err
	}

	var Metadata models.Metadata

	if err = json.Unmarshal([]byte(out.String()), &Metadata); err != nil {
		return err
	}

	// Set new Mediafile
    MediaFile := new(models.Mediafile)
    MediaFile.SetMetadata(Metadata)

    // Set transcoder configuration
    t.SetInputPath(inputPath)
    t.SetOutputPath(outputPath)
    t.SetMediaFile(MediaFile)
	t.SetConfiguration(configuration)

    return nil

}

func (t *Transcoder) Run() (<-chan bool, error) {
	done := make(chan bool)
	var err error

	command := t.GetCommand()

	fmt.Println("FFmpeg command: " + command)

	proc := exec.Command(t.configuration.ExecCmd, t.configuration.ExecArgs, command)

	t.SetProccess(proc)

	go func() {
		err := proc.Start()
		if err != nil {
			return
		}

		proc.Wait()

		done <- true
		close(done)
	}()

	return done, err

}

func (t Transcoder) Output() (<-chan models.Progress, error) {
	out := make(chan models.Progress)
	var err error

		go func() {
			defer close(out)

			stderr, stderror := t.Process().StderrPipe()
			if err != nil {
				err = stderror
				return
			}

			scanner := bufio.NewScanner(stderr)
			filetype := utils.CheckFileType(t.MediaFile().Metadata().Streams)

			split := func(data []byte, atEOF bool) (advance int, token []byte, spliterror error) {

				if atEOF && len(data) == 0 {
					return 0, nil, nil
				}

				Iframe := strings.Index(string(data), "frame=")

				if filetype == "video" {
					if Iframe > 0 {
						return Iframe + 1, data[Iframe:], nil
					}
				} else {
					if i := bytes.IndexByte(data, '\n'); i >= 0 {
						// We have a full newline-terminated line.
						return i + 1, data[0:i], nil
					}
				}

				if atEOF {
					return len(data), data, nil
				}

				return 0, nil, nil
			}

			scanner.Split(split)
			buf := make([]byte, 2)
			scanner.Buffer(buf, bufio.MaxScanTokenSize)

			var lastProgress float64
			for scanner.Scan() {
				Progress := new(models.Progress)
				line := scanner.Text()

				if strings.Contains(line, "time=") && strings.Contains(line, "bitrate=") {
					var re= regexp.MustCompile(`=\s+`)
					st := re.ReplaceAllString(line, `=`)

					f := strings.Fields(st)

					var framesProcessed string
					var currentTime string
					var currentBitrate string

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
						}
					}

					timesec := utils.DurToSec(currentTime)
					dursec, _ := strconv.ParseFloat(t.MediaFile().Metadata().Format.Duration, 64)
					// Progress calculation
					progress := (timesec * 100) / dursec

					Progress.Progress = progress
					Progress.CurrentBitrate = currentBitrate
					Progress.FramesProcessed = framesProcessed
					Progress.CurrentTime = currentTime

					if progress != lastProgress {
						lastProgress = progress
						out <- *Progress
					}
				}
			}

			if scerror := scanner.Err(); err != nil {
				err = scerror
			}
		}()

	return out, err
}
