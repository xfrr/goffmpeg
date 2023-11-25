package progress

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	fferror "github.com/xfrr/goffmpeg/v2/ffmpeg/error"
)

// Reader sets the interface for reading progress from ffmpeg.
type Reader interface {
	// Read reads from r until EOF or an error occurs and writes to progress channel.
	Read(ctx context.Context, r io.Reader, progress chan<- Progress) error
}

type reader struct {
	buffSize int
}

func NewReader() Reader {
	return &reader{
		buffSize: 1024,
	}
}

func (dr *reader) Read(ctx context.Context, r io.Reader, progress chan<- Progress) error {
	scanner := bufio.NewScanner(r)
	scanner.Split(splitFunc)

	buf := make([]byte, dr.buffSize)
	scanner.Buffer(buf, bufio.MaxScanTokenSize)

	p := &Progress{}

	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			line := scanner.Text()

			if err := checkError(line); err != nil {
				return err
			}

			if strings.HasPrefix(line, "frame=") {
				p.FramesProcessed = parseFrame(line)
				continue
			}

			if strings.HasPrefix(line, "fps=") {
				p.Fps = parseFps(line)
				continue
			}

			if strings.HasPrefix(line, "bitrate=") {
				p.Bitrate = parseBitrate(line)
				continue
			}

			if strings.HasPrefix(line, "total_size=") {
				p.Size = parseSize(line)
				continue
			}

			if strings.HasPrefix(line, "out_time_ms=") {
				p.Duration = parseOutTimeMs(line)
				continue
			}

			if strings.HasPrefix(line, "dup_frames=") {
				p.Dup = parseDupFrames(line)
				continue
			}

			if strings.HasPrefix(line, "drop_frames=") {
				p.Drop = parseDropFrames(line)
				continue
			}

			if strings.HasPrefix(line, "speed=") {
				p.Speed = parseSpeed(line)
				continue
			}

			if strings.Contains(line, "progress=continue") {
				progress <- *p
				p = &Progress{}
				continue
			}

			if strings.Contains(line, "progress=end") {
				progress <- *p
				return nil
			}
		}
	}

	return scanner.Err()
}

func checkError(line string) error {
	if strings.Contains(line, "Unable to find a suitable output format") {
		return fferror.ErrUnableToFindOutputFormat
	}

	if strings.Contains(line, "No such file or directory") {
		return fferror.ErrInputFileNotFound
	}

	return nil
}

func parseFrame(line string) int64 {
	var i int64
	_, err := fmt.Sscanf(line, "frame=%d", &i)
	if err != nil {
		return 0
	}

	return i
}

func parseFps(line string) float64 {
	var f float64
	_, err := fmt.Sscanf(line, "fps=%f", &f)
	if err != nil {
		return 0
	}

	return f
}

func parseBitrate(line string) float64 {
	var f float64
	_, err := fmt.Sscanf(line, "bitrate=%f", &f)
	if err != nil {
		return 0
	}

	return f
}

func parseSize(line string) int64 {
	var s int64
	_, err := fmt.Sscanf(line, "total_size=%d", &s)
	if err != nil {
		return 0
	}

	return s
}

func parseOutTimeMs(line string) time.Duration {
	var i int64
	_, err := fmt.Sscanf(line, "out_time_ms=%d", &i)
	if err != nil {
		return 0
	}

	return time.Duration(i/1000) * time.Millisecond
}

func parseDupFrames(line string) int32 {
	var i int32
	_, err := fmt.Sscanf(line, "dup_frames=%d", &i)
	if err != nil {
		return 0
	}

	return i
}

func parseDropFrames(line string) int32 {
	var i int32
	_, err := fmt.Sscanf(line, "drop_frames=%d", &i)
	if err != nil {
		return 0
	}

	return i
}

func parseSpeed(line string) float64 {
	var f float64
	_, err := fmt.Sscanf(line, "speed=%f", &f)
	if err != nil {
		return 0
	}

	return f
}

func splitFunc(data []byte, atEOF bool) (advance int, token []byte, spliterror error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	//windows \r\n
	//so  first \r and then \n can remove unexpected line break
	if i := bytes.IndexByte(data, '\r'); i >= 0 {
		// We have a cr terminated line
		return i + 1, data[0:i], nil
	}
	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, data[0:i], nil
	}
	if atEOF {
		return len(data), data, nil
	}

	return 0, nil, nil
}
