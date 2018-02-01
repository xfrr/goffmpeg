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
	"log"
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

	fmt.Println(rcommand)

	return rcommand

}

/*** FUNCTIONS ***/

func (t *Transcoder) Initialize(inputPath string, outputPath string, configuration *ffmpeg.Configuration) (error) {

	if inputPath == "" {
    	return errors.New("error: transcoder.Initialize -> inputPath missing")
	}

	_, err := os.Stat(inputPath)
	if os.IsNotExist(err) {
		return errors.New("error: transcoder.Initialize -> input file not found")
	}

	command := fmt.Sprintf("%s -i %s -print_format json -show_format -show_streams -show_error", configuration.FfprobeBin, inputPath)

	cmd := exec.Command("/bin/sh", "-c", command)

	var out bytes.Buffer

	cmd.Stdout = &out

	cmdErr := cmd.Start()

	if cmdErr != nil {
		return cmdErr
	}

	_, errProc := cmd.Process.Wait()
	if errProc != nil {
		return errProc
	}

	var Metadata models.Metadata

	if err := json.Unmarshal([]byte(out.String()), &Metadata); err != nil {
		return err
	}

	// Set new Mediafile
    MediaFile := new(models.Mediafile)
    MediaFile.SetMetadata(Metadata)

    // Set transcoder configuration
    t.SetInputPath(inputPath)
    t.SetOutputPath(outputPath)
    t.SetMediaFile(MediaFile)
    t.SetConfiguration(*configuration)

    return nil

}

func (t *Transcoder) Run() (<-chan bool)  {
	done := make(chan bool)
	command := t.GetCommand()

	proc := exec.Command("/bin/sh", "-c", command)

	t.SetProccess(proc)

	go func() {
		perror := proc.Start()
		if perror != nil {
			log.Fatal(perror)
			return
		}

		proc.Wait()

		done <- true
	}()

	return done

}

// TODO: ONLY WORKS FOR VIDEO FILES
func (t Transcoder) Output() (chan models.Progress) {
	out := make(chan models.Progress)

		go func() {

			stderr, serr := t.Process().StderrPipe()
			if serr != nil {
				log.Fatal(serr)
				return
			}

			scanner := bufio.NewScanner(stderr)

			split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {

				if err != nil {
					log.Fatal(err)
					return 0, nil, nil
				}

				if atEOF && len(data) == 0 {
					return 0, nil, nil
				}

				fr := strings.Index(string(data), "frame=")

				if fr > 0 {
					return fr + 1, data[fr:], nil
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
			var lastFrames string
			for scanner.Scan() {
				Progress := new(models.Progress)
				line := scanner.Text()

				if strings.Contains(line, "time=") && strings.Contains(line, "bitrate=") {
					var re= regexp.MustCompile(`=\s+`)
					st := re.ReplaceAllString(line, `=`)

					f := strings.Fields(st)

					// Frames processed
					framesProcessed := strings.Split(f[0], "=")[1]

					// Current time processed
					time := strings.Split(f[4], "=")[1]
					timesec := utils.DurToSec(time)
					dursec, _ := strconv.ParseFloat(t.MediaFile().Metadata().Format.Duration, 64)
					// Progress calculation
					progress := (timesec * 100) / dursec

					// Current bitrate
					currentBitrate := strings.Split(f[5], "=")[1]

					Progress.Progress = progress
					Progress.CurrentBitrate = currentBitrate
					Progress.FramesProcessed = framesProcessed
					Progress.CurrentTime = time

					if progress != lastProgress && framesProcessed != lastFrames{
						lastProgress = progress
						lastFrames = framesProcessed
						out <- *Progress
					}
				}
			}

			defer close(out)

			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
		}()

	return out
}
