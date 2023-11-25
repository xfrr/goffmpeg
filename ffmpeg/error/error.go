package fferror

import "fmt"

var (
	ErrUnableToFindOutputFormat = fmt.Errorf("unable to find a suitable output format")
	ErrInputFileNotFound        = fmt.Errorf("no such file or directory")
)
