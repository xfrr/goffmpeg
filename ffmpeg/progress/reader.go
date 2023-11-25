package progress

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"regexp"
	"strings"
	"time"

	fferror "github.com/xfrr/goffmpeg/v2/ffmpeg/error"
)

var (
	cleanRegex = regexp.MustCompile(`\s+`)
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
			line := strings.ToLower(scanner.Text())
			if err := checkError(line); err != nil {
				return err
			}

			line = cleanLine(line)

			if p.framesProcessed == 0 {
				fp := parseInt64("frame", line)
				if fp != 0 {
					p.framesProcessed = fp
					continue
				}
			}

			if p.fps == 0 {
				fps := parseFloat64("fps", line)
				if fps != 0 {
					p.fps = fps
					continue
				}
			}

			if p.bitrate == 0 {
				bitrate := parseFloat64("bitrate", line)
				if bitrate != 0 {
					p.bitrate = bitrate
					continue
				}
			}

			if p.size == 0 {
				size := parseInt64("total_size", line)
				if size != 0 {
					p.size = size
					continue
				}
			}

			if p.duration == 0 {
				duration := parseDurationMs("out_time_ms", line)
				if duration != 0 {
					p.duration = duration
					continue
				}
			}

			if p.drop == 0 {
				drop := parseInt64("drop_frames", line)
				if drop != 0 {
					p.drop = drop
					continue
				}
			}

			if p.dup == 0 {
				dup := parseInt64("dup_frames", line)
				if dup != 0 {
					p.dup = dup
					continue
				}
			}

			if p.speed == 0 {
				speed := parseFloat64("speed", line)
				if speed != 0 {
					p.speed = speed
					continue
				}
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

// checkError checks if the line contains any of well known errors.
func checkError(line string) error {
	if strings.Contains(line, "unable to find a suitable output format") {
		return fferror.ErrUnableToFindOutputFormat
	}

	if strings.Contains(line, "no such file or directory") {
		return fferror.ErrFileNotFound
	}

	if strings.Contains(line, "invalid data found when processing input") {
		return fferror.ErrInvalidDataFoundWhenProcessingInput
	}

	if strings.Contains(line, "invalid argument") {
		return fferror.ErrInvalidArgument
	}

	return nil
}

func parseInt64(key, line string) int64 {
	var i int64
	_, err := fmt.Sscanf(line, fmt.Sprintf("%s=", key)+"%d", &i)
	if err != nil {
		return 0
	}

	return i
}

func parseFloat64(key, line string) float64 {
	var f float64
	_, err := fmt.Sscanf(line, fmt.Sprintf("%s=", key)+"%f", &f)
	if err != nil {
		return 0
	}

	return f
}

func parseDurationMs(key, line string) time.Duration {
	var i int64
	_, err := fmt.Sscanf(line, fmt.Sprintf("%s=", key)+"%d", &i)
	if err != nil {
		return 0
	}

	dms := time.Duration(i/1000) * time.Millisecond
	return dms
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

func cleanLine(line string) string {
	// Remove whitespace
	cleanedString := strings.ReplaceAll(line, " ", "")

	// Remove non-word characters (using regular expression)
	cleanedString = cleanRegex.ReplaceAllString(cleanedString, "")

	// Remove break lines and new lines
	cleanedString = strings.ReplaceAll(cleanedString, "\n", "")
	cleanedString = strings.ReplaceAll(cleanedString, "\r", "")

	return cleanedString
}
