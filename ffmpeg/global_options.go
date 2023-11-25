package ffmpeg

import (
	"fmt"
	"strconv"
)

// WithReport enables the ffmpeg report.
func (c *Command) WithReport() *Command {
	c.args = append(c.args, "-report")
	return c
}

// WithMaxAlloc sets the maximum size of a single allocated block.
func (c *Command) WithMaxAlloc(bytes int) *Command {
	c.args = append(c.args, "-max_alloc", strconv.Itoa(bytes))
	return c
}

// WithFilterThreads sets the number of non-complex filter threads.
func (c *Command) WithFilterThreads(threads int) *Command {
	c.args = append(c.args, "-filter_threads", strconv.Itoa(threads))
	return c
}

// WithTimeLimit sets the max runtime in seconds.
func (c *Command) WithTimeLimit(seconds int) *Command {
	c.args = append(c.args, "-timelimit", strconv.Itoa(seconds))
	return c
}

// WithAudioDriftThreshold sets the audio drift threshold.
func (c *Command) WithAudioDriftThreshold(threshold int) *Command {
	c.args = append(c.args, "-adrift_threshold", strconv.Itoa(threshold))
	return c
}

// WithCopyTS copies timestamps.
func (c *Command) WithCopyTS() *Command {
	c.args = append(c.args, "-copyts")
	return c
}

// WithCopyTB copies input stream time base when stream copying.
func (c *Command) WithCopyTB(mode string) *Command {
	c.args = append(c.args, "-copytb", mode)
	return c
}

// WithDTSDeltaThreshold sets the timestamp discontinuity delta threshold.
func (c *Command) WithDTSDeltaThreshold(threshold int) *Command {
	c.args = append(c.args, "-dts_delta_threshold", strconv.Itoa(threshold))
	return c
}

// WithDTSErrorThreshold sets the timestamp error delta threshold.
func (c *Command) WithDTSErrorThreshold(threshold int) *Command {
	c.args = append(c.args, "-dts_error_threshold", strconv.Itoa(threshold))
	return c
}

// WithXError exits on error.
func (c *Command) WithXError() *Command {
	c.args = append(c.args, "-xerror")
	return c
}

// WithDebugTS prints timestamp debugging info.
func (c *Command) WithDebugTS() *Command {
	c.args = append(c.args, "-debug_ts")
	return c
}

// WithVStatsFile dumps video coding statistics to file.
func (c *Command) WithVStatsFile(file string) *Command {
	c.args = append(c.args, "-vstats_file", file)
	return c
}

// WithISync assigns an input as a sync source.
func (c *Command) WithISync() *Command {
	c.args = append(c.args, "-isync")
	return c
}

// WithMaxBitrate sets the max bitrate in kbit/s.
func (c *Command) WithMaxBitrate(bitrate int) *Command {
	c.args = append(c.args, "-maxrate", fmt.Sprintf("%dk", bitrate))
	return c
}
