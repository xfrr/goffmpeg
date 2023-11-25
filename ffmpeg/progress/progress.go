package progress

import "time"

type Progress struct {
	FramesProcessed int64
	Bitrate         float64
	Speed           float64
	Size            int64
	Fps             float64
	Drop            int32
	Dup             int32
	Duration        time.Duration
}
