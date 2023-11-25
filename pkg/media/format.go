package media

type FileFormat struct {
	Filename       string
	NbStreams      int            `json:"nb_streams"`
	NbPrograms     int            `json:"nb_programs"`
	FormatName     string         `json:"format_name"`
	FormatLongName string         `json:"format_long_name"`
	Duration       string         `json:"duration"`
	Size           string         `json:"size"`
	BitRate        string         `json:"bit_rate"`
	ProbeScore     int            `json:"probe_score"`
	Tags           FileFormatTags `json:"tags"`
}

type FileFormatTags struct {
	Encoder string `json:"ENCODER"`
}
