package ffprobe

import (
	"context"
	"io"
	"os/exec"

	"github.com/xfrr/goffmpeg/v2/pkg/media"
)

type Command struct {
	cmd          *exec.Cmd
	binPath      string
	args         []string
	inputReader  io.Reader
	outputWriter io.Writer
}

func (c *Command) Args() []string {
	return c.args
}

func NewCommand() *Command {
	return &Command{
		binPath:      "ffprobe",
		cmd:          &exec.Cmd{},
		args:         defaultArgs(),
		inputReader:  nil,
		outputWriter: nil,
	}
}

func defaultArgs() []string {
	return []string{
		"-loglevel", "error",
		"-print_format", "json",
		"-show_format",
		"-show_streams",
		"-show_error",
	}
}

// Run executes the ffprobe command and returns the media file
func (c *Command) Run(ctx context.Context) (media.File, error) {
	cmd := exec.CommandContext(ctx, "ffprobe", c.args...)
	return c.run(ctx, cmd)
}
