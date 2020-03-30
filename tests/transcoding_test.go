package test

import (
	"io/ioutil"
	"os/exec"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/xfrr/goffmpeg/transcoder"
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

func TestTranscodePipes(t *testing.T) {
	c1 := exec.Command("cat", "/tmp/data/testmkv")

	trans := new(transcoder.Transcoder)

	err := trans.InitializeEmptyTranscoder()
	assert.Nil(t, err)

	w, err := trans.CreateInputPipe()
	assert.Nil(t, err)
	c1.Stdout = w

	r, err := trans.CreateOutputPipe("mp4")
	assert.Nil(t, err)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		_, err := ioutil.ReadAll(r)
		assert.Nil(t, err)

		r.Close()
		wg.Done()
	}()

	go func() {
		err := c1.Run()
		assert.Nil(t, err)
		w.Close()
	}()
	done := trans.Run(false)
	err = <-done
	assert.Nil(t, err)

	wg.Wait()
}
