package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vumanhcuongit/gort_commands/pkg/commands"
)

var echoCmd = &cobra.Command{
	Use:     "echo",
	Aliases: []string{"echo"},
	Short:   "Echo a string",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := commands.Echo(args[0])
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(echoCmd)
}
