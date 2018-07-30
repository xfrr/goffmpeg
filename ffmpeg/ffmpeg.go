package ffmpeg

import (
	"bytes"
	"github.com/xfrr/goffmpeg/utils"
	"os/exec"
	"strings"
)

type Configuration struct {
	FfmpegBin  string
	FfprobeBin string
}

func Configure() (Configuration, error) {
	var outFFmpeg bytes.Buffer
	var outProbe bytes.Buffer
	execFFmpegCommand := utils.GetFFmpegExec()
	execFFprobeCommand := utils.GetFFprobeExec()

	cmdFFmpeg := exec.Command(execFFmpegCommand[0], execFFmpegCommand[1])
	cmdProbe := exec.Command(execFFprobeCommand[0], execFFprobeCommand[1])

	cmdFFmpeg.Stdout = &outFFmpeg
	cmdProbe.Stdout = &outProbe

	err := cmdFFmpeg.Run()
	if err != nil {
		return Configuration{}, err
	}

	err = cmdProbe.Run()
	if err != nil {
		return Configuration{}, err
	}

	ffmpeg := strings.Replace(outFFmpeg.String(), utils.LineSeparator(), "", -1)
	fprobe := strings.Replace(outProbe.String(), utils.LineSeparator(), "", -1)

	cnf := Configuration{ffmpeg, fprobe}
	return cnf, nil
}
