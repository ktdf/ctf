package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	cmd.AddCommand(completedCmd)
}

var completedCmd = &cobra.Command{
	Use: "completed",
	Short: "List completed tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("completed")
	},
}