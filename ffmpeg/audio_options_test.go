package ffmpeg_test

import (
	"testing"

	"github.com/xfrr/goffmpeg/v2/ffmpeg"
)

func TestAudioOptions(t *testing.T) {
	testCases := []struct {
		name     string
		expected []string
		fn       func(*ffmpeg.Command) *ffmpeg.Command
	}{
		{
			name:     "WithAudioFrames",
			expected: []string{"-aframes", "10"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithAudioFrames(10)
			},
		},
		{
			name:     "WithAudioQuality",
			expected: []string{"-aq", "10"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithAudioQuality(10)
			},
		},
		{
			name:     "WithAudioRate",
			expected: []string{"-ar", "10"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithAudioRate(10)
			},
		},
		{
			name:     "WithAudioChannels",
			expected: []string{"-ac", "10"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithAudioChannels(10)
			},
		},
		{
			name:     "WithDisableAudio",
			expected: []string{"-an"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithDisableAudio()
			},
		},
		{
			name:     "WithAudioCodec",
			expected: []string{"-acodec", "10"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithAudioCodec("10")
			},
		},
		{
			name:     "WithVolume",
			expected: []string{"-vol", "10"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithVolume(10)
			},
		},
		{
			name:     "WithAudioFilters",
			expected: []string{"-af", "10"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithAudioFilters("10")
			},
		},
		{
			name:     "WithAudioTag",
			expected: []string{"-atag", "10"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithAudioTag("10")
			},
		},
		{
			name:     "WithAudioSampleFormat",
			expected: []string{"-sample_fmt", "10"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithAudioSampleFormat("10")
			},
		},
		{
			name:     "WithAudioChannelLayout",
			expected: []string{"-channel_layout", "10"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithAudioChannelLayout("10")
			},
		},
		{
			name:     "WithAudioGuessLayoutMax",
			expected: []string{"-guess_layout_max", "10"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithAudioGuessLayoutMax(10)
			},
		},
		{
			name:     "WithAudioBitstreamFilters",
			expected: []string{"-absf", "10"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithAudioBitstreamFilters("10")
			},
		},
		{
			name:     "WithAudioPreset",
			expected: []string{"-apreset", "10"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithAudioPreset("10")
			},
		},
		{
			name:     "WithAudioProfile",
			expected: []string{"-profile:a", "10"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithAudioProfile("10")
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cmd := ffmpeg.NewCommand()
			cmd = tc.fn(cmd)
			if len(cmd.Args()) != len(tc.expected)+len(ffmpeg.DefaultArgs()) {
				t.Fatalf("expected %d args, got %d", len(tc.expected), len(cmd.Args()))
			}

			for i := len(ffmpeg.DefaultArgs()) - 1; i < len(tc.expected); i++ {
				if cmd.Args()[i] != tc.expected[i] {
					t.Fatalf("expected %s, got %s", tc.expected[i], cmd.Args()[i])
				}
			}
		})
	}
}
