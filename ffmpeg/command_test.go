package ffmpeg_test

import (
	"context"
	"os"
	"os/exec"
	"path"
	"testing"

	"github.com/xfrr/goffmpeg/v2/ffmpeg"
	"github.com/xfrr/goffmpeg/v2/ffmpeg/progress"
)

const (
	fixturePath = "../testdata/"
	resultsPath = "./results"
)

var (
	inputPath = path.Join(fixturePath, "input.mp4")
)

func TestCommandWithProgress(t *testing.T) {
	createResultsDir(t)

	outputPath := path.Join(resultsPath, "progress.mpg")

	cmd := ffmpeg.NewCommand().
		WithInputPath(inputPath).
		WithOutputPath(outputPath)

	progressCh, err := cmd.Start(context.Background())
	if err != nil {
		panic(err)
	}

	progress := []progress.Progress{}
	for val := range progressCh {
		progress = append(progress, val)
	}

	err = cmd.Wait()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(progress) == 0 {
		t.Fatal("expected progress to be reported")
	}

	if !fileExists(outputPath) {
		t.Fatal("expected output file to exist")
	}
}

func TestCommandWithoutProgress(t *testing.T) {
	createResultsDir(t)

	outputPath := path.Join(resultsPath, "no-progress.mpg")

	cmd := ffmpeg.NewCommand().
		WithInputPath(inputPath).
		WithOutputPath(outputPath)

	err := cmd.Run(context.Background())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !fileExists(outputPath) {
		t.Fatal("expected output file to exist")
	}

}

func createResultsDir(t *testing.T) {
	err := exec.Command("mkdir", "-p", resultsPath).Run()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func fileExists(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	return !fileInfo.IsDir()
}
