package ffprobe

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"

	"github.com/xfrr/goffmpeg/v2/pkg/media"
)

func (c *Command) run(ctx context.Context, cmd *exec.Cmd, args ...string) (media.File, error) {
	if c.inputReader != nil {
		cmd.Stdin = c.inputReader
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return media.File{}, err
	}

	errb := new(bytes.Buffer)
	go io.Copy(errb, stderr)

	outb := new(bytes.Buffer)
	cmd.Stdout = outb

	err = cmd.Run()
	if err != nil {
		if errb.Len() > 0 {
			return media.File{}, fmt.Errorf(errb.String())
		}
		return media.File{}, err
	}

	// make a copy of the output buffer
	// because it will be modified by json.Unmarshal
	// and we want to return it as is
	outbCopy := make([]byte, outb.Len())
	copy(outbCopy, outb.Bytes())

	if c.outputWriter != nil {
		n, err := c.outputWriter.Write(outbCopy)
		if err != nil {
			return media.File{}, err
		}
		if n != len(outbCopy) {
			return media.File{}, fmt.Errorf("failed to write all output to writer")
		}
	}

	mediafile := new(media.File)
	err = json.Unmarshal(outb.Bytes(), mediafile)
	if err != nil {
		return media.File{}, err
	}

	return *mediafile, nil
}
