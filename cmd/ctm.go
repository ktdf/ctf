package cmd

import (
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "ctm",
	Short: "cli task manager",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	cmd.Execute()
}
