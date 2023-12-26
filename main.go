package main

import (
	"os"

	"github.com/apodeixis/backend/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
