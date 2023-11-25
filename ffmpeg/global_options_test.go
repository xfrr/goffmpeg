package ffmpeg_test

import (
	"testing"

	"github.com/xfrr/goffmpeg/v2/ffmpeg"
)

func TestGlobalOptions(t *testing.T) {
	testCases := []struct {
		name     string
		expected []string
		fn       func(*ffmpeg.Command) *ffmpeg.Command
	}{
		{
			name:     "WithReport",
			expected: []string{"-report"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithReport()
			},
		},
		{
			name:     "WithMaxAlloc",
			expected: []string{"-max_alloc", "100"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithMaxAlloc(100)
			},
		},
		{
			name:     "WithFilterThreads",
			expected: []string{"-filter_threads", "2"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithFilterThreads(2)
			},
		},
		{
			name:     "WithTimeLimit",
			expected: []string{"-timelimit", "100"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithTimeLimit(100)
			},
		},
		{
			name:     "WithAudioDriftThreshold",
			expected: []string{"-adrift_threshold", "100"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithAudioDriftThreshold(100)
			},
		},
		{
			name:     "WithCopyTS",
			expected: []string{"-copyts"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithCopyTS()
			},
		},
		{
			name:     "WithCopyTB",
			expected: []string{"-copytb", "mode"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithCopyTB("mode")
			},
		},
		{
			name:     "WithDTSDeltaThreshold",
			expected: []string{"-dts_delta_threshold", "100"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithDTSDeltaThreshold(100)
			},
		},
		{
			name:     "WithDTSErrorThreshold",
			expected: []string{"-dts_error_threshold", "100"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithDTSErrorThreshold(100)
			},
		},
		{
			name:     "WithXError",
			expected: []string{"-xerror"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithXError()
			},
		},
		{
			name:     "WithDebugTS",
			expected: []string{"-debug_ts"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithDebugTS()
			},
		},
		{
			name:     "WithVStatsFile",
			expected: []string{"-vstats_file", "file"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithVStatsFile("file")
			},
		},
		{
			name:     "WithISync",
			expected: []string{"-isync"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithISync()
			},
		},
		{
			name:     "WithMaxBitrate",
			expected: []string{"-maxrate", "100"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithMaxBitrate(100)
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
