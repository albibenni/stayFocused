package main

import (
	"fmt"
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

func main() {
	if amIWindows() {
		fmt.Println("Hello from Windows")
	}
	if amIMacOS() {
		fmt.Println("Hello from Mac")
	}
}
