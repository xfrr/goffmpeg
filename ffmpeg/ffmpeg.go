package ffmpeg

import (
	"os/exec"
	"bytes"
	"strings"
	"goffmpeg/utils"
)

type Configuration struct {
	FfmpegBin  string
    FfprobeBin string
	ExecCmd    string
	ExecArgs   string
}

func Configure() (Configuration, error) {
	var outFFmpeg bytes.Buffer
	var outProbe bytes.Buffer

	execCmd := utils.GetExec()
	execFFmpegCommand := utils.GetFFmpegExec()
	execFFprobeCommand := utils.GetFFprobeExec()
	execArgs := utils.GetExecArgs()

	cmdFFmpeg := exec.Command(execCmd, execArgs, execFFmpegCommand)
	cmdProbe := exec.Command(execCmd, execArgs, execFFprobeCommand)

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

	cnf := Configuration{ffmpeg, fprobe, execCmd, execArgs}

	return cnf, nil
}
