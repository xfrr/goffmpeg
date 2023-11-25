package progress_test

import (
	"context"
	"strings"
	"testing"
	"time"

	fferror "github.com/xfrr/goffmpeg/v2/ffmpeg/error"
	"github.com/xfrr/goffmpeg/v2/ffmpeg/progress"
)

func TestReader_Read(t *testing.T) {
	ctx := context.Background()

	// Create a mock reader with test data
	mockReader := strings.NewReader(`
	frame=100
			fps=25
	bitrate=2000
			total_size=102400
	out_time_ms=5000
	  dup_frames=5
	drop_frames=2
			speed=1.5
	progress=continue
	  frame=200
	fps=30
	bitrate=2500
			total_size=204800
	out_time_ms=1000		
	  dup_frames=10
	drop_frames=4
		speed=2.0
	progress=end`)

	// Create a progress channel to receive progress updates
	progressCh := make(chan progress.Progress)

	// Create a reader instance
	reader := progress.NewReader()

	// Start reading from the mock reader
	go func() {
		defer close(progressCh)
		err := reader.Read(ctx, mockReader, progressCh)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	}()

	// Verify the progress updates
	expectedProgress := []progress.Progress{
		{
			FramesProcessed: 100,
			Fps:             25,
			Bitrate:         2000,
			Size:            102400,
			Duration:        time.Duration(5) * time.Millisecond,
			Dup:             5,
			Drop:            2,
			Speed:           1.5,
		},
		{
			FramesProcessed: 200,
			Fps:             30,
			Bitrate:         2500,
			Size:            204800,
			Duration:        time.Duration(1) * time.Millisecond,
			Dup:             10,
			Drop:            4,
			Speed:           2.0,
		},
	}

	idx := 0
	for p := range progressCh {
		expected := expectedProgress[idx]
		if p.FramesProcessed != expected.FramesProcessed {
			t.Errorf("expected frames processed %d, got %d", expected.FramesProcessed, p.FramesProcessed)
		}

		if p.Fps != expected.Fps {
			t.Errorf("expected fps %f, got %f", expected.Fps, p.Fps)
		}

		if p.Bitrate != expected.Bitrate {
			t.Errorf("expected bitrate %f, got %f", expected.Bitrate, p.Bitrate)
		}

		if p.Size != expected.Size {
			t.Errorf("expected size %d, got %d", expected.Size, p.Size)
		}

		if p.Duration.String() != expected.Duration.String() {
			t.Errorf("expected duration %s, got %s", expected.Duration, p.Duration)
		}

		if p.Dup != expected.Dup {
			t.Errorf("expected dup frames %d, got %d", expected.Dup, p.Dup)
		}

		if p.Drop != expected.Drop {
			t.Errorf("expected drop frames %d, got %d", expected.Drop, p.Drop)
		}

		if p.Speed != expected.Speed {
			t.Errorf("expected speed %f, got %f", expected.Speed, p.Speed)
		}

		idx++
	}

	if idx != len(expectedProgress) {
		t.Errorf("expected %d progress updates, got %d", len(expectedProgress), idx)
	}
}

func TestReader_ReadWithError(t *testing.T) {
	ctx := context.Background()

	cases := []struct {
		name  string
		input string
		err   error
	}{
		{
			name:  "unable to find a suitable output format",
			input: "[NULL @ 0x7f8b9b800000] Unable to find a suitable output format for 'pipe:'",
			err:   fferror.ErrUnableToFindOutputFormat,
		},
		{
			name:  "pipe:: Invalid data found when processing input",
			input: "[pipe @ 0x7f8b9b800000] Invalid data found when processing input",
			err:   fferror.ErrInvalidDataFoundWhenProcessingInput,
		},
		{
			name:  "pipe:: Invalid argument",
			input: "[pipe @ 0x7f8b9b800000] Invalid argument",
			err:   fferror.ErrInvalidArgument,
		},
		{
			name:  "no such file or directory",
			input: "No such file or directory",
			err:   fferror.ErrFileNotFound,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mockReader := strings.NewReader(c.input)
			progressCh := make(chan progress.Progress)
			defer close(progressCh)

			reader := progress.NewReader()
			err := reader.Read(ctx, mockReader, progressCh)
			if err != c.err {
				t.Errorf("expected error %v, got %v", c.err, err)
			}
		})
	}
}

func TestReader_ReadContextDone(t *testing.T) {
	mockReader := strings.NewReader(`
	frame=100
			fps=25
	bitrate=2000
			total_size=102400
	out_time_ms=5000
	  dup_frames=5
	drop_frames=2
			speed=1.5
	progress=continue
	  frame=200
	fps=30
	bitrate=2500
			total_size=204800
	out_time_ms=1000		
	  dup_frames=10
	drop_frames=4
		speed=2.0
	progress=end`)

	progressCh := make(chan progress.Progress)
	defer close(progressCh)

	reader := progress.NewReader()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := reader.Read(ctx, mockReader, progressCh)
	if err != context.Canceled {
		t.Errorf("expected error %v, got %v", context.Canceled, err)
	}
}
