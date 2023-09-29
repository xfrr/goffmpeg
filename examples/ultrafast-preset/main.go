package main

import (
	"fmt"

	"github.com/xfrr/goffmpeg/transcoder"
)

const (
	inputPath  = "../fixtures/input.3gp"
	outputPath = "../test_results/ultrafast-output.mp4"
)

func main() {
	trans := new(transcoder.Transcoder)

	err := trans.Initialize(inputPath, outputPath)
	if err != nil {
		panic(err)
	}

	trans.MediaFile().SetPreset("ultrafast")

	done := trans.Run(true)
	progress := trans.Output()
	for p := range progress {
		fmt.Println(p)
	}

	fmt.Println(<-done)
}
