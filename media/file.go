package media

import (
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
)

type File struct {
	aspect                string
	resolution            string
	videoBitRate          string
	videoBitRateTolerance int
	videoMaxBitRate       int
	videoMinBitrate       int
	videoCodec            string
	vframes               int
	frameRate             int
	audioRate             int
	maxKeyframe           int
	minKeyframe           int
	keyframeInterval      int
	audioCodec            string
	audioBitrate          string
	audioChannels         int
	audioVariableBitrate  bool
	bufferSize            int
	threadset             bool
	threads               int
	preset                string
	tune                  string
	audioProfile          string
	videoProfile          string
	target                string
	duration              string
	durationInput         string
	seekTime              string
	qscale                uint32
	crf                   uint32
	strict                int
	singleFile            int
	muxDelay              string
	seekUsingTsInput      bool
	seekTimeInput         string
	inputPath             string
	inputPipe             bool
	inputPipeReader       io.ReadCloser
	inputPipeWriter       io.Writer
	outputPipe            bool
	outputPipeReader      io.Reader
	outputPipeWriter      io.WriteCloser
	movFlags              string
	hideBanner            bool
	outputPath            string
	outputFormat          string
	copyTs                bool
	nativeFramerateInput  bool
	inputInitialOffset    string
	rtmpLive              string
	hlsPlaylistType       string
	hlsListSize           int
	hlsSegmentDuration    int
	hlsMasterPlaylistName string
	hlsSegmentFilename    string
	httpMethod            string
	httpKeepAlive         bool
	hwaccel               string
	streamIds             map[int]string
	metadata              Metadata
	videoFilter           string
	audioFilter           string
	skipVideo             bool
	skipAudio             bool
	compressionLevel      int
	mapMetadata           string
	tags                  map[string]string
	encryptionKey         string
	movflags              string
	bframe                int
	pixFmt                string
	rawInputArgs          []string
	rawOutputArgs         []string
}

/*** SETTERS ***/
func (m *File) SetAudioFilter(v string) {
	m.audioFilter = v
}

func (m *File) SetVideoFilter(v string) {
	m.videoFilter = v
}

// Deprecated: Use SetVideoFilter instead.
func (m *File) SetFilter(v string) {
	m.SetVideoFilter(v)
}

func (m *File) SetAspect(v string) {
	m.aspect = v
}

func (m *File) SetResolution(v string) {
	m.resolution = v
}

func (m *File) SetVideoBitRate(v string) {
	m.videoBitRate = v
}

func (m *File) SetVideoBitRateTolerance(v int) {
	m.videoBitRateTolerance = v
}

func (m *File) SetVideoMaxBitrate(v int) {
	m.videoMaxBitRate = v
}

func (m *File) SetVideoMinBitRate(v int) {
	m.videoMinBitrate = v
}

func (m *File) SetVideoCodec(v string) {
	m.videoCodec = v
}

func (m *File) SetVframes(v int) {
	m.vframes = v
}

func (m *File) SetFrameRate(v int) {
	m.frameRate = v
}

func (m *File) SetAudioRate(v int) {
	m.audioRate = v
}

func (m *File) SetAudioVariableBitrate() {
	m.audioVariableBitrate = true
}

func (m *File) SetMaxKeyFrame(v int) {
	m.maxKeyframe = v
}

func (m *File) SetMinKeyFrame(v int) {
	m.minKeyframe = v
}

func (m *File) SetKeyframeInterval(v int) {
	m.keyframeInterval = v
}

func (m *File) SetAudioCodec(v string) {
	m.audioCodec = v
}

func (m *File) SetAudioBitRate(v string) {
	m.audioBitrate = v
}

func (m *File) SetAudioChannels(v int) {
	m.audioChannels = v
}

func (m *File) SetPixFmt(v string) {
	m.pixFmt = v
}

func (m *File) SetBufferSize(v int) {
	m.bufferSize = v
}

