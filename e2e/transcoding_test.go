package test

import (
	"os/exec"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/xfrr/goffmpeg/transcoder"
)

const (
	fixturePath = "./fixtures"
	resultsPath = "./test_results"
)

var (
	// Input files
	input3gp = path.Join(fixturePath, "input.3gp")
)

func TestInputNotFound(t *testing.T) {
	createResultsDir(t)
	var outputPath = path.Join(resultsPath, "notfound.mp4")

	trans := new(transcoder.Transcoder)

	err := trans.Initialize("notfound.3gp", outputPath)
	assert.NotNil(t, err)
}

func TestTranscodingProgress(t *testing.T) {
	createResultsDir(t)

	outputPath := path.Join(resultsPath, "progress.mp4")
	trans := new(transcoder.Transcoder)

	err := trans.Initialize(input3gp, outputPath)
	assert.Nil(t, err)

	errCh := trans.Run(true)

	progress := []transcoder.Progress{}
	for val := range trans.Output() {
		progress = append(progress, val)
	}
	err = <-errCh
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(progress), 1)
	checkFileExists(t, outputPath)
}

func createResultsDir(t *testing.T) {
	err := exec.Command("mkdir", "-p", resultsPath).Run()
	assert.Nil(t, err)
}

func checkFileExists(t *testing.T, filepath string) {
	res, err := exec.Command("cat", filepath).Output()
	assert.Nil(t, err)
	assert.Greater(t, len(res), 0)
}
