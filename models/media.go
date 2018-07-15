package models

import (
  "strings"
  "strconv"
  "fmt"
  "reflect"
)

type Mediafile struct {
    aspect                string
    resolution            string
    videoBitRate          int
    videoBitRateTolerance int
    videoMaxBitRate       int
    videoMinBitrate       int
    videoCodec            string
    frameRate             int
    audioRate             int
    maxKeyframe           int
    minKeyframe           int
    keyframeInterval      int
    audioCodec            string
    audioBitrate          int
    audioChannels         int
    bufferSize            int
    threads               int
    preset                string
    tune                  string
    audioProfile          string
    videoProfile          string
    target                string
    duration              string
    durationInput         string
    seekTime              string
    quality               int
    strict                int
    muxDelay              string
    seekUsingTsInput      bool
    seekTimeInput         string
    inputPath             string
    outputPath            string
    outputFormat          string
    copyTs                bool
    nativeFramerateInput  bool
    inputInitialOffset    string
    rtmpLive              string
    hlsPlaylistType       string
    hlsSegmentDuration    int
    httpMethod            string
    httpKeepAlive         bool
    streamIds             map[int]string
    metadata              Metadata
    filter                string
}

/*** SETTERS ***/
func (m *Mediafile) SetFilter(v string){
  m.filter = v
}

func (m *Mediafile) SetAspect(v string) {
  m.aspect = v
}

func (m *Mediafile) SetResolution(v string) {
  m.resolution = v
}

func (m *Mediafile) SetVideoBitRate(v int) {
  m.videoBitRate = v
}

func (m *Mediafile) SetVideoBitRateTolerance(v int) {
  m.videoBitRateTolerance = v
}

func (m *Mediafile) SetVideoMaxBitrate(v int) {
  m.videoMaxBitRate = v
}

func (m *Mediafile) SetVideoMinBitRate(v int) {
  m.videoMinBitrate = v
}

func (m *Mediafile) SetVideoCodec(v string) {
  m.videoCodec = v
}

func (m *Mediafile) SetFrameRate(v int) {
  m.frameRate = v
}

func (m *Mediafile) SetAudioRate(v int) {
  m.audioRate = v
}

func (m *Mediafile) SetMaxKeyFrame(v int) {
  m.maxKeyframe = v
}

func (m *Mediafile) SetMinKeyFrame(v int) {
  m.minKeyframe = v
}

func (m *Mediafile) SetKeyframeInterval(v int) {
  m.keyframeInterval = v
}

func (m *Mediafile) SetAudioCodec(v string) {
  m.audioCodec = v
}

func (m *Mediafile) SetAudioBitRate(v int) {
  m.audioBitrate = v
}

func (m *Mediafile) SetAudioChannels(v int) {
  m.audioChannels = v
}

func (m *Mediafile) SetBufferSize(v int) {
  m.bufferSize = v
}

func (m *Mediafile) SetThreads(v int) {
  m.threads = v
}

func (m *Mediafile) SetPreset(v string) {
  m.preset = v
}

func (m *Mediafile) SetTune(v string) {
  m.tune = v
}

func (m *Mediafile) SetAudioProfile(v string) {
  m.audioProfile = v
}

func (m *Mediafile) SetVideoProfile(v string) {
  m.videoProfile = v
}

func (m *Mediafile) SetDuration(v string) {
  m.duration = v
}

func (m *Mediafile) SetDurationInput(v string) {
  m.durationInput = v
}

func (m *Mediafile) SetSeekTime(v string) {
  m.seekTime = v
}

func (m *Mediafile) SetSeekTimeInput(v string) {
  m.seekTimeInput = v
}

func (m *Mediafile) SetQuality(v int) {
  m.quality = v
}

func (m *Mediafile) SetStrict(v int) {
	m.strict = v
}

func (m *Mediafile) SetSeekUsingTsInput(val bool) {
  m.seekUsingTsInput = val
}

func (m *Mediafile) SetCopyTs(val bool) {
  m.copyTs = val
}

func (m *Mediafile) SetInputPath(val string) {
  m.inputPath = val
}

func (m *Mediafile) SetMuxDelay(val string) {
  m.muxDelay = val
}

func (m *Mediafile) SetOutputPath(val string) {
  m.outputPath = val
}

func (m *Mediafile) SetOutputFormat(val string) {
  m.outputFormat = val
}

func (m *Mediafile) SetNativeFramerateInput(val bool) {
  m.nativeFramerateInput = val
}

func (m *Mediafile) SetRtmpLive(val string) {
  m.rtmpLive = val
}

