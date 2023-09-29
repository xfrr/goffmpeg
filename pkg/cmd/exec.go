package cmd

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
)

func FindBinPath(ctx context.Context, command string) (string, error) {
	if command == "" {
		return "", fmt.Errorf("command cannot be empty")
	}

	path, err := execBufferOutput(ctx, getFindCommand(), command)
	if err != nil {
		return "", err
	}

	return path, nil
}

func execBufferOutput(ctx context.Context, command string, args ...string) (string, error) {
	var out bytes.Buffer

	c := exec.CommandContext(ctx, command, args...)
	c.Stdout = &out

	err := c.Run()
	if err != nil {
		return "", fmt.Errorf("%s: %s", c.String(), err)
	}

	return out.String(), nil
}
