package cmd

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

func init() {
	cmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use: "add",
	Short: "Adds new task to the list",
	Run: func(cmd *cobra.Command, args []string) {
		var con conn
		db , err := bolt.Open("ctm.db", 0755, nil)
		checkPanic(err)
		con.db = db
		err = con.Generate()
		checkPanic(err)
		err = con.AddTask(args)
		checkPanic(err)
		fmt.Println("\""+sliceToString(args)+"\" has been added to the task list")
	},
}