package models

type Ffmpeg struct {
    FfmpegBinPath         string
    FfprobeBinPath        string
}

type Mediafile struct {
    Aspect                string
    Resolution            string
    VideoBitRate          string
    VideoMaxBitrate       string
    VideoMinBitrate       string
    VideoCodec            string
    FrameRate             string
    MaxKeyframe           string
    MinKeyframe           string
    AudioCodec            string
    AudioBitrate          string
    AudioSampleRate       string
    AudioChannels         string
    BufferSize            string
    Threads               string
    Target                string
    Duration              string
    SeekTime              string
    Quality               string
}
