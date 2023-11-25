package duration

import (
	"fmt"
	"strconv"
	"strings"
)

// DurToSec converts duration string to seconds
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

// SecToDur converts seconds to duration string
func SecToDur(sec float64) (dur string) {
	hr := int(sec / (60 * 60))
	sec -= float64(hr) * (60 * 60)
	min := int(sec / (60))
	sec -= float64(min) * (60)
	if sec < 0 {
		sec = 0
	}

	format := "%02d:%02d:%02d"
	return fmt.Sprintf(format, hr, min, int(sec))
}
