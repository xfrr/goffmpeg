package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/xfrr/goffmpeg/v2/ffmpeg"
	"github.com/xfrr/goffmpeg/v2/ffmpeg/progress"
)

var (
	// use your own rtsp stream
	defaultInputPath = "rtsp://localhost:8554/test"
	outputPath       = flag.String("o", "../results/rtsp.mp4", "output path")
)

func main() {
	flag.Parse()

	inputPath := flag.Arg(0)
	if inputPath == "" {
		inputPath = defaultInputPath
	}

	err := createOutputDir(*outputPath)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// the order of the arguments matters
	cmd := ffmpeg.NewCommand().
		WithInputPath(inputPath).
		WithOutputFormat("mp4").
		WithOutputPath(*outputPath)

	go func() {
		progress, err := cmd.Start(ctx)
		if err != nil {
			panic(err)
		}

		go func() {
			for p := range progress {
				printProgress(p)
			}
		}()
	}()

	err = cmd.Wait()
	if err != nil {
		panic(err)
	}
}

func printProgress(p progress.Progress) {
	fmt.Printf(`
Progress:
- Time Recorded: %s
- Frames Processed: %d
- Bitrate: %f
- Size: %d
- Speed: %f
- Fps: %f
- Dup: %d
- Drop: %d

`,
		p.Duration,
		p.FramesProcessed,
		p.Bitrate,
		p.Size,
		p.Speed,
		p.Fps,
		p.Dup,
		p.Drop,
	)
}

func createOutputDir(outputPath string) error {
	err := os.MkdirAll(filepath.Dir(outputPath), 0755)
	if err != nil {
		return err
	}

	return nil
}
