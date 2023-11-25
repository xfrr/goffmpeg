package ffmpeg

import (
	"io"
	"strconv"
)

// WithInputPath sets the input path to the given path.
func (c *Command) WithInputPath(path string) *Command {
	c.args = append(c.args, "-i", path)
	return c
}

// WithInputPipe sets the input pipe reader (stdin).
func (c *Command) WithInputPipe(r io.Reader) *Command {
	c.args = append(c.args, "-i", "pipe:0")
	c.inputReader = r
	return c
}

// WithInputFormat sets the input format to the given format.
func (c *Command) WithOutputPath(path string) *Command {
	c.args = append(c.args, path)
	return c
}

// WithOutputPipe sets the output pipe writer to the given writer (stdout)
func (c *Command) WithOutputPipe(w io.Writer) *Command {
	c.args = append(c.args, "pipe:1")
	c.outputWriter = w
	return c
}

// WithOutputFormat sets the output format to the given format.
func (c *Command) WithOutputFormat(format string) *Command {
	c.args = append(c.args, "-f", format)
	return c
}

// WithCodec selects an encoder (when used before an output file) or a decoder (when used before an input file) for one or more streams.
// Codec is the name of a decoder/encoder or a special value copy (output only) to indicate that the stream is not to be re-encoded.
func (c *Command) WithCodec(codec string) *Command {
	c.args = append(c.args, "-c", codec)
	return c
}

// WithPreset sets the preset for matching stream(s).
func (c *Command) WithPreset(preset string) *Command {
	c.args = append(c.args, "-preset", preset)
	return c
}

// WithDuration when used as an input option (before -i), limit the duration of data read from the input file.
// When used as an output option (before an output url), stop writing the output after its duration reaches duration.
// duration must be a time duration specification, see (ffmpeg-utils)the Time duration section in the ffmpeg-utils(1) manual.
// -to and -t are mutually exclusive and -t has priority.
func (c *Command) WithDuration(duration string) *Command {
	c.args = append(c.args, "-t", duration)
	return c
}

// WithStopTime stop writing the output or reading the input at position. position must be a time duration specification.
func (c *Command) WithStopTime(time string) *Command {
	c.args = append(c.args, "-to", time)
	return c
}

// WithOutputLimitFileSize sets the file size limit, expressed in bytes.
// No further chunk of bytes is written after the limit is exceeded.
// The size of the output file is slightly more than the requested file size.
func (c *Command) WithLimitFileSize(size int) *Command {
	c.args = append(c.args, "-fs", strconv.Itoa(size))
	return c
}

// WithStartTimeOffset when used as an input option (before -i), seeks in this input file to position.
// Note that in most formats it is not possible to seek exactly, so ffmpeg will seek to the closest seek point before position.
// When transcoding and -accurate_seek is enabled (the default), this extra segment between the seek point and position will be decoded and discarded.
// When doing stream copy or when -noaccurate_seek is used, it will be preserved.
// When used as an output option (before an output url), decodes but discards input until the timestamps reach position.
func (c *Command) WithStartTimeOffset(offset string) *Command {
	c.args = append(c.args, "-ss", offset)
	return c
}

// WithTimestamp sets the recording timestamp in the container.
func (c *Command) WithTimestamp(timestamp string) *Command {
	c.args = append(c.args, "-timestamp", timestamp)
	return c
}

// WithMetadata sets metadata information to the output file.
func (c *Command) WithMetadata(metadata string) *Command {
	c.args = append(c.args, "-metadata", metadata)
	return c
}

// WithTarget sets the target file type.
func (c *Command) WithTarget(typ string) *Command {
	c.args = append(c.args, "-target", typ)
	return c
}

// WithAudioPad adds "padding" audio frames.
func (c *Command) WithAudioPad() *Command {
	c.args = append(c.args, "-apad")
	return c
}

// WithFrames sets the number of frames to record.
func (c *Command) WithFrames(number int) *Command {
	c.args = append(c.args, "-frames", strconv.Itoa(number))
	return c
}

