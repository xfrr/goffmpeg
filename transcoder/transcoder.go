package transcoder

import (
    //"fmt"
    //"os"
	"errors"
	"os"
	"goffmpeg/models"
)

type Transcoder struct {
	Process               *os.Process
	InputPath             string
	OutputPath            string
	MediaFile			  *models.Mediafile
}

func New(inputPath *string) (*Transcoder, error) {
    transcoding := new(Transcoder)

	if inputPath == nil {
    	return nil, errors.New("error: transcoder.Initialize -> inputPath missing")
	}

    transcoding.InputPath = *inputPath
    transcoding.MediaFile = new(models.Mediafile)

    return transcoding, nil

}

func (t *Transcoder) SetBitRate(v *string) string {
	t.MediaFile.VideoBitRate = *v
	return t.MediaFile.VideoBitRate
}
