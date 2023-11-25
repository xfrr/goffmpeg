package ffprobe_test

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/xfrr/goffmpeg/v2/ffprobe"
)

var (
	inputPath   = "../testdata/input.mp4"
	defaultArgs = []string{
		"-loglevel", "error",
		"-print_format", "json",
		"-show_format",
		"-show_streams",
		"-show_error",
	}
)

func Test_Ffprobe_NewCommand(t *testing.T) {
	tests := []struct {
		name         string
		command      *ffprobe.Command
		expectedArgs []string
	}{
		{
			name:         "default",
			command:      ffprobe.NewCommand(),
			expectedArgs: defaultArgs,
		},
		{
			name:    "with input path",
			command: ffprobe.NewCommand().WithInputPath(inputPath),
			expectedArgs: append(defaultArgs,
				"-i", inputPath,
			),
		},
		{
			name:    "with input reader",
			command: ffprobe.NewCommand().WithInputReader(bytes.NewBuffer(nil)),
			expectedArgs: append(defaultArgs,
				"-i", "pipe:0",
			),
		},
		{
			name:    "with output writer",
			command: ffprobe.NewCommand().WithOutputWriter(bytes.NewBuffer(nil)),
			expectedArgs: append(defaultArgs,
				"-o", "pipe:1",
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if len(test.command.Args()) != len(test.expectedArgs) {
				t.Errorf("expected %d args, got %d", len(test.expectedArgs), len(test.command.Args()))
			}

			for i, arg := range test.command.Args() {
				if arg != test.expectedArgs[i] {
					t.Errorf("expected arg %s, got %s", test.expectedArgs[i], arg)
				}
			}
		})
	}
}

func Test_Ffprobe_Command_Run(t *testing.T) {
	tests := []struct {
		name             string
		command          *ffprobe.Command
		withInputReader  bool
		withOutputWriter bool
		shouldFail       bool
	}{
		{
			name:       "default",
			command:    ffprobe.NewCommand(),
			shouldFail: true,
		},
		{
			name:    "with input path",
			command: ffprobe.NewCommand().WithInputPath(inputPath),
		},
		{
			name:            "with input reader",
			command:         ffprobe.NewCommand(),
			withInputReader: true,
		},
		{
			name:             "with input path and output writer",
			command:          ffprobe.NewCommand().WithInputPath(inputPath),
			withOutputWriter: true,
		},
		{
			name:             "with input reader and output writer",
			command:          ffprobe.NewCommand(),
			withInputReader:  true,
			withOutputWriter: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var err error

			if test.withInputReader {
				inputReader, err := os.Open(inputPath)
				if err != nil {
					t.Fatalf("expected no error, got %s", err)
				}
				defer inputReader.Close()

				test.command.WithInputReader(inputReader)
			}

			if test.withOutputWriter {
				outputWriter, err := os.Create(fmt.Sprintf("../testdata/%s.json", test.name))
				if err != nil {
					t.Fatalf("expected no error, got %s", err)
				}
				defer outputWriter.Close()

				test.command.WithOutputWriter(outputWriter)
			}

			file, err := test.command.Run(context.Background())
			if err != nil && !test.shouldFail {
				t.Fatalf("expected no error, got %s", err)
			} else if err == nil && test.shouldFail {
				t.Fatalf("expected error, got nil")
			} else if err != nil && test.shouldFail {
				return
			}

			if len(file.Streams) == 0 {
				t.Errorf("expected streams, got none")
			}

			if !test.withInputReader && file.Format.Filename != inputPath {
				t.Errorf("expected filename %s, got %s", inputPath, file.Format.Filename)
			}

			if test.withOutputWriter {
				if _, err := os.Stat(fmt.Sprintf("../testdata/%s.json", test.name)); os.IsNotExist(err) {
					t.Errorf("expected output file to exist, got %s", err)
				}
			}
		})
	}
}