func (m *Mediafile) SetHlsSegmentDuration(val int) {
  m.hlsSegmentDuration = val
}

func (m *Mediafile) SetHlsPlaylistType(val string) {
  m.hlsPlaylistType = val
}

func (m *Mediafile) SetHttpMethod(val string) {
  m.httpMethod = val
}

func (m *Mediafile) SetHttpKeepAlive(val bool) {
  m.httpKeepAlive = val
}

func (m *Mediafile) SetInputInitialOffset(val string) {
  m.inputInitialOffset = val
}

func (m *Mediafile) SetStreamIds(val map[int]string) {
  m.streamIds = val
}

func (m *Mediafile) SetMetadata(v Metadata) {
  m.metadata = v
}

/*** GETTERS ***/

func (m *Mediafile) Filter() string {
  return m.filter
}

func (m *Mediafile) Aspect() string {
  return m.aspect
}

func (m *Mediafile) Resolution() string {
  return m.resolution
}

func (m *Mediafile) VideoBitrate() int {
  return m.videoBitRate
}

func (m *Mediafile) VideoBitRateTolerance() int {
  return m.videoBitRateTolerance
}

func (m *Mediafile) VideoMaxBitRate() int {
  return m.videoMaxBitRate
}

func (m *Mediafile) VideoMinBitRate() int {
  return m.videoMinBitrate
}

func (m *Mediafile) VideoCodec() string {
  return m.videoCodec
}

func (m *Mediafile) FrameRate() int {
  return m.frameRate
}

func (m *Mediafile) AudioRate() int {
  return m.audioRate
}

func (m *Mediafile) MaxKeyFrame() int {
  return m.maxKeyframe
}

func (m *Mediafile) MinKeyFrame() int {
  return m.minKeyframe
}

func (m *Mediafile) KeyFrameInterval() int {
  return m.keyframeInterval
}

func (m *Mediafile) AudioCodec() string {
  return m.audioCodec
}

func (m *Mediafile) AudioBitrate() int {
  return m.audioBitrate
}

func (m *Mediafile) AudioChannels() int {
  return m.audioChannels
}

func (m *Mediafile) BufferSize() int {
  return m.bufferSize
}

func (m *Mediafile) Threads() int {
  return m.threads
}

func (m *Mediafile) Target() string {
  return m.target
}

func (m *Mediafile) Duration() string {
  return m.duration
}

func (m *Mediafile) DurationInput() string {
  return m.durationInput
}

func (m *Mediafile) SeekTime() string {
  return m.seekTime
}

func (m *Mediafile) Preset() string {
  return m.preset
}

func (m *Mediafile) AudioProfile() string {
  return m.audioProfile
}

func (m *Mediafile) VideoProfile() string {
  return m.videoProfile
}

func (m *Mediafile) Tune() string {
  return m.tune
}

func (m *Mediafile) SeekTimeInput() string {
  return m.seekTimeInput
}

func (m *Mediafile) Quality() int {
  return m.quality
}

func (m *Mediafile) Strict() int {
  return m.strict
}

func (m *Mediafile) MuxDelay() string {
  return m.muxDelay
}

func (m *Mediafile) SeekUsingTsInput() bool {
  return m.seekUsingTsInput
}

func (m *Mediafile) CopyTs() bool {
  return m.copyTs
}

func (m *Mediafile) InputPath() string {
  return m.inputPath
}

func (m *Mediafile) OutputPath() string {
  return m.outputPath
}

func (m *Mediafile) OutputFormat() string {
  return m.outputFormat
}

func (m *Mediafile) NativeFramerateInput() bool {
  return m.nativeFramerateInput
}

func (m *Mediafile) RtmpLive() string {
  return m.rtmpLive
}

func (m *Mediafile) HlsSegmentDuration() int {
  return m.hlsSegmentDuration
}

func (m *Mediafile) HlsPlaylistType() string {
  return m.hlsPlaylistType
}

func (m *Mediafile) InputInitialOffset() string {
  return m.inputInitialOffset
}

func (m *Mediafile) HttpMethod() string {
  return m.httpMethod
}

func (m *Mediafile) HttpKeepAlive() bool {
  return m.httpKeepAlive
}

func (m *Mediafile) StreamIds() map[int]string {
  return m.streamIds
}

func (m *Mediafile) Metadata() Metadata {
  return m.metadata
}

