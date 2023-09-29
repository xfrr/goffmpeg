package cmd

import "runtime"

var (
	platform = runtime.GOOS
)

func getFindCommand() string {
	switch platform {
	case "windows":
		return "where"
	default:
		return "which"
	}
}
