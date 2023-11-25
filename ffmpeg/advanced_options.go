package ffmpeg

// WithMap creates one or more streams in the output file. This option has two forms for specifying the data source(s): the first selects one or more streams from some input file (specified with -i), the second takes an output from some complex filtergraph (specified with -filter_complex or -filter_complex_script).
func (c *Command) WithMap(m string) *Command {
	c.args = append(c.args, "-map", m)
	return c
}

// WithIgnoreUnknown ignore input streams with unknown type instead of failing if copying such streams is attempted.
func (c *Command) WithIgnoreUnknown() *Command {
	c.args = append(c.args, "-ignore_unknown")
	return c
}

// WithCopyUnknown allow input streams with unknown type to be copied instead of failing if copying such streams is attempted.
func (c *Command) WithCopyUnknown() *Command {
	c.args = append(c.args, "-copy_unknown")
	return c
}

// WithMapChannel creates an audio channel from a given input to an output. If output_file_id.stream_specifier is not set, the audio channel will be mapped on all the audio streams.
func (c *Command) WithMapChannel(m string) *Command {
	c.args = append(c.args, "-map_channel", m)
	return c
}

// WithMapMetadata sets metadata information of the next output file from infile. Note that those are file indices (zero-based), not filenames. Optional metadata_spec_in/out parameters specify, which metadata to copy. A metadata specifier can have the following forms:
func (c *Command) WithMapMetadata(m string) *Command {
	c.args = append(c.args, "-map_metadata", m)
	return c
}

// WithMapChapters copy chapters from input file with index input_file_index to the next output file. If no chapter mapping is specified, then chapters are copied from the first input file with at least one chapter. Use a negative file index to disable any chapter copying.
func (c *Command) WithMapChapters(m string) *Command {
	c.args = append(c.args, "-map_chapters", m)
	return c
}

// WithTimelimit exit after ffmpeg has been running for duration seconds in CPU user time.
func (c *Command) WithTimelimit(t string) *Command {
	c.args = append(c.args, "-timelimit", t)
	return c
}

// WithDump dump each input packet to stderr.
func (c *Command) WithDump() *Command {
	c.args = append(c.args, "-dump")
	return c
}

// WithHex when dumping packets, also dump the payload.
func (c *Command) WithHex() *Command {
	c.args = append(c.args, "-hex")
	return c
}

// WithVideoSync set video sync method.
func (c *Command) WithVideoSync(v string) *Command {
	c.args = append(c.args, "-vsync", v)
	return c
}

// WithFpsMode set video sync method.
func (c *Command) WithFpsMode(v string) *Command {
	c.args = append(c.args, "-fps_mode", v)
	return c
}

// WithFrameDropThreshold set frame drop threshold, which specifies how much behind video frames can be before they are dropped. In frame rate units, so 1.0 is one frame. The default is -1.1. One possible usecase is to avoid framedrops in case of noisy timestamps or to increase frame drop precision in case of exact timestamps.
func (c *Command) WithFrameDropThreshold(f string) *Command {
	c.args = append(c.args, "-frame_drop_threshold", f)
	return c
}

// WithApad pad the output audio stream(s). This is the same as applying -af apad. Argument is a string of filter parameters composed the same as with the apad filter. -shortest must be set for this output for the option to take effect.
func (c *Command) WithApad(a string) *Command {
	c.args = append(c.args, "-apad", a)
	return c
}

// WithCopyts do not process input timestamps, but keep their values without trying to sanitize them. In particular, do not remove the initial start time offset value.
func (c *Command) WithCopyts() *Command {
	c.args = append(c.args, "-copyts")
	return c
}

// WithStartAtZero when used with copyts, shift input timestamps so they start at zero.
func (c *Command) WithStartAtZero() *Command {
	c.args = append(c.args, "-start_at_zero")
	return c
}

// WithCopytb specify how to set the encoder timebase when stream copying. mode is an integer numeric value, and can assume one of the following values:
func (c *Command) WithCopytb(m string) *Command {
	c.args = append(c.args, "-copytb", m)
	return c
}

// WithEncTimeBase set the encoder timebase. timebase can assume one of the following values:
func (c *Command) WithEncTimeBase(e string) *Command {
	c.args = append(c.args, "-enc_time_base", e)
	return c
}

// WithBitexact enable bitexact mode for (de)muxer and (de/en)coder
func (c *Command) WithBitexact() *Command {
	c.args = append(c.args, "-bitexact")
	return c
}

// WithShortest finish encoding when the shortest output stream ends.
func (c *Command) WithShortest() *Command {
	c.args = append(c.args, "-shortest")
	return c
}