func (m *File) SetThreads(v int) {
	m.threadset = true
	m.threads = v
}

func (m *File) SetPreset(v string) {
	m.preset = v
}

func (m *File) SetTune(v string) {
	m.tune = v
}

func (m *File) SetAudioProfile(v string) {
	m.audioProfile = v
}

func (m *File) SetVideoProfile(v string) {
	m.videoProfile = v
}

func (m *File) SetDuration(v string) {
	m.duration = v
}

func (m *File) SetDurationInput(v string) {
	m.durationInput = v
}

func (m *File) SetSeekTime(v string) {
	m.seekTime = v
}

func (m *File) SetSeekTimeInput(v string) {
	m.seekTimeInput = v
}

// Q Scale must be integer between 1 to 31 - https://trac.ffmpeg.org/wiki/Encode/MPEG-4
func (m *File) SetQScale(v uint32) {
	m.qscale = v
}

func (m *File) SetCRF(v uint32) {
	m.crf = v
}

func (m *File) SetStrict(v int) {
	m.strict = v
}

func (m *File) SetSingleFile(v int) {
	m.singleFile = v
}

func (m *File) SetSeekUsingTsInput(val bool) {
	m.seekUsingTsInput = val
}

func (m *File) SetCopyTs(val bool) {
	m.copyTs = val
}

func (m *File) SetInputPath(val string) {
	m.inputPath = val
}

func (m *File) SetInputPipe(val bool) {
	m.inputPipe = val
}

func (m *File) SetInputPipeReader(r io.ReadCloser) {
	m.inputPipeReader = r
}

func (m *File) SetInputPipeWriter(w io.Writer) {
	m.inputPipeWriter = w
}

func (m *File) SetOutputPipe(val bool) {
	m.outputPipe = val
}

func (m *File) SetOutputPipeReader(r io.Reader) {
	m.outputPipeReader = r
}

func (m *File) SetOutputPipeWriter(w io.WriteCloser) {
	m.outputPipeWriter = w
}

func (m *File) SetMovFlags(val string) {
	m.movFlags = val
}

func (m *File) SetHideBanner(val bool) {
	m.hideBanner = val
}

func (m *File) SetMuxDelay(val string) {
	m.muxDelay = val
}

func (m *File) SetOutputPath(val string) {
	m.outputPath = val
}

func (m *File) SetOutputFormat(val string) {
	m.outputFormat = val
}

func (m *File) SetNativeFramerateInput(val bool) {
	m.nativeFramerateInput = val
}

func (m *File) SetRtmpLive(val string) {
	m.rtmpLive = val
}

func (m *File) SetHlsListSize(val int) {
	m.hlsListSize = val
}

func (m *File) SetHlsSegmentDuration(val int) {
	m.hlsSegmentDuration = val
}

func (m *File) SetHlsPlaylistType(val string) {
	m.hlsPlaylistType = val
}

func (m *File) SetHlsMasterPlaylistName(val string) {
	m.hlsMasterPlaylistName = val
}

func (m *File) SetHlsSegmentFilename(val string) {
	m.hlsSegmentFilename = val
}

func (m *File) SetHttpMethod(val string) {
	m.httpMethod = val
}

func (m *File) SetHttpKeepAlive(val bool) {
	m.httpKeepAlive = val
}

func (m *File) SetHardwareAcceleration(val string) {
	m.hwaccel = val
}

func (m *File) SetInputInitialOffset(val string) {
	m.inputInitialOffset = val
}

func (m *File) SetStreamIds(val map[int]string) {
	m.streamIds = val
}

func (m *File) SetSkipVideo(val bool) {
	m.skipVideo = val
}

func (m *File) SetSkipAudio(val bool) {
	m.skipAudio = val
}

func (m *File) SetMetadata(v Metadata) {
	m.metadata = v
}

func (m *File) SetCompressionLevel(val int) {
	m.compressionLevel = val
}

