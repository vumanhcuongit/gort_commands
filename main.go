package main

import (
	"os"

	"github.com/vumanhcuongit/gort_commands/cmd/commands"
)

func main() {
	err := commands.Execute()
	if err != nil {
		os.Exit(1)
	}
}
