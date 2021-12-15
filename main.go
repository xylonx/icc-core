package main

import (
	"os"

	"github.com/xylonx/go-template/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
