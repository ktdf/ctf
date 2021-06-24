package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	cmd.AddCommand(rmCmd)
}

var rmCmd = &cobra.Command{
	Use: "rm",
	Short: "Removes task from the list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rm")
	},
}