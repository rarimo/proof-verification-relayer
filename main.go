package main

import (
	"os"

	"github.com/rarimo/proof-verification-relayer/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
