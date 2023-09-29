package transcoder

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xfrr/goffmpeg/media"
)

func TestTranscoder(t *testing.T) {
	t.Run("#SetWhiteListProtocols", func(t *testing.T) {
		t.Run("Should not set -protocol_whitelist option if it isn't present", func(t *testing.T) {
			ts := Transcoder{}

			ts.SetMediaFile(&media.File{})
			require.NotEqual(t, ts.GetCommand()[0:2], []string{"-protocol_whitelist", "file,http,https,tcp,tls"})
			require.NotContains(t, ts.GetCommand(), "protocol_whitelist")
		})

		t.Run("Should set -protocol_whitelist option if it's present", func(t *testing.T) {
			ts := Transcoder{}

			ts.SetMediaFile(&media.File{})
			ts.SetWhiteListProtocols([]string{"file", "http", "https", "tcp", "tls"})

			require.Equal(t, ts.GetCommand()[0:2], []string{"-protocol_whitelist", "file,http,https,tcp,tls"})
		})
	})
}
