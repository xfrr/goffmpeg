package progress

import "time"

type Progress struct {
	framesProcessed int64
	bitrate         float64
	speed           float64
	size            int64
	fps             float64
	drop            int64
	dup             int64
	duration        time.Duration
}

func (p Progress) FramesProcessed() int64 {
	return p.framesProcessed
}

func (p Progress) Bitrate() float64 {
	return p.bitrate
}

func (p Progress) Speed() float64 {
	return p.speed
}

func (p Progress) Size() int64 {
	return p.size
}

func (p Progress) Fps() float64 {
	return p.fps
}

func (p Progress) Drop() int64 {
	return p.drop
}

func (p Progress) Dup() int64 {
	return p.dup
}

func (p Progress) Duration() time.Duration {
	return p.duration
}
