package ffmpeg

import "strconv"

// WithVideoFrames set the number of video frames to record
func (c *Command) WithVideoFrames(frames int) *Command {
	c.args = append(c.args, "-vframes", strconv.Itoa(frames))
	return c
}

// WithFrameRate set frame rate (Hz value, fraction or abbreviation)
func (c *Command) WithFrameRate(rate string) *Command {
	c.args = append(c.args, "-r", rate)
	return c
}

// WithFrameSize set frame size (WxH or abbreviation)
func (c *Command) WithFrameSize(size string) *Command {
	c.args = append(c.args, "-s", size)
	return c
}

// WithAspectRatio set aspect ratio (4:3, 16:9 or 1.3333, 1.7777)
func (c *Command) WithAspectRatio(aspect string) *Command {
	c.args = append(c.args, "-aspect", aspect)
	return c
}

// WithDisableVideo disable video
func (c *Command) WithDisableVideo() *Command {
	c.args = append(c.args, "-vn")
	return c
}

// WithVideoCodec force video codec (‘copy’ to copy stream)
func (c *Command) WithVideoCodec(codec string) *Command {
	c.args = append(c.args, "-vcodec", codec)
	return c
}

// WithTimeCode set initial TimeCode value.
func (c *Command) WithTimeCode(timecode string) *Command {
	c.args = append(c.args, "-timecode", timecode)
	return c
}

// WithPass select the pass number (1 to 3)
func (c *Command) WithPass(pass int) *Command {
	c.args = append(c.args, "-pass", strconv.Itoa(pass))
	return c
}

// WithVideoFilters set video filters
func (c *Command) WithVideoFilters(filters string) *Command {
	c.args = append(c.args, "-vf", filters)
	return c
}

// WithVideoBitrate set video bitrate (please use -b:v)
func (c *Command) WithVideoBitrate(bitrate string) *Command {
	c.args = append(c.args, "-b:v", bitrate)
	return c
}

// WithDisableData disable data
func (c *Command) WithDisableData() *Command {
	c.args = append(c.args, "-dn")
	return c
}

// WithPixelFormat set pixel format
func (c *Command) WithPixelFormat(format string) *Command {
	c.args = append(c.args, "-pix_fmt", format)
	return c
}

// WithIntra deprecated use -g 1
func (c *Command) WithDiscardThreshold(discard int) *Command {
	c.args = append(c.args, "-vdt", strconv.Itoa(discard))
	return c
}

// WithRateControlOverride override rate control override for specific intervals
func (c *Command) WithRateControlOverride(override string) *Command {
	c.args = append(c.args, "-rc_override", override)
	return c
}

// WithPassLogFile select two pass log file name prefix
func (c *Command) WithPassLogFile(prefix string) *Command {
	c.args = append(c.args, "-passlogfile", prefix)
	return c
}

// WithDeinterlace this option is deprecated, use the yadif filter instead
func (c *Command) WithDeinterlace() *Command {
	c.args = append(c.args, "-deinterlace")
	return c
}

// WithPSNR calculate PSNR of compressed frames
func (c *Command) WithPSNR() *Command {
	c.args = append(c.args, "-psnr")
	return c
}

// WithVideoStats dump video coding statistics to file
func (c *Command) WithVideoStats() *Command {
	c.args = append(c.args, "-vstats")
	return c
}

// WithVideoStatsFile dump video coding statistics to file
func (c *Command) WithVideoStatsFile(file string) *Command {
	c.args = append(c.args, "-vstats_file", file)
	return c
}

// WithIntraMatrix specify intra matrix coeffs
func (c *Command) WithIntraMatrix(matrix string) *Command {
	c.args = append(c.args, "-intra_matrix", matrix)
	return c
}

// WithInterMatrix specify inter matrix coeffs
func (c *Command) WithInterMatrix(matrix string) *Command {
	c.args = append(c.args, "-inter_matrix", matrix)
	return c
}

// WithChromaIntraMatrix specify intra matrix coeffs
func (c *Command) WithChromaIntraMatrix(matrix string) *Command {
	c.args = append(c.args, "-chroma_intra_matrix", matrix)
	return c
}

// WithTop top=1/bottom=0/auto=-1 field first
func (c *Command) WithTop(top int) *Command {
	c.args = append(c.args, "-top", strconv.Itoa(top))
	return c
}

// WithDC precision intra_dc_precision
func (c *Command) WithDC(precision int) *Command {
	c.args = append(c.args, "-dc", strconv.Itoa(precision))
	return c
}

// WithVideoTag force video tag/fourcc
func (c *Command) WithVideoTag(tag string) *Command {
	c.args = append(c.args, "-vtag", tag)
	return c
}

func (c *Command) WithVideoProfile(profile string) *Command {
	c.args = append(c.args, "-profile:v", profile)
	return c
}

// WithQPHist show QP histogram
func (c *Command) WithQPHist() *Command {
	c.args = append(c.args, "-qphist")
	return c
}

// WithForceFPS force the selected framerate, disable the best supported framerate selection
func (c *Command) WithForceFPS() *Command {
	c.args = append(c.args, "-force_fps")
	return c
}

// WithForceKeyFrames force key frames at specified timestamps
func (c *Command) WithForceKeyFrames(timestamps string) *Command {
	c.args = append(c.args, "-force_key_frames", timestamps)
	return c
}

// WithHWAccel use HW accelerated decoding
func (c *Command) WithHWAccel(name string) *Command {
	c.args = append(c.args, "-hwaccel", name)
	return c
}

// WithHWAccelDevice select a device for HW accelerationdevicename
func (c *Command) WithHWAccelDevice(device string) *Command {
	c.args = append(c.args, "-hwaccel_device", device)
	return c
}

// WithChannel deprecated, use -channel
func (c *Command) WithChannel(channel string) *Command {
	c.args = append(c.args, "-channel", channel)
	return c
}

// WithTVStandard deprecated, use -standard
func (c *Command) WithTVStandard(standard string) *Command {
	c.args = append(c.args, "-standard", standard)
	return c
}

// WithVideoBitstreamFilters deprecated
func (c *Command) WithVideoBitstreamFilters(filters string) *Command {
	c.args = append(c.args, "-vbsf", filters)
	return c
}

// WithVideoPreset set the video options to the indicated preset
func (c *Command) WithVideoPreset(preset string) *Command {
	c.args = append(c.args, "-vpre", preset)
	return c
}