// WithShortestBufDuration the -shortest option may require buffering potentially large amounts of data when at least one of the streams is "sparse" (i.e. has large gaps between frames – this is typically the case for subtitles).
func (c *Command) WithShortestBufDuration(s string) *Command {
	c.args = append(c.args, "-shortest_buf_duration", s)
	return c
}

// WithDtsDeltaThreshold timestamp discontinuity delta threshold, expressed as a decimal number of seconds.
func (c *Command) WithDtsDeltaThreshold(d string) *Command {
	c.args = append(c.args, "-dts_delta_threshold", d)
	return c
}

// WithDtsErrorThreshold timestamp error delta threshold, expressed as a decimal number of seconds.
func (c *Command) WithDtsErrorThreshold(d string) *Command {
	c.args = append(c.args, "-dts_error_threshold", d)
	return c
}

// WithMuxdelay set the maximum demux-decode delay.
func (c *Command) WithMuxdelay(m string) *Command {
	c.args = append(c.args, "-muxdelay", m)
	return c
}

// WithMuxpreload set the initial demux-decode delay.
func (c *Command) WithMuxpreload(m string) *Command {
	c.args = append(c.args, "-muxpreload", m)
	return c
}

// WithStreamID assign a new stream-id value to an output stream. This option should be specified prior to the output filename to which it applies. For the situation where multiple output files exist, a streamid may be reassigned to a different value.
func (c *Command) WithStreamID(s string) *Command {
	c.args = append(c.args, "-streamid", s)
	return c
}

// WithBsf set bitstream filters for matching streams. bitstream_filters is a comma-separated list of bitstream filters. Use the -bsfs option to get the list of bitstream filters.
func (c *Command) WithBsf(b string) *Command {
	c.args = append(c.args, "-bsf", b)
	return c
}

// WithTag force a tag/fourcc for matching streams.
func (c *Command) WithTag(t string) *Command {
	c.args = append(c.args, "-tag", t)
	return c
}

// WithTimecode specify Timecode for writing. SEP is ’:’ for non drop timecode and ’;’ (or ’.’) for drop.
func (c *Command) WithTimecode(t string) *Command {
	c.args = append(c.args, "-timecode", t)
	return c
}

// WithFilterComplex define a complex filtergraph, i.e. one with arbitrary number of inputs and/or outputs. For simple graphs – those with one input and one output of the same type – see the -filter options. filtergraph is a description of the filtergraph, as described in the “Filtergraph syntax” section of the ffmpeg-filters manual.
func (c *Command) WithFilterComplex(f string) *Command {
	c.args = append(c.args, "-filter_complex", f)
	return c
}

// WithFilterComplexThreads defines how many threads are used to process a filter_complex graph. Similar to filter_threads but used for -filter_complex graphs only. The default is the number of available CPUs.
func (c *Command) WithFilterComplexThreads(f string) *Command {
	c.args = append(c.args, "-filter_complex_threads", f)
	return c
}

// WithLavfi define a complex filtergraph, i.e. one with arbitrary number of inputs and/or outputs. Equivalent to -filter_complex.
func (c *Command) WithLavfi(l string) *Command {
	c.args = append(c.args, "-lavfi", l)
	return c
}

// WithFilterComplexScript this option is similar to -filter_complex, the only difference is that its argument is the name of the file from which a complex filtergraph description is to be read.
func (c *Command) WithFilterComplexScript(f string) *Command {
	c.args = append(c.args, "-filter_complex_script", f)
	return c
}

// WithAccurateSeek this option enables or disables accurate seeking in input files with the -ss option. It is enabled by default, so seeking is accurate when transcoding. Use -noaccurate_seek to disable it, which may be useful e.g. when copying some streams and transcoding the others.
func (c *Command) WithAccurateSeek() *Command {
	c.args = append(c.args, "-accurate_seek")
	return c
}

// WithSeekTimestamp this option enables or disables seeking by timestamp in input files with the -ss option. It is disabled by default. If enabled, the argument to the -ss option is considered an actual timestamp, and is not offset by the start time of the file. This matters only for files which do not start from timestamp 0, such as transport streams.
func (c *Command) WithSeekTimestamp() *Command {
	c.args = append(c.args, "-seek_timestamp")
	return c
}

// WithThreadQueueSize for input, this option sets the maximum number of queued packets when reading from the file or device. With low latency / high rate live streams, packets may be discarded if they are not read in a timely manner; setting this value can force ffmpeg to use a separate input thread and read packets as soon as they arrive. By default ffmpeg only does this if multiple inputs are specified.
func (c *Command) WithThreadQueueSize(t string) *Command {
	c.args = append(c.args, "-thread_queue_size", t)
	return c
}

