package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	cmd.AddCommand(addDo)
}

var addDo = &cobra.Command{
	Use: "do",
	Short: "Marks task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Do")
	},
}