package ffmpeg

import (
	"os/exec"
	"bytes"
	"strings"
)

type Configuration struct {
    FfmpegBin string
    FfprobeBin string
}

func Configure() (Configuration, error) {
	var outFFmpeg bytes.Buffer
	var outProbe bytes.Buffer

	cmdFFmpeg := exec.Command("/bin/sh", "-c", "which ffmpeg")
	cmdProbe := exec.Command("/bin/sh", "-c", "which ffprobe")

	cmdFFmpeg.Stdout = &outFFmpeg
	cmdProbe.Stdout = &outProbe

	err := cmdFFmpeg.Start()
	if err != nil {
		return Configuration{}, err
	}

	_, err = cmdFFmpeg.Process.Wait()
	if err != nil {
		return Configuration{}, err
	}

	err = cmdProbe.Start()
	if err != nil {
		return Configuration{}, err
	}

	_, err = cmdProbe.Process.Wait()
	if err != nil {
		return Configuration{}, err
	}

	ffmpeg := strings.Replace(outFFmpeg.String(), "\n", "", -1)
	fprobe := strings.Replace(outProbe.String(), "\n", "", -1)

	cnf := Configuration{ffmpeg, fprobe}

	return cnf, nil
}
