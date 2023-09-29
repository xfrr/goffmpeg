package main

import (
	"fmt"

	"github.com/xfrr/goffmpeg/transcoder"
)

const (
	inputPath   = "../fixtures/input.3gp"
	outputPath  = "../test_results/hls-output.mp4"
	keyinfoPath = "keyinfo"
)

func main() {
	trans := new(transcoder.Transcoder)

	err := trans.Initialize(inputPath, outputPath)
	if err != nil {
		panic(err)
	}

	trans.MediaFile().SetVideoCodec("libx264")
	trans.MediaFile().SetHlsSegmentDuration(4)
	trans.MediaFile().SetEncryptionKey(keyinfoPath)

	done := trans.Run(true)
	progress := trans.Output()
	for p := range progress {
		fmt.Println(p)
	}

	fmt.Println(<-done)
}
