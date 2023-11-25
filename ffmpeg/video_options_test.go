package ffmpeg_test

import (
	"testing"

	"github.com/xfrr/goffmpeg/v2/ffmpeg"
)

func TestVideoOptions(t *testing.T) {
	testCases := []struct {
		name     string
		expected []string
		fn       func(*ffmpeg.Command) *ffmpeg.Command
	}{
		{
			name:     "WithVideoFrames",
			expected: []string{"-vframes", "10"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithVideoFrames(10)
			},
		},
		{
			name:     "WithFrameRate",
			expected: []string{"-r", "10"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithFrameRate("10")
			},
		},
		{
			name:     "WithFrameSize",
			expected: []string{"-s", "10"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithFrameSize("10")
			},
		},
		{
			name:     "WithAspectRatio",
			expected: []string{"-aspect", "10"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithAspectRatio("10")
			},
		},
		{
			name:     "WithDisableVideo",
			expected: []string{"-vn"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithDisableVideo()
			},
		},
		{
			name:     "WithVideoCodec",
			expected: []string{"-vcodec", "10"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithVideoCodec("10")
			},
		},
		{
			name:     "WithTimeCode",
			expected: []string{"-timecode", "10"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithTimeCode("10")
			},
		},
		{
			name:     "WithPass",
			expected: []string{"-pass", "10"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithPass(10)
			},
		},
		{
			name:     "WithVideoFilters",
			expected: []string{"-vf", "10"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithVideoFilters("10")
			},
		},
		{
			name:     "WithVideoBitrate",
			expected: []string{"-b:v", "10"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithVideoBitrate("10")
			},
		},
		{
			name:     "WithDisableData",
			expected: []string{"-dn"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithDisableData()
			},
		},
		{
			name:     "WithPixelFormat",
			expected: []string{"-pix_fmt", "sample"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithPixelFormat("sample")
			},
		},
		{
			name:     "WithDiscardThreshold",
			expected: []string{"-vdt", "10"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithDiscardThreshold(10)
			},
		},
		{
			name:     "WithRateControlOverride",
			expected: []string{"-rc_override", "10"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithRateControlOverride("10")
			},
		},
		{
			name:     "WithPassLogFile",
			expected: []string{"-passlogfile", "sample"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithPassLogFile("sample")
			},
		},
		{
			name:     "WithDeinterlace",
			expected: []string{"-deinterlace"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithDeinterlace()
			},
		},
		{
			name:     "WithPSNR",
			expected: []string{"-psnr"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithPSNR()
			},
		},
		{
			name:     "WithVideoStats",
			expected: []string{"-vstats"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithVideoStats()
			},
		},
		{
			name:     "WithVideoStatsFile",
			expected: []string{"-vstats_file", "sample"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithVideoStatsFile("sample")
			},
		},
		{
			name:     "WithIntraMatrix",
			expected: []string{"-intra_matrix", "sample"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithIntraMatrix("sample")
			},
		},
		{
			name:     "WithInterMatrix",
			expected: []string{"-inter_matrix", "sample"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithInterMatrix("sample")
			},
		},
		{
			name:     "WithChromaIntraMatrix",
			expected: []string{"-chroma_intra_matrix", "sample"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithChromaIntraMatrix("sample")
			},
		},
		{
			name:     "WithTop",
			expected: []string{"-top", "10"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithTop(10)
			},
		},
		{
			name:     "WithDC",
			expected: []string{"-dc", "10"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithDC(10)
			},
		},
		{
			name:     "WithVideoTag",
			expected: []string{"-vtag", "sample"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithVideoTag("sample")
			},
		},
		{
			name:     "WithVideoProfile",
			expected: []string{"-profile:v", "sample"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithVideoProfile("sample")
			},
		},
		{
			name:     "WithQPHist",
			expected: []string{"-qphist"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithQPHist()
			},
		},
		{
			name:     "WithForceFPS",
			expected: []string{"-force_fps"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithForceFPS()
			},
		},
		{
			name:     "WithForceKeyFrames",
			expected: []string{"-force_key_frames", "sample"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithForceKeyFrames("sample")
			},
		},
		{
			name:     "WithHWAccel",
			expected: []string{"-hwaccel", "sample"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithHWAccel("sample")
			},
		},
		{
			name:     "WithHWAccelDevice",
			expected: []string{"-hwaccel_device", "sample"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithHWAccelDevice("sample")
			},
		},
		{
			name:     "WithChannel",
			expected: []string{"-channel", "sample"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithChannel("sample")
			},
		},
		{
			name:     "WithTVStandard",
			expected: []string{"-standard", "sample"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithTVStandard("sample")
			},
		},
		{
			name:     "WithVideoBitstreamFilters",
			expected: []string{"-vbsf", "sample"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithVideoBitstreamFilters("sample")
			},
		},
		{
			name:     "WithVideoPreset",
			expected: []string{"-vpre", "sample"},
			fn: func(c *ffmpeg.Command) *ffmpeg.Command {
				return c.WithVideoPreset("sample")
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
