package ffprobe

import "io"

// WithBinPath sets the path to the ffmpeg binary.
func (c *Command) WithBinPath(binPath string) *Command {
	c.binPath = binPath
	return c
}

func (c *Command) WithInputPath(inputPath string) *Command {
	c.args = append(c.args, "-i", inputPath)
	return c
}

func (c *Command) WithInputReader(inputReader io.Reader) *Command {
	c.args = append(c.args, "-i", "pipe:0")
	c.inputReader = inputReader
	return c
}

func (c *Command) WithOutputWriter(outputWriter io.Writer) *Command {
	c.args = append(c.args, "-o", "pipe:1")
	c.outputWriter = outputWriter
	return c
}
