package goffmpeg

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigure(t *testing.T) {
	t.Run("Should set the correct command", func(t *testing.T) {
		ctx := context.Background()
		cfg, err := Configure(ctx)
		assert.Nil(t, err)
		assert.NotEmpty(t, cfg.FFmpegBinPath())
		assert.NotEmpty(t, cfg.FFprobeBinPath())
	})
}
