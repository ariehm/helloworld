package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	fmt.Printf("Hello, world!\n")
	fmt.Printf("build: %s\n", getVCSInfo())
}

func getVCSInfo() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				return setting.Value
			}
		}
	}

	return "could not retrieve build info"
}
