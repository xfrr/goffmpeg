package models

import "time"

type Ffmpeg struct {
	FfmpegBinPath  string
	FfprobeBinPath string
}

type Metadata struct {
	Streams []Streams `json:"streams"`
	Format  Format    `json:"format"`
}

type Streams struct {
	Index              int
	ID                 string      `json:"id"`
	CodecName          string      `json:"codec_name"`
	CodecLongName      string      `json:"codec_long_name"`
	Profile            string      `json:"profile"`
	CodecType          string      `json:"codec_type"`
	CodecTimeBase      string      `json:"codec_time_base"`
	CodecTagString     string      `json:"codec_tag_string"`
	CodecTag           string      `json:"codec_tag"`
	Width              int         `json:"width"`
	Height             int         `json:"height"`
	CodedWidth         int         `json:"coded_width"`
	CodedHeight        int         `json:"coded_height"`
	HasBFrames         int         `json:"has_b_frames"`
	SampleAspectRatio  string      `json:"sample_aspect_ratio"`
	DisplayAspectRatio string      `json:"display_aspect_ratio"`
	PixFmt             string      `json:"pix_fmt"`
	Level              int         `json:"level"`
	ChromaLocation     string      `json:"chroma_location"`
	Refs               int         `json:"refs"`
	QuarterSample      string      `json:"quarter_sample"`
	DivxPacked         string      `json:"divx_packed"`
	RFrameRrate        string      `json:"r_frame_rate"`
	AvgFrameRate       string      `json:"avg_frame_rate"`
	TimeBase           string      `json:"time_base"`
	DurationTs         int         `json:"duration_ts"`
	Duration           string      `json:"duration"`
	BitRate            string      `json:"bit_rate"`
	Disposition        Disposition `json:"disposition"`
	SideDataList       []SideData  `json:"side_data_list"`
	Tags               *StreamTags `json:"tags"`
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
}

type Format struct {
	Filename       string
	NbStreams      int    `json:"nb_streams"`
	NbPrograms     int    `json:"nb_programs"`
	FormatName     string `json:"format_name"`
	FormatLongName string `json:"format_long_name"`
	Duration       string `json:"duration"`
	Size           string `json:"size"`
	BitRate        string `json:"bit_rate"`
	ProbeScore     int    `json:"probe_score"`
	Tags           Tags   `json:"tags"`
}

type SideData struct {
	SideDataType  *string `json:"side_data_type"`
	DisplayMatrix *string `json:"displaymatrix"`
	Rotation      *int    `json:"rotation"`
	MaxContent    *int    `json:"max_content"`
	MaxAverage    *int    `json:"max_average"`
	RedX          *string `json:"red_x"`
	RedY          *string `json:"red_y"`
	GreenX        *string `json:"green_x"`
	GreenY        *string `json:"green_y"`
	BlueX         *string `json:"blue_x"`
	BlueY         *string `json:"blue_y"`
	WhitePointX   *string `json:"white_point_x"`
	WhitePointY   *string `json:"white_point_y"`
	MinLuminance  *string `json:"min_luminance"`
	MaxLuminance  *string `json:"max_luminance"`
}

type StreamTags struct {
	CreationTime *time.Time `json:"creation_time"`
	Language     *string    `json:"language"`
	HandlerName  *string    `json:"handler_name"`
	VendorID     *string    `json:"vendor_id"`
	Encoder      *string    `json:"encoder"`
}

type Progress struct {
	FramesProcessed string
	CurrentTime     string
	CurrentBitrate  string
	Progress        float64
	Speed           string
}

type Tags struct {
	Encoder string `json:"ENCODER"`
}
