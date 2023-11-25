package ffmpeg_test

import (
	"testing"

	"github.com/xfrr/goffmpeg/v2/ffmpeg"
)

func TestMainOptions(t *testing.T) {
	testCases := []struct {
		name     string
		expected []string
		fn       func(*ffmpeg.Command) *ffmpeg.Command
	}{
		{
			name:     "WithInputPath",
			expected: []string{"-i", "input"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithInputPath("input")
			},
		},
		{
			name:     "WithInputPipe",
			expected: []string{"-i", "pipe:0"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithInputPipe(nil)
			},
		},
		{
			name:     "WithOutputPath",
			expected: []string{"output"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithOutputPath("output")
			},
		},
		{
			name:     "WithOutputPipe",
			expected: []string{"pipe:1"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithOutputPipe(nil)
			},
		},
		{
			name:     "WithOutputFormat",
			expected: []string{"-f", "format"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithOutputFormat("format")
			},
		},
		{
			name:     "WithCodec",
			expected: []string{"-c", "codec"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithCodec("codec")
			},
		},
		{
			name:     "WithPreset",
			expected: []string{"-preset", "preset"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithPreset("preset")
			},
		},
		{
			name:     "WithDuration",
			expected: []string{"-t", "duration"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithDuration("duration")
			},
		},
		{
			name:     "WithStopTime",
			expected: []string{"-to", "time"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithStopTime("time")
			},
		},
		{
			name:     "WithLimitFileSize",
			expected: []string{"-fs", "100"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithLimitFileSize(100)
			},
		},
		{
			name:     "WithStartTimeOffset",
			expected: []string{"-ss", "100"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithStartTimeOffset("100")
			},
		},
		{
			name:     "WithTimestamp",
			expected: []string{"-timestamp", "100"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithTimestamp("100")
			},
		},
		{
			name:     "WithMetadata",
			expected: []string{"-metadata", "metadata"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithMetadata("metadata")
			},
		},
		{
			name:     "WithTarget",
			expected: []string{"-target", "typ"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithTarget("typ")
			},
		},
		{
			name:     "WithAudioPad",
			expected: []string{"-apad"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithAudioPad()
			},
		},
		{
			name:     "WithFrames",
			expected: []string{"-frames", "100"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithFrames(100)
			},
		},
		{
			name:     "WithFilter",
			expected: []string{"-filter", "filter"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithFilter("filter")
			},
		},
		{
			name:     "WithFilterScript",
			expected: []string{"-filter_script", "filename"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithFilterScript("filename")
			},
		},
		{
			name:     "WithReinitFilter",
			expected: []string{"-reinit_filter"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithReinitFilter()
			},
		},
		{
			name:     "WithInputTsOffset",
			expected: []string{"-itsoffset", "offset"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithInputTsOffset("offset")
			},
		},
		{
			name:     "WithInputTsScale",
			expected: []string{"-itsscale", "scale"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithInputTsScale("scale")
			},
		},
		{
			name:     "WithDataFrames",
			expected: []string{"-data_frames", "100"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithDataFrames(100)
			},
		},
		{
			name:     "WithReadInputAtNativeFrameRate",
			expected: []string{"-re"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithReadInputAtNativeFrameRate()
			},
		},
		{
			name:     "WithCopyInitialNonKeyframes",
			expected: []string{"-copyinkf"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithCopyInitialNonKeyframes()
			},
		},
		{
			name:     "WithCopyOrDiscardFramesBeforeStartTime",
			expected: []string{"-copyinkf"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithCopyOrDiscardFramesBeforeStartTime()
			},
		},
		{
			name:     "WithForceCodecTag",
			expected: []string{"-codec_tag", "force"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithForceCodecTag("force")
			},
		},
		{
			name:     "WithFixedQualityScale",
			expected: []string{"-q", "100"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithFixedQualityScale(100)
			},
		},
		{
			name:     "WithAttachment",
			expected: []string{"-attach", "filename"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithAttachment("filename")
			},
		},
		{
			name:     "WithExtractAttachment",
			expected: []string{"-dump_attachment", "filename"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithExtractAttachment("filename")
			},
		},
		{
			name:     "WithMaximumDemuxDecodeDelay",
			expected: []string{"-max_delay", "100"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithMaximumDemuxDecodeDelay(100)
			},
		},
		{
			name:     "WithInitialDemuxDecodeDelay",
			expected: []string{"-muxpreload", "100"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithInitialDemuxDecodeDelay(100)
			},
		},
		{
			name:     "WithBitstreamFilters",
			expected: []string{"-bsf", "filters"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithBitstreamFilters("filters")
			},
		},
		{
			name:     "WithPresetFile",
			expected: []string{"-fpre", "filename"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithPresetFile("filename")
			},
		},
		{
			name:     "WithForceDataCodec",
			expected: []string{"-codec", "codec"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithForceDataCodec("codec")
			},
		},
		{
			name:     "WithStreamLoop",
			expected: []string{"-stream_loop", "100"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithStreamLoop(100)
			},
		},
		{
			name:     "WithRecastMedia",
			expected: []string{"-recast_media"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithRecastMedia()
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
