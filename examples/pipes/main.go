package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/xfrr/goffmpeg/v2/ffmpeg"
	"github.com/xfrr/goffmpeg/v2/ffmpeg/progress"
	"github.com/xfrr/goffmpeg/v2/ffprobe"
)

var (
	defaultInputPath = "../../testdata/input.mp4"
	outputPath       = flag.String("o", "../results/pipes.mp4", "output path")
)

func main() {
	flag.Parse()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	inputPath := flag.Arg(0)
	if inputPath == "" {
		inputPath = defaultInputPath
	}

	inputFile, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}

	err = createOutputDir(*outputPath)
	if err != nil {
		panic(err)
	}

	outputFile, err := os.Create(*outputPath)
	if err != nil {
		panic(err)
	}

	// the order of the arguments matters
	cmd := ffmpeg.NewCommand().
		WithInputPipe(inputFile).
		WithOutputFormat("mpeg").
		WithOutputPipe(outputFile)

	go func() {
		progress, err := cmd.Start(ctx)
		if err != nil {
			panic(err)
		}

		mediafile, err := ffprobe.NewCommand().
			WithInputPath(inputPath).
			Run(ctx)
		if err != nil {
			panic(err)
		}

		go func() {
			for msg := range progress {
				printProgress(msg, mediafile.Duration())
			}
		}()
	}()

	err = cmd.Wait()
	if err != nil {
		panic(err)
	}
}

func printProgress(p progress.Progress, tdur time.Duration) {
	// calculate the progress percentage
	progressPct := fmt.Sprintf("%d%s", p.Duration().Milliseconds()/tdur.Milliseconds()*100, "%")

	// convert the total duration to milliseconds
	totalDuration := time.Duration(tdur.Milliseconds()) * time.Millisecond

	// calculate the remaining duration
	remainingDuration := p.Duration() - totalDuration
	if remainingDuration < 0 {
		remainingDuration = -remainingDuration
	}

	fmt.Printf(`
Progress:
- Progress: %s
- Time Remaining: %s
- Time Processed: %s
- Frames Processed: %d
- Current Bitrate: %f
- Size: %d
- Speed: %f
- Fps: %f
- Dup: %d
- Drop: %d
`,
		progressPct,
		remainingDuration,
		p.Duration(),
		p.FramesProcessed(),
		p.Bitrate(),
		p.Size(),
		p.Speed(),
		p.Fps(),
		p.Dup(),
		p.Drop(),
	)
}

func createOutputDir(outputPath string) error {
	err := os.MkdirAll(filepath.Dir(outputPath), 0755)
	if err != nil {
		return err
	}

	return nil
}
