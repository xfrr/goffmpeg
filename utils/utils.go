package utils

import (
	"strings"
	"strconv"
	"github.com/xfrr/goffmpeg/models"
	"runtime"
)

func DurToSec(dur string) (sec float64) {
	durAry := strings.Split(dur, ":")
	var secs float64
	if len(durAry) != 3 {
		return
	}
	hr, _ := strconv.ParseFloat(durAry[0], 64)
	secs = hr * (60 * 60)
	min, _ := strconv.ParseFloat(durAry[1], 64)
	secs += min * (60)
	second, _ := strconv.ParseFloat(durAry[2], 64)
	secs += second
	return secs
}

func GetExec() string {
	var platform = runtime.GOOS
	var command = ""

	switch platform {
	case "windows":
		command = "cmd"
		break
	default:
		command = "/bin/sh"
		break
	}

	return command
}

func GetExecArgs() string {
	var platform = runtime.GOOS
	var args = ""

	switch platform {
	case "windows":
		args = "/C"
		break
	default:
		args = "-c"
		break
	}

	return args
}

func GetFFmpegExec() string {
	var platform = runtime.GOOS
	var command = ""

	switch platform {
	case "windows":
		command = "where ffmpeg"
		break
	default:
		command = "which ffmpeg"
		break
	}

	return command
}

func GetFFprobeExec() string {
	var platform = runtime.GOOS
	var command = ""

	switch platform {
	case "windows":
		command = "where ffprobe"
		break
	default:
		command = "which ffprobe"
		break
	}

	return command
}

func CheckFileType(streams []models.Streams) string {
	for i := 0; i < len(streams); i++ {
		st := streams[i]
		if st.CodecType == "video" {
			return "video"
			break
		}
	}

	return "audio"
}



