package utility

import (
	"runtime"
	"strings"
)

func WinCharacter(s string) string {
	if runtime.GOOS == "windows" {
		s = strings.Replace(s, "/", "-", -1)
		s = strings.Replace(s, "  ", " ", -1)
		return strings.Replace(s, ":", "-", -1)
	}
	return s
}