// WithFilter sets stream filtergraph.
func (c *Command) WithFilter(filter string) *Command {
	c.args = append(c.args, "-filter", filter)
	return c
}

// WithFilterScript sets read stream filtergraph description from a file.
func (c *Command) WithFilterScript(filename string) *Command {
	c.args = append(c.args, "-filter_script", filename)
	return c
}

// WithReinitFilter reinit filtergraph on input parameter changes.
func (c *Command) WithReinitFilter() *Command {
	c.args = append(c.args, "-reinit_filter")
	return c
}

// WithInputTsOffset set the input ts offset.
func (c *Command) WithInputTsOffset(offset string) *Command {
	c.args = append(c.args, "-itsoffset", offset)
	return c
}

// WithInputTsScale set the input ts scale.
func (c *Command) WithInputTsScale(scale string) *Command {
	c.args = append(c.args, "-itsscale", scale)
	return c
}

// WithDataFrames set the number of data frames to record.
func (c *Command) WithDataFrames(number int) *Command {
	c.args = append(c.args, "-dframes", strconv.Itoa(number))
	return c
}

// WithReadInputAtNativeFrameRate reads input at native frame rate.
func (c *Command) WithReadInputAtNativeFrameRate() *Command {
	c.args = append(c.args, "-re")
	return c
}

// WithCopyInitialNonKeyframes copies non-keyframes initially.
func (c *Command) WithCopyInitialNonKeyframes() *Command {
	c.args = append(c.args, "-copyinkf")
	return c
}

// WithCopyOrDiscardFramesBeforeStartTime copies or discards frames before start time.
func (c *Command) WithCopyOrDiscardFramesBeforeStartTime() *Command {
	c.args = append(c.args, "-copypriorss")
	return c
}

// WithForceCodecTag forces codec tag/fourcc.
func (c *Command) WithForceCodecTag(tag string) *Command {
	c.args = append(c.args, "-tag", tag)
	return c
}

// WithFixedQualityScale sets fixed quality scale (VBR).
func (c *Command) WithFixedQualityScale(q int) *Command {
	c.args = append(c.args, "-q", strconv.Itoa(q))
	return c
}

// WithAttachment adds an attachment to the output file.
func (c *Command) WithAttachment(filename string) *Command {
	c.args = append(c.args, "-attach", filename)
	return c
}

// WithExtractAttachment extracts an attachment into a file.
func (c *Command) WithExtractAttachment(filename string) *Command {
	c.args = append(c.args, "-dump_attachment", filename)
	return c
}

// WithMaximumDemuxDecodeDelay sets maximum demux-decode delay.
func (c *Command) WithMaximumDemuxDecodeDelay(seconds int) *Command {
	c.args = append(c.args, "-muxdelay", strconv.Itoa(seconds))
	return c
}

// WithInitialDemuxDecodeDelay sets initial demux-decode delay.
func (c *Command) WithInitialDemuxDecodeDelay(seconds int) *Command {
	c.args = append(c.args, "-muxpreload", strconv.Itoa(seconds))
	return c
}

// WithBitstreamFilters sets bitstream filters for matching streams.
func (c *Command) WithBitstreamFilters(filters string) *Command {
	c.args = append(c.args, "-bsf", filters)
	return c
}

// WithPresetFile sets the preset file for matching stream(s).
func (c *Command) WithPresetFile(filename string) *Command {
	c.args = append(c.args, "-fpre", filename)
	return c
}

// WithForceDataCodec forces data codec (‘copy’ to copy stream).
func (c *Command) WithForceDataCodec(codec string) *Command {
	c.args = append(c.args, "-dcodec", codec)
	return c
}

// WithStreamLoop sets number of times input stream shall be looped. Loop 0 means no loop, loop -1 means infinite loop.
func (c *Command) WithStreamLoop(number int) *Command {
	c.args = append(c.args, "-stream_loop", strconv.Itoa(number))
	return c
}

// WithRecastMedia forces a decoder of a different media type than the one detected or designated by the demuxer.
// Useful for decoding media data muxed as data streams.
func (c *Command) WithRecastMedia() *Command {
	c.args = append(c.args, "-recast_media")
	return c
}
