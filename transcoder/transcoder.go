package transcoder

import (
    //"fmt"
    //"os"
	"errors"
	"os"
	"goffmpeg/models"
    "os/exec"
	"fmt"
	"goffmpeg/ffmpeg"
	"bytes"
)

type Transcoder struct {
	Process               *os.Process
	InputPath             string
	OutputPath            string
	MediaFile			  *models.Mediafile
}

func New(inputPath *string, configuration *ffmpeg.Configuration) (*Transcoder, error) {
    transcoding := new(Transcoder)

	if inputPath == nil {
    	return nil, errors.New("error: transcoder.Initialize -> inputPath missing")
	}

	_, err := os.Stat(*inputPath)
	if os.IsNotExist(err) {
		return nil, errors.New("error: transcoder.Initialize -> input file not found")
	}

	// Set input path
    transcoding.InputPath = *inputPath

    // TODO: Get file metadata from ffprobe and set MediaFile
	command := fmt.Sprintf("%s -i %s -print_format json -show_format -show_streams -show_error", configuration.FfprobeBin, *inputPath)

	cmd := exec.Command("/bin/sh", "-c", command)

	var out bytes.Buffer

	cmd.Stdout = &out

	cmdErr := cmd.Start()

	if cmdErr != nil {
		return nil, cmdErr
	}

	_, errProc := cmd.Process.Wait()
	if errProc != nil {
		return nil, errProc
	}

	stdout := out.String()

	fmt.Println(stdout)

    transcoding.MediaFile = new(models.Mediafile)

    return transcoding, nil

}

func (t *Transcoder) SetBitRate(v *string) string {
	t.MediaFile.VideoBitRate = *v
	return t.MediaFile.VideoBitRate
}
