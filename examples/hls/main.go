package main

// Execute the following script to make it work:
/*
BASE_URL=${1:-'.'}
openssl rand 16 > file.key
echo $BASE_URL/file.key > file.keyinfo
echo file.key >> file.keyinfo
echo $(openssl rand -hex 16) >> file.keyinfo
*/

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
	keyinfoPath = flag.String("k", "file.keyinfo", "Encryption key path")
)

var (
	defaultInputPath = "../../testdata/input.mp4"
	outputPath       = flag.String("o", "../results/hls.m3u8", "output path")
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

	mediafile, err := ffprobe.NewCommand().
		WithInputPath(inputPath).
		Run(ctx)
	if err != nil {
		panic(err)
	}

	// the order of the arguments matters
	cmd := ffmpeg.NewCommand().
		WithInputPath(inputPath).
		WithHLSKeyInfoPath(*keyinfoPath).
		WithHLSSegmentTime(4).
		WithOutputPath(*outputPath)

	progress, err := cmd.Start(ctx)
	if err != nil {
		panic(err)
	}

	go func() {
		for msg := range progress {
			printProgress(msg, mediafile.GetDuration())
		}
	}()

	err = cmd.Wait()
	if err != nil {
		panic(err)
	}
}

func printProgress(p progress.Progress, tdur time.Duration) {
	// totalDurMs is the total duration in milliseconds
	tdur = time.Duration(tdur.Milliseconds()) * time.Millisecond

	// rdur is the remaining duration in milliseconds
	rdur := p.Duration - tdur

	// convert to positive if negative
	if rdur < 0 {
		rdur = -rdur
	}

	// remainingTimeStr is the remaining time in the format 00:00:00.000
	remainingTime := rdur.String()

	// progress is the percentage of the progress
	progress := int(float64(p.Duration.Milliseconds()) / float64(tdur.Milliseconds()) * 100)

	fmt.Printf(`
Progress:
- Current Time: %s
- Remaining Time: %s
- Progress: %d%s
- Frames Processed: %d
- Current Bitrate: %f
- Size: %d
- Speed: %f
- Fps: %f
- Dup: %d
- Drop: %d
`,
		p.Duration,
		remainingTime,
		progress, "%",
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
