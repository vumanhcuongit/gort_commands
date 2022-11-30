package commands

import (
	"github.com/spf13/cobra"
)

var version = "0.0.1"

var rootCmd = &cobra.Command{
	Use:     "cbt",
	Version: version,
	Short:   "cbt - a simple CLI to transform and inspect strings",
	Long: `cbt aka Chu Ba Thong is a super fancy CLI (kidding)
   
One can use cbt to modify or inspect strings straight from the terminal`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
