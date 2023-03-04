package main

import (
	"fmt"
	"runtime"
)

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Hello from Windows")
	}
	if runtime.GOOS == "darwin" {
		fmt.Println("Hello from Mac")
	}
}
