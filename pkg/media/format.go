package media

type Format struct {
	NbStreams      int    `json:"nb_streams"`
	NbPrograms     int    `json:"nb_programs"`
	Filename       string `json:"filename"`
	FormatName     string `json:"format_name"`
	FormatLongName string `json:"format_long_name"`
	Duration       string `json:"duration"`
	Size           string `json:"size"`
	BitRate        string `json:"bit_rate"`
	ProbeScore     int    `json:"probe_score"`
	StartTime      string `json:"start_time"`
	Tags           Tags   `json:"tags"`
}

type Tags struct {
	MajorBrand       string `json:"major_brand,omitempty"`
	MinorVersion     string `json:"minor_version,omitempty"`
	CompatibleBrands string `json:"compatible_brands,omitempty"`
	CreationTime     string `json:"creation_time,omitempty"`
	Language         string `json:"language,omitempty"`
	HandlerName      string `json:"handler_name,omitempty"`
	HandlerType      string `json:"handler_type,omitempty"`
	Encoder          string `json:"encoder,omitempty"`
	EncoderLavc      string `json:"encoder_lavc,omitempty"`
	StreamKind       string `json:"STREAM_KIND,omitempty"`
	StreamKindID     string `json:"STREAM_KIND_ID,omitempty"`
	VendorID         string `json:"vendor_id,omitempty"`
}
