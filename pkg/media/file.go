package media

import (
	"strconv"
	"time"
)

type File struct {
	Streams []FileStream `json:"streams"`
	Format  FileFormat   `json:"format"`
}

func (f File) GetDuration() time.Duration {
	fileDurationSeconds, err := strconv.ParseFloat(f.Format.Duration, 64)
	if err != nil {
		panic(err)
	}
	return time.Duration(fileDurationSeconds) * time.Second
}
