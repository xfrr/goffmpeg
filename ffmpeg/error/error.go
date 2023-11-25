package fferror

import "fmt"

var (
	ErrUnableToFindOutputFormat            = fmt.Errorf("unable to find a suitable output format")
	ErrFileNotFound                        = fmt.Errorf("no such file or directory")
	ErrInvalidDataFoundWhenProcessingInput = fmt.Errorf("invalid data found when processing input")
	ErrInvalidArgument                     = fmt.Errorf("invalid argument")
)
