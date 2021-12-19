package main

import (
	"os"

	"github.com/xylonx/icc-core/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