func (m *File) SetMapMetadata(val string) {
	m.mapMetadata = val
}

func (m *File) SetTags(val map[string]string) {
	m.tags = val
}

func (m *File) SetBframe(v int) {
	m.bframe = v
}

func (m *File) SetRawInputArgs(args []string) {
	m.rawInputArgs = args
}

func (m *File) SetRawOutputArgs(args []string) {
	m.rawOutputArgs = args
}

/*** GETTERS ***/

// Deprecated: Use VideoFilter instead.
func (m *File) Filter() string {
	return m.VideoFilter()
}

func (m *File) VideoFilter() string {
	return m.videoFilter
}

func (m *File) AudioFilter() string {
	return m.audioFilter
}

func (m *File) Aspect() string {
	return m.aspect
}

func (m *File) Resolution() string {
	return m.resolution
}

func (m *File) VideoBitrate() string {
	return m.videoBitRate
}

func (m *File) VideoBitRateTolerance() int {
	return m.videoBitRateTolerance
}

func (m *File) VideoMaxBitRate() int {
	return m.videoMaxBitRate
}

func (m *File) VideoMinBitRate() int {
	return m.videoMinBitrate
}

func (m *File) VideoCodec() string {
	return m.videoCodec
}

func (m *File) Vframes() int {
	return m.vframes
}

func (m *File) FrameRate() int {
	return m.frameRate
}

func (m *File) GetPixFmt() string {
	return m.pixFmt
}

func (m *File) AudioRate() int {
	return m.audioRate
}

func (m *File) MaxKeyFrame() int {
	return m.maxKeyframe
}

func (m *File) MinKeyFrame() int {
	return m.minKeyframe
}

func (m *File) KeyFrameInterval() int {
	return m.keyframeInterval
}

func (m *File) AudioCodec() string {
	return m.audioCodec
}

func (m *File) AudioBitrate() string {
	return m.audioBitrate
}

func (m *File) AudioChannels() int {
	return m.audioChannels
}

func (m *File) BufferSize() int {
	return m.bufferSize
}

func (m *File) Threads() int {
	return m.threads
}

func (m *File) Target() string {
	return m.target
}

func (m *File) Duration() string {
	return m.duration
}

func (m *File) DurationInput() string {
	return m.durationInput
}

func (m *File) SeekTime() string {
	return m.seekTime
}

func (m *File) Preset() string {
	return m.preset
}

func (m *File) AudioProfile() string {
	return m.audioProfile
}

func (m *File) VideoProfile() string {
	return m.videoProfile
}

func (m *File) Tune() string {
	return m.tune
}

func (m *File) SeekTimeInput() string {
	return m.seekTimeInput
}

func (m *File) QScale() uint32 {
	return m.qscale
}

func (m *File) CRF() uint32 {
	return m.crf
}

func (m *File) Strict() int {
	return m.strict
}

func (m *File) SingleFile() int {
	return m.singleFile
}

func (m *File) MuxDelay() string {
	return m.muxDelay
}

func (m *File) SeekUsingTsInput() bool {
	return m.seekUsingTsInput
}

func (m *File) CopyTs() bool {
	return m.copyTs
}

func (m *File) InputPath() string {
	return m.inputPath
}

func (m *File) InputPipe() bool {
	return m.inputPipe
}

func (m *File) InputPipeReader() io.ReadCloser {
	return m.inputPipeReader
}

func (m *File) InputPipeWriter() io.Writer {
	return m.inputPipeWriter
}

func (m *File) OutputPipe() bool {
	return m.outputPipe
}

func (m *File) OutputPipeReader() io.Reader {
	return m.outputPipeReader
}

func (m *File) OutputPipeWriter() io.WriteCloser {
	return m.outputPipeWriter
}

func (m *File) MovFlags() string {
	return m.movFlags
}

func (m *File) HideBanner() bool {
	return m.hideBanner
}

