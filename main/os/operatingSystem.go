package os

import (
	"runtime"
)

func amIWindows() bool {
	if runtime.GOOS == "windows" {
		return true
	}
	return false
}

func amIMacOS() bool {
	if runtime.GOOS == "darwin" {
		return true
	}
	return false
}
