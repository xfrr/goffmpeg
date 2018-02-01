package utils

import (
	"strings"
	"strconv"
	"goffmpeg/models"
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



