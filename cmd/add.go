package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	cmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use: "add",
	Short: "Adds new task to the list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("List")
	},
}