func (m *File) OutputPath() string {
	return m.outputPath
}

func (m *File) OutputFormat() string {
	return m.outputFormat
}

func (m *File) NativeFramerateInput() bool {
	return m.nativeFramerateInput
}

func (m *File) RtmpLive() string {
	return m.rtmpLive
}

func (m *File) HlsListSize() int {
	return m.hlsListSize
}

func (m *File) HlsSegmentDuration() int {
	return m.hlsSegmentDuration
}

func (m *File) HlsMasterPlaylistName() string {
	return m.hlsMasterPlaylistName
}

func (m *File) HlsSegmentFilename() string {
	return m.hlsSegmentFilename
}

func (m *File) HlsPlaylistType() string {
	return m.hlsPlaylistType
}

func (m *File) InputInitialOffset() string {
	return m.inputInitialOffset
}

func (m *File) HttpMethod() string {
	return m.httpMethod
}

func (m *File) HttpKeepAlive() bool {
	return m.httpKeepAlive
}

func (m *File) HardwareAcceleration() string {
	return m.hwaccel
}

func (m *File) StreamIds() map[int]string {
	return m.streamIds
}

func (m *File) SkipVideo() bool {
	return m.skipVideo
}

func (m *File) SkipAudio() bool {
	return m.skipAudio
}

func (m *File) Metadata() Metadata {
	return m.metadata
}

func (m *File) CompressionLevel() int {
	return m.compressionLevel
}

func (m *File) MapMetadata() string {
	return m.mapMetadata
}

func (m *File) Tags() map[string]string {
	return m.tags
}

func (m *File) SetEncryptionKey(v string) {
	m.encryptionKey = v
}

func (m *File) EncryptionKey() string {
	return m.encryptionKey
}

func (m *File) RawInputArgs() []string {
	return m.rawInputArgs
}

func (m *File) RawOutputArgs() []string {
	return m.rawOutputArgs
}

/** OPTS **/
func (m *File) ToStrCommand() []string {
	var strCommand []string

	opts := []string{
		"SeekTimeInput",
		"SeekUsingTsInput",
		"NativeFramerateInput",
		"DurationInput",
		"RtmpLive",
		"InputInitialOffset",
		"HardwareAcceleration",
		"RawInputArgs",
		"InputPath",
		"InputPipe",
		"HideBanner",
		"Aspect",
		"Resolution",
		"FrameRate",
		"AudioRate",
		"VideoCodec",
		"Vframes",
		"VideoBitRate",
		"VideoBitRateTolerance",
		"VideoMaxBitRate",
		"VideoMinBitRate",
		"VideoProfile",
		"SkipVideo",
		"AudioCodec",
		"AudioBitRate",
		"AudioChannels",
		"AudioProfile",
		"SkipAudio",
		"CRF",
		"QScale",
		"Strict",
		"SingleFile",
		"BufferSize",
		"MuxDelay",
		"Threads",
		"KeyframeInterval",
		"Preset",
		"PixFmt",
		"Tune",
		"Target",
		"SeekTime",
		"Duration",
		"CopyTs",
		"StreamIds",
		"MovFlags",
		"RawOutputArgs",
		"OutputFormat",
		"OutputPipe",
		"HlsListSize",
		"HlsSegmentDuration",
		"HlsPlaylistType",
		"HlsMasterPlaylistName",
		"HlsSegmentFilename",
		"AudioFilter",
		"VideoFilter",
		"HttpMethod",
		"HttpKeepAlive",
		"CompressionLevel",
		"MapMetadata",
		"Tags",
		"EncryptionKey",
		"OutputPath",
		"Bframe",
		"MovFlags",
	}

	for _, name := range opts {
		opt := reflect.ValueOf(m).MethodByName(fmt.Sprintf("Obtain%s", name))
		if (opt != reflect.Value{}) {
			result := opt.Call([]reflect.Value{})

			if val, ok := result[0].Interface().([]string); ok {
				strCommand = append(strCommand, val...)
			}
		}
	}

	return strCommand
}

