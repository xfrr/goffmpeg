package ffmpeg

import "strconv"

// WithAudioFrames sets the number of audio frames to record.
func (c *Command) WithAudioFrames(frames int) *Command {
	c.args = append(c.args, "-aframes", strconv.Itoa(frames))
	return c
}

// WithAudioQuality sets the audio quality.
func (c *Command) WithAudioQuality(quality int) *Command {
	c.args = append(c.args, "-aq", strconv.Itoa(quality))
	return c
}

// WithAudioRate sets the audio sampling rate.
func (c *Command) WithAudioRate(rate int) *Command {
	c.args = append(c.args, "-ar", strconv.Itoa(rate))
	return c
}

// WithAudioChannels sets the number of audio channels.
func (c *Command) WithAudioChannels(channels int) *Command {
	c.args = append(c.args, "-ac", strconv.Itoa(channels))
	return c
}

// WithDisableAudio disables audio.
func (c *Command) WithDisableAudio() *Command {
	c.args = append(c.args, "-an")
	return c
}

// WithAudioCodec sets the audio codec.
func (c *Command) WithAudioCodec(codec string) *Command {
	c.args = append(c.args, "-acodec", codec)
	return c
}

// WithVolume sets the audio volume.
func (c *Command) WithVolume(volume int) *Command {
	c.args = append(c.args, "-vol", strconv.Itoa(volume))
	return c
}

// WithAudioFilters creates a complex filtergraph.
func (c *Command) WithAudioFilters(filters string) *Command {
	c.args = append(c.args, "-af", filters)
	return c
}

// WithAudioTag sets the audio tag.
func (c *Command) WithAudioTag(tag string) *Command {
	c.args = append(c.args, "-atag", tag)
	return c
}

// WithAudioSampleFormat sets the sample format.
func (c *Command) WithAudioSampleFormat(format string) *Command {
	c.args = append(c.args, "-sample_fmt", format)
	return c
}

// WithAudioChannelLayout sets the channel layout.
func (c *Command) WithAudioChannelLayout(layout string) *Command {
	c.args = append(c.args, "-channel_layout", layout)
	return c
}

// WithAudioGuessLayoutMax sets the maximum number of channels to try to guess the channel layout.
func (c *Command) WithAudioGuessLayoutMax(max int) *Command {
	c.args = append(c.args, "-guess_layout_max", strconv.Itoa(max))
	return c
}

// WithAudioBitstreamFilters sets the audio bitstream filters.
func (c *Command) WithAudioBitstreamFilters(filters string) *Command {
	c.args = append(c.args, "-absf", filters)
	return c
}

// WithAudioPreset sets the audio options to the indicated preset.
func (c *Command) WithAudioPreset(preset string) *Command {
	c.args = append(c.args, "-apre", preset)
	return c
}

// WithAudioProfile sets the audio profile.
func (c *Command) WithAudioProfile(profile string) *Command {
	c.args = append(c.args, "-profile:a", profile)
	return c
}