// WithSdpFile print sdp information for an output stream to file. This allows dumping sdp information when at least one output isn’t an rtp stream. (Requires at least one of the output formats to be rtp).
func (c *Command) WithSdpFile(s string) *Command {
	c.args = append(c.args, "-sdp_file", s)
	return c
}

// WithDiscard allows discarding specific streams or frames from streams. Any input stream can be fully discarded, using value all whereas selective discarding of frames from a stream occurs at the demuxer and is not supported by all demuxers.
func (c *Command) WithDiscard(d string) *Command {
	c.args = append(c.args, "-discard", d)
	return c
}

// WithAbortOn Stop and abort on various conditions. The following flags are available:
//
// empty_output: No packets were passed to the muxer, the output is empty.
//
// empty_output_stream: No packets were passed to the muxer in some of the output streams.
func (c *Command) WithAbortOn(a string) *Command {
	c.args = append(c.args, "-abort_on", a)
	return c
}

// WithMaxErrorRate set fraction of decoding frame failures across all inputs which when crossed ffmpeg will return exit code 69. Crossing this threshold does not terminate processing. Range is a floating-point number between 0 to 1. Default is 2/3.
func (c *Command) WithMaxErrorRate(m string) *Command {
	c.args = append(c.args, "-max_error_rate", m)
	return c
}

// WithMaxMuxingQueueSize when transcoding audio and/or video streams, ffmpeg will not begin writing into the output until it has one packet for each such stream. While waiting for that to happen, packets for other streams are buffered. This option sets the size of this buffer, in packets, for the matching output stream.
func (c *Command) WithMaxMuxingQueueSize(m string) *Command {
	c.args = append(c.args, "-max_muxing_queue_size", m)
	return c
}

// WithMuxingQueueDataThreshold this is a minimum threshold until which the muxing queue size is not taken into account. Defaults to 50 megabytes per stream, and is based on the overall size of packets passed to the muxer.
func (c *Command) WithMuxingQueueDataThreshold(m string) *Command {
	c.args = append(c.args, "-muxing_queue_data_threshold", m)
	return c
}

// WithAutoConversionFilters enable automatically inserting format conversion filters in all filter graphs, including those defined by -vf, -af, -filter_complex and -lavfi. If filter format negotiation requires a conversion, the initialization of the filters will fail. Conversions can still be performed by inserting the relevant conversion filter (scale, aresample) in the graph. On by default, to explicitly disable it you need to specify -noauto_conversion_filters.
func (c *Command) WithAutoConversionFilters() *Command {
	c.args = append(c.args, "-auto_conversion_filters")
	return c
}

// WithBitsPerRawSample declare the number of bits per raw sample in the given output stream to be value. Note that this option sets the information provided to the encoder/muxer, it does not change the stream to conform to this value. Setting values that do not match the stream properties may result in encoding failures or invalid output files.
func (c *Command) WithBitsPerRawSample(b string) *Command {
	c.args = append(c.args, "-bits_per_raw_sample", b)
	return c
}

// WithStatsEncPre write per-frame encoding information about the matching streams into the file given by path.
func (c *Command) WithStatsEncPre(s string) *Command {
	c.args = append(c.args, "-stats_enc_pre", s)
	return c
}

// WithStatsEncPost write per-frame encoding information about the matching streams into the file given by path.
func (c *Command) WithStatsEncPost(s string) *Command {
	c.args = append(c.args, "-stats_enc_post", s)
	return c
}

// WithStatsMuxPre write per-frame encoding information about the matching streams into the file given by path.
func (c *Command) WithStatsMuxPre(s string) *Command {
	c.args = append(c.args, "-stats_mux_pre", s)
	return c
}

// WithStatsEncPreFmt specify the format for the lines written with -stats_enc_pre / -stats_enc_post / -stats_mux_pre.
func (c *Command) WithStatsEncPreFmt(s string) *Command {
	c.args = append(c.args, "-stats_enc_pre_fmt", s)
	return c
}

// WithStatsEncPostFmt specify the format for the lines written with -stats_enc_pre / -stats_enc_post / -stats_mux_pre.
func (c *Command) WithStatsEncPostFmt(s string) *Command {
	c.args = append(c.args, "-stats_enc_post_fmt", s)
	return c
}

// WithStatsMuxPreFmt specify the format for the lines written with -stats_enc_pre / -stats_enc_post / -stats_mux_pre.
func (c *Command) WithStatsMuxPreFmt(s string) *Command {
	c.args = append(c.args, "-stats_mux_pre_fmt", s)
	return c
}
