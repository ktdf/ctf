package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	cmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use: "list",
	Short: "Print list of commands",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("List")
	},
}