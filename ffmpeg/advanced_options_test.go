package ffmpeg_test

import (
	"testing"

	"github.com/xfrr/goffmpeg/v2/ffmpeg"
)

func TestAdvancedOptions(t *testing.T) {
	testCases := []struct {
		name     string
		expected []string
		fn       func(*ffmpeg.Command) *ffmpeg.Command
	}{
		{
			name:     "WithMap",
			expected: []string{"-map", "0"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithMap("0")
			},
		},
		{
			name:     "WithIgnoreUnknown",
			expected: []string{"-ignore_unknown"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithIgnoreUnknown()
			},
		},
		{
			name:     "WithCopyUnknown",
			expected: []string{"-copy_unknown"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithCopyUnknown()
			},
		},
		{
			name:     "WithMapChannel",
			expected: []string{"-map_channel", "0.0.0"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithMapChannel("0.0.0")
			},
		},
		{
			name:     "WithMapMetadata",
			expected: []string{"-map_metadata", "0"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithMapMetadata("0")
			},
		},
		{
			name:     "WithMapChapters",
			expected: []string{"-map_chapters", "0"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithMapChapters("0")
			},
		},
		{
			name:     "WithTimelimit",
			expected: []string{"-timelimit", "0"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithTimelimit("0")
			},
		},
		{
			name:     "WithDump",
			expected: []string{"-dump"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithDump()
			},
		},
		{
			name:     "WithHex",
			expected: []string{"-hex"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithHex()
			},
		},
		{
			name:     "WithVideoSync",
			expected: []string{"-vsync", "0"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithVideoSync("0")
			},
		},
		{
			name:     "WithFpsMode",
			expected: []string{"-fpsprobesize", "0"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithFpsMode("0")
			},
		},
		{
			name:     "WithFrameDropThreshold",
			expected: []string{"-frame_drop_threshold", "0"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithFrameDropThreshold("0")
			},
		},
		{
			name:     "WithApad",
			expected: []string{"-apad", "0"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithApad("0")
			},
		},
		{
			name:     "WithCopyts",
			expected: []string{"-copyts"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithCopyts()
			},
		},
		{
			name:     "WithStartAtZero",
			expected: []string{"-start_at_zero"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithStartAtZero()
			},
		},
		{
			name:     "WithCopytb",
			expected: []string{"-copytb", "0"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithCopytb("0")
			},
		},
		{
			name:     "WithEncTimeBase",
			expected: []string{"-enc_time_base", "0"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithEncTimeBase("0")
			},
		},
		{
			name:     "WithBitexact",
			expected: []string{"-bitexact"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithBitexact()
			},
		},
		{
			name:     "WithShortest",
			expected: []string{"-shortest"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithShortest()
			},
		},
		{
			name:     "WithShortestBufDuration",
			expected: []string{"-shortest_bufsize", "0"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithShortestBufDuration("0")
			},
		},
		{
			name:     "WithDtsDeltaThreshold",
			expected: []string{"-dts_delta_threshold", "0"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithDtsDeltaThreshold("0")
			},
		},
		{
			name:     "WithDtsErrorThreshold",
			expected: []string{"-dts_error_threshold", "0"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithDtsErrorThreshold("0")
			},
		},
		{
			name:     "WithMuxdelay",
			expected: []string{"-muxdelay", "0"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithMuxdelay("0")
			},
		},
		{
			name:     "WithMuxpreload",
			expected: []string{"-muxpreload", "0"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithMuxpreload("0")
			},
		},
		{
			name:     "WithStreamID",
			expected: []string{"-streamId", "0"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithStreamID("0")
			},
		},
		{
			name:     "WithBsf",
			expected: []string{"-bsf", "0"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithBsf("0")
			},
		},
		{
			name:     "WithTag",
			expected: []string{"-tag", "0"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithTag("0")
			},
		},
		{
			name:     "WithTimecode",
			expected: []string{"-timecode", "0"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithTimecode("0")
			},
		},
		{
			name:     "WithFilterComplex",
			expected: []string{"-filter_complex", "dummy"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithFilterComplex("dummy")
			},
		},
		{
			name:     "WithFilterComplexThreads",
			expected: []string{"-filter_complex_threads", "0"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithFilterComplexThreads("0")
			},
		},
		{
			name:     "WithLavfi",
			expected: []string{"-lavfi", "0"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithLavfi("0")
			},
		},
		{
			name:     "WithFilterComplexScript",
			expected: []string{"-filter_complex_script", "dummy"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithFilterComplexScript("dummy")
			},
		},
		{
			name:     "WithAccurateSeek",
			expected: []string{"-accurate_seek"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithAccurateSeek()
			},
		},
		{
			name:     "WithSeekTimestamp",
			expected: []string{"-seek_timestamp"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithSeekTimestamp()
			},
		},
		{
			name:     "WithThreadQueueSize",
			expected: []string{"-thread_queue_size", "0"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithThreadQueueSize("0")
			},
		},
		{
			name:     "WithSdpFile",
			expected: []string{"-sdp_file", "dummy"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithSdpFile("dummy")
			},
		},
		{
			name:     "WithDiscard",
			expected: []string{"-discard", "0"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithDiscard("0")
			},
		},
		{
			name:     "WithAbortOn",
			expected: []string{"-abort_on", "empty_output"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithAbortOn("empty_output")
			},
		},
		{
			name:     "WithMaxErrorRate",
			expected: []string{"-max_error_rate", "0.0"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithMaxErrorRate("0.0")
			},
		},
		{
			name:     "WithMaxMuxingQueueSize",
			expected: []string{"-max_muxing_queue_size", "0"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithMaxMuxingQueueSize("0")
			},
		},
		{
			name:     "WithMuxingQueueDataThreshold",
			expected: []string{"-muxing_queue_data_threshold", "0"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithMuxingQueueDataThreshold("0")
			},
		},
		{
			name:     "WithAutoConversionFilters",
			expected: []string{"-auto_conversion_filters"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithAutoConversionFilters()
			},
		},
		{
			name:     "WithBitsPerRawSample",
			expected: []string{"-bits_per_raw_sample", "0"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithBitsPerRawSample("0")
			},
		},
		{
			name:     "WithStatsEncPre",
			expected: []string{"-stats_enc_pre", "0"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithStatsEncPre("0")
			},
		},
		{
			name:     "WithStatsEncPost",
			expected: []string{"-stats_enc_post", "0"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithStatsEncPost("0")
			},
		},
		{
			name:     "WithStatsMuxPre",
			expected: []string{"-stats_mux_pre", "0"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithStatsMuxPre("0")
			},
		},
		{
			name:     "WithStatsEncPreFmt",
			expected: []string{"-stats_enc_pre_fmt", "0"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithStatsEncPreFmt("0")
			},
		},
		{
			name:     "WithStatsEncPostFmt",
			expected: []string{"-stats_enc_post_fmt", "0"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithStatsEncPostFmt("0")
			},
		},
		{
			name:     "WithStatsMuxPreFmt",
			expected: []string{"-stats_mux_pre_fmt", "0"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithStatsMuxPreFmt("0")
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cmd := ffmpeg.NewCommand()
			cmd = tc.fn(cmd)

			expectedLen := len(tc.expected) + len(ffmpeg.DefaultArgs())
			if len(cmd.Args()) != expectedLen {
				t.Fatalf("expected %d, got %d", expectedLen, len(cmd.Args()))
			}

			for i := len(ffmpeg.DefaultArgs()) - 1; i < len(tc.expected); i++ {
				if cmd.Args()[i] != tc.expected[i] {
					t.Fatalf("expected %s, got %s", tc.expected[i], cmd.Args()[i])
				}
			}
		})
	}
}
