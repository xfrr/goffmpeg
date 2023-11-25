package media

type Stream struct {
	AvgFrameRate       string      `json:"avg_frame_rate,omitempty"`
	BitRate            string      `json:"bit_rate,omitempty"`
	BitsPerRawSample   string      `json:"bits_per_raw_sample,omitempty"`
	BitsPerSample      int         `json:"bits_per_sample,omitempty"`
	ChanneLayout       string      `json:"channel_layout,omitempty"`
	Channels           int         `json:"channels,omitempty"`
	ChromaLocation     string      `json:"chroma_location,omitempty"`
	ClosedCaptions     int         `json:"closed_captions,omitempty"`
	CodecLongName      string      `json:"codec_long_name,omitempty"`
	CodecName          string      `json:"codec_name,omitempty"`
	CodecTag           string      `json:"codec_tag,omitempty"`
	CodecTagString     string      `json:"codec_tag_string,omitempty"`
	CodecTimeBase      string      `json:"codec_time_base,omitempty"`
	CodecType          string      `json:"codec_type,omitempty"`
	CodedHeight        int         `json:"coded_height,omitempty"`
	CodedWidth         int         `json:"coded_width,omitempty"`
	DisplayAspectRatio string      `json:"display_aspect_ratio,omitempty"`
	Disposition        Disposition `json:"disposition"`
	DivxPacked         string      `json:"divx_packed,omitempty"`
	Duration           string      `json:"duration,omitempty"`
	DurationTs         int         `json:"duration_ts,omitempty"`
	ExtraDataSize      int         `json:"extradata_size,omitempty"`
	FieldOrder         string      `json:"field_order,omitempty"`
	HasBFrames         int         `json:"has_b_frames,omitempty"`
	Height             int         `json:"height,omitempty"`
	ID                 string      `json:"id"`
	Index              int         `json:"index,omitempty"`
	IsAvc              string      `json:"is_avc,omitempty"`
	InitialPadding     int         `json:"initial_padding,omitempty"`
	Level              int         `json:"level,omitempty"`
	NalLengthSize      string      `json:"nal_length_size,omitempty"`
	NbFrames           string      `json:"nb_frames,omitempty"`
	PixFmt             string      `json:"pix_fmt,omitempty"`
	Profile            string      `json:"profile,omitempty"`
	QuarterSample      string      `json:"quarter_sample,omitempty"`
	Refs               int         `json:"refs,omitempty"`
	RFrameRrate        string      `json:"r_frame_rate,omitempty"`
	SampleAspectRatio  string      `json:"sample_aspect_ratio,omitempty"`
	SampleFmt          string      `json:"sample_fmt,omitempty"`
	SampleRate         string      `json:"sample_rate,omitempty"`
	StartPts           int         `json:"start_pts,omitempty"`
	StartTime          string      `json:"start_time,omitempty"`
	Tags               Tags        `json:"tags"`
	TimeBase           string      `json:"time_base,omitempty"`
	Width              int         `json:"width,omitempty"`
}

type Disposition struct {
	Default         int `json:"default"`
	Dub             int `json:"dub"`
	Original        int `json:"original"`
	Comment         int `json:"comment"`
	Lyrics          int `json:"lyrics"`
	Karaoke         int `json:"karaoke"`
	Forced          int `json:"forced"`
	HearingImpaired int `json:"hearing_impaired"`
	VisualImpaired  int `json:"visual_impaired"`
	CleanEffects    int `json:"clean_effects"`
	AttachedPic     int `json:"attached_pic"`
	TimedThumbnails int `json:"timed_thumbnails"`
	Captions        int `json:"captions"`
	Descriptions    int `json:"descriptions"`
	Metadata        int `json:"metadata"`
	Dependent       int `json:"dependent"`
	StillImage      int `json:"still_image"`
}
