package test

import (
	"os/exec"
	"testing"

	"github.com/muhammadharis/goffmpeg/transcoder"
)

func TestInputNotFound(t *testing.T) {

	var inputPath = "/data/testmkv"
	var outputPath = "/data/out/testmp4.mp4"

	trans := new(transcoder.Transcoder)

	err := trans.Initialize(inputPath, outputPath)
	if err != nil {
		return
	}
}

func TestTranscoding3GP(t *testing.T) {

	var inputPath = "/data/test3gp"
	var outputPath = "/data/testmp4.mp4"

	trans := new(transcoder.Transcoder)

	err := trans.Initialize(inputPath, outputPath)
	if err != nil {
		t.Error(err)
		return
	}

	done := trans.Run(false)
	err = <-done
	if err != nil {
		t.Error(err)
		return
	}
}

func TestTranscodingAVI(t *testing.T) {

	var inputPath = "/data/testavi"
	var outputPath = "/data/out/testmp4.mp4"

	trans := new(transcoder.Transcoder)

	err := trans.Initialize(inputPath, outputPath)
	if err != nil {
		t.Error(err)
		return
	}

	done := trans.Run(false)
	err = <-done
	if err != nil {
		t.Error(err)
		return
	}
}

func TestTranscodingFLV(t *testing.T) {

	var inputPath = "/data/testflv"
	var outputPath = "/data/out/testmp4.mp4"

	trans := new(transcoder.Transcoder)

	err := trans.Initialize(inputPath, outputPath)
	if err != nil {
		t.Error(err)
		return
	}

	done := trans.Run(false)
	err = <-done
	if err != nil {
		t.Error(err)
		return
	}
}

func TestTranscodingMKV(t *testing.T) {

	var inputPath = "/data/testmkv"
	var outputPath = "/data/out/testmp4.mp4"

	trans := new(transcoder.Transcoder)

	err := trans.Initialize(inputPath, outputPath)
	if err != nil {
		t.Error(err)
		return
	}

	done := trans.Run(false)
	err = <-done
	if err != nil {
		t.Error(err)
		return
	}
}

func TestTranscodingMOV(t *testing.T) {

	var inputPath = "/data/testmov"
	var outputPath = "/data/out/testmp4.mp4"

	trans := new(transcoder.Transcoder)

	err := trans.Initialize(inputPath, outputPath)
	if err != nil {
		t.Error(err)
		return
	}

	done := trans.Run(false)
	err = <-done
	if err != nil {
		t.Error(err)
		return
	}
}

func TestTranscodingMPEG(t *testing.T) {

	var inputPath = "/data/testmpeg"
	var outputPath = "/data/out/testmp4.mp4"

	trans := new(transcoder.Transcoder)

	err := trans.Initialize(inputPath, outputPath)
	if err != nil {
		t.Error(err)
		return
	}

	done := trans.Run(false)
	err = <-done
	if err != nil {
		t.Error(err)
		return
	}
}

func TestTranscodingOGG(t *testing.T) {

	var inputPath = "/data/testogg"
	var outputPath = "/data/out/testmp4.mp4"

	trans := new(transcoder.Transcoder)

	err := trans.Initialize(inputPath, outputPath)
	if err != nil {
		t.Error(err)
		return
	}

	done := trans.Run(false)
	err = <-done
	if err != nil {
		t.Error(err)
		return
	}
}

func TestTranscodingWAV(t *testing.T) {

	var inputPath = "/data/testwav"
	var outputPath = "/data/out/testmp4.mp4"

	trans := new(transcoder.Transcoder)

	err := trans.Initialize(inputPath, outputPath)
	if err != nil {
		t.Error(err)
		return
	}

	done := trans.Run(false)
	err = <-done
	if err != nil {
		t.Error(err)
		return
	}
}

func TestTranscodingWEBM(t *testing.T) {

	var inputPath = "/data/testwebm"
	var outputPath = "/data/out/testmp4.mp4"

	trans := new(transcoder.Transcoder)

	err := trans.Initialize(inputPath, outputPath)
	if err != nil {
		t.Error(err)
		return
	}

	done := trans.Run(false)
	err = <-done
	if err != nil {
		t.Error(err)
		return
	}
}

func TestTranscodingWMV(t *testing.T) {

	var inputPath = "/data/testwmv"
	var outputPath = "/data/out/testmp4.mp4"

	trans := new(transcoder.Transcoder)

	err := trans.Initialize(inputPath, outputPath)
	if err != nil {
		t.Error(err)
		return
	}

	done := trans.Run(false)
	err = <-done
	if err != nil {
		t.Error(err)
		return
	}
}

func TestTranscodingInputPipe(t *testing.T) {

	// Tests pipe with input mpeg, output mp4 using cat command for pipe-in
	var outputPath = "/data/out/testmp4.mp4"

	trans := new(transcoder.Transcoder)
	err := trans.InitializeEmptyTranscoder()
	trans.SetOutputPath(outputPath)
	trans.CreateInputPipe(exec.Command("cat", "/data/testmpeg"))

	if err != nil {
		t.Error(err)
		return
	}

	done := trans.Run(false)
	err = <-done
	if err != nil {
		t.Error(err)
		return
	}
}

func TestTranscodingProgress(t *testing.T) {

	var inputPath = "/data/testavi"
	var outputPath = "/data/out/testmp4.mp4"

	trans := new(transcoder.Transcoder)

	err := trans.Initialize(inputPath, outputPath)
	if err != nil {
		t.Error(err)
		return
	}

	done := trans.Run(true)
	for val := range trans.Output() {
		if &val != nil {
			break
		}
	}

	err = <-done
	if err != nil {
		t.Error(err)
		return
	}
}
