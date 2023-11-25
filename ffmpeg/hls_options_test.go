package ffmpeg_test

import (
	"testing"

	"github.com/xfrr/goffmpeg/v2/ffmpeg"
)

func TestHlsOptions(t *testing.T) {
	testCases := []struct {
		name     string
		expected []string
		fn       func(*ffmpeg.Command) *ffmpeg.Command
	}{
		{
			name:     "WithHLSKeyInfoPath",
			expected: []string{"-hls_key_info_file", "path"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithHLSKeyInfoPath("path")
			},
		},
		{
			name:     "WithHLSSegmentTime",
			expected: []string{"-hls_time", "100"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithHLSSegmentTime(100)
			},
		},
		{
			name:     "WithHLSSegmentFilename",
			expected: []string{"-hls_segment_filename", "filename"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithHLSSegmentFilename("filename")
			},
		},
		{
			name:     "WithHLSSegmentStartNumber",
			expected: []string{"-start_number", "100"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithHLSSegmentStartNumber(100)
			},
		},
		{
			name:     "WithHLSSegmentListSize",
			expected: []string{"-hls_list_size", "100"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithHLSSegmentListSize(100)
			},
		},
		{
			name:     "WithHLSSegmentListType",
			expected: []string{"-hls_playlist_type", "typ"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithHLSSegmentListType("typ")
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
