package ffmpeg

type Configuration struct {
    FfmpegBin string
    FfprobeBin string
}

func Configure(ffmpegBin, ffprobeBin string) *Configuration {
    return &Configuration{ffmpegBin, ffprobeBin}
}
