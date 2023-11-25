package media

import (
	"encoding/json"
	"strconv"
	"time"
)

type FileProperties struct {
	Format  Format   `json:"format"`
	Streams []Stream `json:"streams"`
}

type File struct {
	properties FileProperties
}

func (f *File) UnmarshalJSON(data []byte) error {
	var fileProperties FileProperties
	if err := json.Unmarshal(data, &fileProperties); err != nil {
		return err
	}
	f.properties = fileProperties
	return nil
}

func (f File) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.properties)
}

func (f File) Filename() string {
	return f.properties.Format.Filename
}

func (f File) FormatName() string {
	return f.properties.Format.FormatName
}

func (f File) FormatLongName() string {
	return f.properties.Format.FormatLongName
}

func (f File) BitRate() string {
	return f.properties.Format.BitRate
}

func (f File) Size() string {
	return f.properties.Format.Size
}

func (f File) Encoder() string {
	return f.properties.Format.Tags.Encoder
}

func (f File) NbStreams() int {
	return f.properties.Format.NbStreams
}

func (f File) NbPrograms() int {
	return f.properties.Format.NbPrograms
}

func (f File) ProbeScore() int {
	return f.properties.Format.ProbeScore
}

func (f File) Tags() Tags {
	return f.properties.Format.Tags
}

func (f File) Duration() time.Duration {
	fileDurationSeconds, err := strconv.ParseFloat(f.properties.Format.Duration, 64)
	if err != nil {
		panic(err)
	}
	return time.Duration(fileDurationSeconds) * time.Second
}

func (f File) Streams() []Stream {
	return f.properties.Streams
}

func (f File) StreamByIndex(index int) *Stream {
	if len(f.properties.Streams) <= index {
		return nil
	}
	return &f.properties.Streams[index]
}