/** OPTS **/
func (m *Mediafile) ToStrCommand() []string {
  var strCommand []string

  opts := []string {
    "SeekTimeInput",
    "SeekUsingTsInput",
    "NativeFramerateInput",
    "DurationInput",
    "RtmpLive",
    "InputInitialOffset",
    "InputPath",

    "Aspect",
    "Resolution",
    "FrameRate",
    "AudioRate",
    "VideoCodec",
    "VideoBitRate",
    "VideoBitRateTolerance",
    "VideoMaxBitRate",
    "VideoMinBitRate",
    "VideoProfile",
    "AudioCodec",
    "AudioBitRate",
    "AudioChannels",
    "AudioProfile",
    "Quality",
    "Strict",
    "BufferSize",
    "MuxDelay",
    "Threads",
    "KeyframeInterval",
    "Preset",
    "Tune",
    "Target",
    "SeekTime",
    "Duration",
    "CopyTs",
    "StreamIds",
    "OutputFormat",
    "HlsSegmentDuration",
    "HlsPlaylistType",
    "Filter",
    "HttpMethod",
    "HttpKeepAlive",
    "OutputPath",
  }
  for _, name := range opts {
    opt :=  reflect.ValueOf(m).MethodByName(fmt.Sprintf("Obtain%s", name))
    if (opt != reflect.Value{}) {
      result := opt.Call([]reflect.Value{})

      if val, ok := result[0].Interface().([]string); ok {
        strCommand = append(strCommand, val...)
      }
    }
  }

  return strCommand
}

func (m *Mediafile) ObtainFilter() []string{
  if m.filter != "" {
    return []string{"-vf", m.filter}
  }
  return nil
}

func (m *Mediafile) ObtainAspect() []string {
  // Set aspect
  if m.resolution != "" {
    resolution := strings.Split(m.resolution, "x")
    if len(resolution) != 0 {
      width, _ := strconv.ParseFloat(resolution[0], 64)
      height, _ := strconv.ParseFloat(resolution[1], 64)
      return []string{"-aspect", fmt.Sprintf("%f", width/height)}
    }
  }

  if m.aspect != "" {
    return []string{"-aspect", m.aspect}
  }
	return nil
}

func (m *Mediafile) ObtainInputPath() []string {
  return []string{"-i", m.inputPath}
}

func (m *Mediafile) ObtainNativeFramerateInput() []string {
  if m.nativeFramerateInput {
    return []string{"-re"}
  }
  return nil
}

func (m *Mediafile) ObtainOutputPath() []string {
  return []string{m.outputPath}
}

func (m *Mediafile) ObtainVideoCodec() []string {
  if m.videoCodec != "" {
    return []string{"-c:v", m.videoCodec}
  }
  return nil
}

func (m *Mediafile) ObtainFrameRate() []string {
  if m.frameRate != 0 {
    return []string{"-r",fmt.Sprintf("%d", m.frameRate)}
  }
  return nil
}

func (m *Mediafile) ObtainAudioRate() []string {
  if m.audioRate != 0 {
    return []string{"-ar",fmt.Sprintf("%d", m.audioRate)}
  }
  return nil
}

func (m *Mediafile) ObtainResolution() []string {
  if m.resolution != "" {
    return []string{"-s", m.resolution}
  }
  return nil
}

func (m *Mediafile) ObtainVideoBitRate() []string {
  if m.videoBitRate != 0 {
    return []string{"-b:v", fmt.Sprintf("%d", m.videoBitRate)}
  }
  return nil
}

func (m *Mediafile) ObtainAudioCodec() []string {
  if m.audioCodec != "" {
    return []string{"-c:a", m.audioCodec}
  }
  return nil
}

func (m *Mediafile) ObtainAudioBitRate() []string {
  if m.audioBitrate != 0 {
    return []string{"-b:a", fmt.Sprintf("%d", m.audioBitrate)}
  }
  return nil
}

func (m *Mediafile) ObtainAudioChannels() []string {
  if m.audioChannels != 0 {
    return []string{"-ac",fmt.Sprintf("%d", m.audioChannels)}
  }
  return nil
}

func (m *Mediafile) ObtainVideoMaxBitRate() []string {
  if m.videoMaxBitRate != 0 {
    return []string{"-maxrate",fmt.Sprintf("%dk", m.videoMaxBitRate)}
  }
  return nil
}

func (m *Mediafile) ObtainVideoMinBitRate() []string {
  if m.videoMinBitrate != 0 {
    return []string{"-minrate",fmt.Sprintf("%dk", m.videoMinBitrate)}
  }
  return nil
}

func (m *Mediafile) ObtainBufferSize() []string {
  if m.bufferSize != 0 {
    return []string{"-bufsize",fmt.Sprintf("%dk", m.bufferSize)}
  }
  return nil
}

func (m *Mediafile) ObtainVideoBitRateTolerance() []string {
  if m.videoBitRateTolerance != 0 {
    return []string{"-bt",fmt.Sprintf("%dk", m.videoBitRateTolerance)}
  }
  return nil
}

