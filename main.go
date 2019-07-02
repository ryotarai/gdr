package main

import (
	"os"

	"github.com/ryotarai/gdr/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
