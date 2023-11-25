package ffmpeg

import "strconv"

func (c *Command) WithHLSKeyInfoPath(path string) *Command {
	c.args = append(c.args, "-hls_key_info_file", path)
	return c
}

func (c *Command) WithHLSSegmentTime(seconds int) *Command {
	c.args = append(c.args, "-hls_time", strconv.Itoa(seconds))
	return c
}

func (c *Command) WithHLSSegmentFilename(filename string) *Command {
	c.args = append(c.args, "-hls_segment_filename", filename)
	return c
}

func (c *Command) WithHLSSegmentStartNumber(number int) *Command {
	c.args = append(c.args, "-start_number", strconv.Itoa(number))
	return c
}

func (c *Command) WithHLSSegmentListSize(size int) *Command {
	c.args = append(c.args, "-hls_list_size", strconv.Itoa(size))
	return c
}

func (c *Command) WithHLSSegmentListType(typ string) *Command {
	c.args = append(c.args, "-hls_playlist_type", typ)
	return c
}