func (m *Mediafile) ObtainThreads() []string {
  if m.threads != 0 {
    return []string{"-threads",fmt.Sprintf("%d", m.threads)}
  }
  return nil
}

func (m *Mediafile) ObtainTarget() []string {
  if m.target != "" {
    return []string{"-target",m.target}
  }
  return nil
}

func (m *Mediafile) ObtainDuration() []string {
  if m.duration != "" {
    return []string{"-t",m.duration}
  }
  return nil
}

func (m *Mediafile) ObtainDurationInput() []string {
  if m.durationInput != "" {
    return []string{"-t",m.durationInput}
  }
  return nil
}

func (m *Mediafile) ObtainKeyframeInterval() []string {
  if m.keyframeInterval != 0 {
    return []string{"-g",fmt.Sprintf("%d", m.keyframeInterval)}
  }
  return nil
}

func (m *Mediafile) ObtainSeekTime() []string {
  if m.seekTime != "" {
    return []string{"-ss",m.seekTime}
  }
  return nil
}

func (m *Mediafile) ObtainSeekTimeInput() []string {
  if m.seekTimeInput != "" {
    return []string{"-ss",m.seekTimeInput}
  }
  return nil
}

func (m *Mediafile) ObtainPreset() []string {
  if m.preset != "" {
    return []string{"-preset",m.preset}
  }
  return nil
}

func (m *Mediafile) ObtainTune() []string {
  if m.tune != "" {
    return []string{"-tune",m.tune}
  }
  return nil
}

func (m *Mediafile) ObtainQuality() []string {
  if m.quality != 0 {
    return []string{"-crf",fmt.Sprintf("%d", m.quality)}
  }
  return nil
}

func (m *Mediafile) ObtainStrict() []string {
  if m.bufferSize != 0 {
    return []string{"-strict",fmt.Sprintf("%d", m.strict)}
  }
  return nil
}

func (m *Mediafile) ObtainVideoProfile() []string {
  if m.videoProfile != "" {
    return []string{"-profile:v",m.videoProfile}
  }
  return nil
}

func (m *Mediafile) ObtainAudioProfile() []string {
  if m.audioProfile != "" {
    return []string{"-profile:a", m.audioProfile}
  }
  return nil
}

func (m *Mediafile) ObtainCopyTs() []string {
  if m.copyTs {
    return []string{"-copyts"}
  }
  return nil
}

func (m *Mediafile) ObtainOutputFormat() []string {
  if m.outputFormat != "" {
    return []string{"-f",m.outputFormat}
  }
  return nil
}

func (m *Mediafile) ObtainMuxDelay() []string {
  if m.muxDelay != "" {
    return []string{"-muxdelay",m.muxDelay}
  }
  return nil
}

func (m *Mediafile) ObtainSeekUsingTsInput() []string {
  if m.seekUsingTsInput {
    return []string{"-seek_timestamp", "1"}
  }
	return nil
}

func (m *Mediafile) ObtainRtmpLive() []string {
  if m.rtmpLive != "" {
    return []string{"-rtmp_live",m.rtmpLive}
  } else {
    return nil
  }
}

func (m *Mediafile) ObtainHlsPlaylistType() []string {
  if m.hlsPlaylistType != "" {
    return []string{"-hls_playlist_type",m.hlsPlaylistType}
  } else {
    return nil
  }
}

func (m *Mediafile) ObtainInputInitialOffset() []string {
  if m.inputInitialOffset != "" {
    return []string{"-itsoffset",m.inputInitialOffset}
  } else {
    return nil
  }
}

func (m *Mediafile) ObtainHlsSegmentDuration() []string {
  if m.hlsSegmentDuration != 0 {
    return []string{"-hls_time",fmt.Sprintf("%d", m.hlsSegmentDuration)}
  } else {
    return nil
  }
}

func (m *Mediafile) ObtainHttpMethod() []string {
  if m.httpMethod != "" {
    return []string{"-method",m.httpMethod}
  } else {
    return nil
  }
}

func (m *Mediafile) ObtainHttpKeepAlive() []string {
  if m.httpKeepAlive {
    return []string{"-multiple_requests","1"}
  } else {
    return nil
  }
}

func (m *Mediafile) ObtainStreamIds() []string {
  if m.streamIds != nil && len(m.streamIds) != 0 {
    result := []string{}
    for i, val := range m.streamIds {
      result = append(result, []string{"-streamid", fmt.Sprintf("%d:%s", i, val)}...)
    }
    return result
  }
  return nil
}
