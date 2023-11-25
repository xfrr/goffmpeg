package ffmpeg

import (
	"context"
	"fmt"
	"io"
	"os/exec"

	"github.com/xfrr/goffmpeg/v2/ffmpeg/progress"
)

type Command struct {
	cmd            *exec.Cmd
	binPath        string
	args           []string
	inputReader    io.Reader
	outputWriter   io.Writer
	progressReader progress.Reader
	err            chan error
}

func (c Command) Args() []string {
	return c.args
}

// WithBinPath sets the path to the ffmpeg binary.
func (c *Command) WithBinPath(binPath string) *Command {
	c.binPath = binPath
	return c
}

func NewCommand() *Command {
	return &Command{
		binPath:        "ffmpeg",
		cmd:            &exec.Cmd{},
		args:           DefaultArgs(),
		inputReader:    nil,
		outputWriter:   nil,
		progressReader: progress.NewReader(),
		err:            make(chan error, 1),
	}
}

func DefaultArgs() []string {
	return []string{
		"-nostats",
		"-loglevel", "error",
		"-y",
	}
}

func (c *Command) Run(ctx context.Context) error {
	if len(c.args) == 0 {
		return fmt.Errorf("command not initialized")
	}

	c.cmd = exec.CommandContext(ctx, c.binPath, c.args...)
	if c.inputReader != nil {
		c.cmd.Stdin = c.inputReader
	}

	if c.outputWriter != nil {
		c.cmd.Stdout = c.outputWriter
	}

	return c.cmd.Run()
}

func (c *Command) Start(ctx context.Context) (<-chan progress.Progress, error) {
	if len(c.args) == 0 {
		return nil, fmt.Errorf("command not initialized")
	}

	c.args = append([]string{"-progress", "pipe:2"}, c.args...)
	c.cmd = exec.CommandContext(ctx, c.binPath, c.args...)

	if c.inputReader != nil {
		c.cmd.Stdin = c.inputReader
	}

	if c.outputWriter != nil {
		c.cmd.Stdout = c.outputWriter
	}

	stderr, err := c.cmd.StderrPipe()
	if err != nil {
		return nil, err
	}

	progressCh := make(chan progress.Progress, 10)

	go func() {
		defer close(progressCh)
		c.err <- c.progressReader.Read(ctx, stderr, progressCh)
	}()

	err = c.cmd.Start()
	if err != nil {
		return nil, err
	}

	return progressCh, nil
}

func (c *Command) Wait() error {
	defer close(c.err)
	err := <-c.err
	if err != nil {
		return err
	}
	return c.cmd.Wait()
}