func (m *File) ObtainAudioFilter() []string {
	if m.audioFilter != "" {
		return []string{"-af", m.audioFilter}
	}
	return nil
}

func (m *File) ObtainVideoFilter() []string {
	if m.videoFilter != "" {
		return []string{"-vf", m.videoFilter}
	}
	return nil
}

func (m *File) ObtainAspect() []string {
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

func (m *File) ObtainHardwareAcceleration() []string {
	if m.hwaccel != "" {
		return []string{"-hwaccel", m.hwaccel}
	}
	return nil
}

func (m *File) ObtainInputPath() []string {
	if m.inputPath != "" {
		return []string{"-i", m.inputPath}
	}
	return nil
}

func (m *File) ObtainInputPipe() []string {
	if m.inputPipe {
		return []string{"-i", "pipe:0"}
	}
	return nil
}

func (m *File) ObtainOutputPipe() []string {
	if m.outputPipe {
		return []string{"pipe:1"}
	}
	return nil
}

func (m *File) ObtainMovFlags() []string {
	if m.movFlags != "" {
		return []string{"-movflags", m.movFlags}
	}
	return nil
}

func (m *File) ObtainHideBanner() []string {
	if m.hideBanner {
		return []string{"-hide_banner"}
	}
	return nil
}

func (m *File) ObtainNativeFramerateInput() []string {
	if m.nativeFramerateInput {
		return []string{"-re"}
	}
	return nil
}

func (m *File) ObtainOutputPath() []string {
	if m.outputPath != "" {
		return []string{m.outputPath}
	}
	return nil
}

func (m *File) ObtainVideoCodec() []string {
	if m.videoCodec != "" {
		return []string{"-c:v", m.videoCodec}
	}
	return nil
}

func (m *File) ObtainVframes() []string {
	if m.vframes != 0 {
		return []string{"-vframes", fmt.Sprintf("%d", m.vframes)}
	}
	return nil
}

func (m *File) ObtainFrameRate() []string {
	if m.frameRate != 0 {
		return []string{"-r", fmt.Sprintf("%d", m.frameRate)}
	}
	return nil
}

func (m *File) ObtainAudioRate() []string {
	if m.audioRate != 0 {
		return []string{"-ar", fmt.Sprintf("%d", m.audioRate)}
	}
	return nil
}

func (m *File) ObtainResolution() []string {
	if m.resolution != "" {
		return []string{"-s", m.resolution}
	}
	return nil
}

func (m *File) ObtainVideoBitRate() []string {
	if m.videoBitRate != "" {
		return []string{"-b:v", m.videoBitRate}
	}
	return nil
}

func (m *File) ObtainAudioCodec() []string {
	if m.audioCodec != "" {
		return []string{"-c:a", m.audioCodec}
	}
	return nil
}

func (m *File) ObtainAudioBitRate() []string {
	switch {
	case !m.audioVariableBitrate && m.audioBitrate != "":
		return []string{"-b:a", m.audioBitrate}
	case m.audioVariableBitrate && m.audioBitrate != "":
		return []string{"-q:a", m.audioBitrate}
	case m.audioVariableBitrate:
		return []string{"-q:a", "0"}
	default:
		return nil
	}
}

func (m *File) ObtainAudioChannels() []string {
	if m.audioChannels != 0 {
		return []string{"-ac", fmt.Sprintf("%d", m.audioChannels)}
	}
	return nil
}

func (m *File) ObtainVideoMaxBitRate() []string {
	if m.videoMaxBitRate != 0 {
		return []string{"-maxrate", fmt.Sprintf("%dk", m.videoMaxBitRate)}
	}
	return nil
}

func (m *File) ObtainVideoMinBitRate() []string {
	if m.videoMinBitrate != 0 {
		return []string{"-minrate", fmt.Sprintf("%dk", m.videoMinBitrate)}
	}
	return nil
}

func (m *File) ObtainBufferSize() []string {
	if m.bufferSize != 0 {
		return []string{"-bufsize", fmt.Sprintf("%dk", m.bufferSize)}
	}
	return nil
}

func (m *File) ObtainVideoBitRateTolerance() []string {
	if m.videoBitRateTolerance != 0 {
		return []string{"-bt", fmt.Sprintf("%dk", m.videoBitRateTolerance)}
	}
	return nil
}

func (m *File) ObtainThreads() []string {
	if m.threadset {
		return []string{"-threads", fmt.Sprintf("%d", m.threads)}
	}
	return nil
}

func (m *File) ObtainTarget() []string {
	if m.target != "" {
		return []string{"-target", m.target}
	}
	return nil
}

func (m *File) ObtainDuration() []string {
	if m.duration != "" {
		return []string{"-t", m.duration}
	}
	return nil
}

func (m *File) ObtainDurationInput() []string {
	if m.durationInput != "" {
		return []string{"-t", m.durationInput}
	}
	return nil
}

func (m *File) ObtainKeyframeInterval() []string {
	if m.keyframeInterval != 0 {
		return []string{"-g", fmt.Sprintf("%d", m.keyframeInterval)}
	}
	return nil
}

func (m *File) ObtainSeekTime() []string {
	if m.seekTime != "" {
		return []string{"-ss", m.seekTime}
	}
	return nil
}

func (m *File) ObtainSeekTimeInput() []string {
	if m.seekTimeInput != "" {
		return []string{"-ss", m.seekTimeInput}
	}
	return nil
}

func (m *File) ObtainPreset() []string {
	if m.preset != "" {
		return []string{"-preset", m.preset}
	}
	return nil
}

func (m *File) ObtainTune() []string {
	if m.tune != "" {
		return []string{"-tune", m.tune}
	}
	return nil
}

func (m *File) ObtainCRF() []string {
	if m.crf != 0 {
		return []string{"-crf", fmt.Sprintf("%d", m.crf)}
	}
	return nil
}

func (m *File) ObtainQScale() []string {
	if m.qscale != 0 {
		return []string{"-qscale", fmt.Sprintf("%d", m.qscale)}
	}
	return nil
}

func (m *File) ObtainStrict() []string {
	if m.strict != 0 {
		return []string{"-strict", fmt.Sprintf("%d", m.strict)}
	}
	return nil
}

func (m *File) ObtainSingleFile() []string {
	if m.singleFile != 0 {
		return []string{"-single_file", fmt.Sprintf("%d", m.singleFile)}
	}
	return nil
}

func (m *File) ObtainVideoProfile() []string {
	if m.videoProfile != "" {
		return []string{"-profile:v", m.videoProfile}
	}
	return nil
}

func (m *File) ObtainAudioProfile() []string {
	if m.audioProfile != "" {
		return []string{"-profile:a", m.audioProfile}
	}
	return nil
}

func (m *File) ObtainCopyTs() []string {
	if m.copyTs {
		return []string{"-copyts"}
	}
	return nil
}

func (m *File) ObtainOutputFormat() []string {
	if m.outputFormat != "" {
		return []string{"-f", m.outputFormat}
	}
	return nil
}

func (m *File) ObtainMuxDelay() []string {
	if m.muxDelay != "" {
		return []string{"-muxdelay", m.muxDelay}
	}
	return nil
}

func (m *File) ObtainSeekUsingTsInput() []string {
	if m.seekUsingTsInput {
		return []string{"-seek_timestamp", "1"}
	}
	return nil
}

func (m *File) ObtainRtmpLive() []string {
	if m.rtmpLive != "" {
		return []string{"-rtmp_live", m.rtmpLive}
	} else {
		return nil
	}
}

func (m *File) ObtainHlsPlaylistType() []string {
	if m.hlsPlaylistType != "" {
		return []string{"-hls_playlist_type", m.hlsPlaylistType}
	} else {
		return nil
	}
}

func (m *File) ObtainInputInitialOffset() []string {
	if m.inputInitialOffset != "" {
		return []string{"-itsoffset", m.inputInitialOffset}
	} else {
		return nil
	}
}

func (m *File) ObtainHlsListSize() []string {
	return []string{"-hls_list_size", fmt.Sprintf("%d", m.hlsListSize)}
}

func (m *File) ObtainHlsSegmentDuration() []string {
	if m.hlsSegmentDuration != 0 {
		return []string{"-hls_time", fmt.Sprintf("%d", m.hlsSegmentDuration)}
	} else {
		return nil
	}
}

func (m *File) ObtainHlsMasterPlaylistName() []string {
	if m.hlsMasterPlaylistName != "" {
		return []string{"-master_pl_name", fmt.Sprintf("%s", m.hlsMasterPlaylistName)}
	} else {
		return nil
	}
}

func (m *File) ObtainHlsSegmentFilename() []string {
	if m.hlsSegmentFilename != "" {
		return []string{"-hls_segment_filename", fmt.Sprintf("%s", m.hlsSegmentFilename)}
	} else {
		return nil
	}
}

func (m *File) ObtainHttpMethod() []string {
	if m.httpMethod != "" {
		return []string{"-method", m.httpMethod}
	} else {
		return nil
	}
}

func (m *File) ObtainPixFmt() []string {
	if m.pixFmt != "" {
		return []string{"-pix_fmt", m.pixFmt}
	} else {
		return nil
	}
}

func (m *File) ObtainHttpKeepAlive() []string {
	if m.httpKeepAlive {
		return []string{"-multiple_requests", "1"}
	} else {
		return nil
	}
}

func (m *File) ObtainSkipVideo() []string {
	if m.skipVideo {
		return []string{"-vn"}
	} else {
		return nil
	}
}

func (m *File) ObtainSkipAudio() []string {
	if m.skipAudio {
		return []string{"-an"}
	} else {
		return nil
	}
}

func (m *File) ObtainStreamIds() []string {
	if m.streamIds != nil && len(m.streamIds) != 0 {
		result := []string{}
		for i, val := range m.streamIds {
			result = append(result, []string{"-streamid", fmt.Sprintf("%d:%s", i, val)}...)
		}
		return result
	}
	return nil
}

func (m *File) ObtainCompressionLevel() []string {
	if m.compressionLevel != 0 {
		return []string{"-compression_level", fmt.Sprintf("%d", m.compressionLevel)}
	}
	return nil
}

func (m *File) ObtainMapMetadata() []string {
	if m.mapMetadata != "" {
		return []string{"-map_metadata", m.mapMetadata}
	}
	return nil
}

func (m *File) ObtainEncryptionKey() []string {
	if m.encryptionKey != "" {
		return []string{"-hls_key_info_file", m.encryptionKey}
	}

	return nil
}

func (m *File) ObtainBframe() []string {
	if m.bframe != 0 {
		return []string{"-bf", fmt.Sprintf("%d", m.bframe)}
	}
	return nil
}

func (m *File) ObtainTags() []string {
	if m.tags != nil && len(m.tags) != 0 {
		result := []string{}
		for key, val := range m.tags {
			result = append(result, []string{"-metadata", fmt.Sprintf("%s=%s", key, val)}...)
		}
		return result
	}
	return nil
}

func (m *File) ObtainRawInputArgs() []string {
	return m.rawInputArgs
}

func (m *File) ObtainRawOutputArgs() []string {
	return m.rawOutputArgs
}

func CheckFileType(streams []Streams) string {
	for i := 0; i < len(streams); i++ {
		st := streams[i]
		if st.CodecType == "video" {
			return "video"
		}
	}

	return "audio"
}
