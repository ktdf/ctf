package cmd

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

func init() {
	cmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use: "list",
	Short: "Print list of commands",
	Run: func(cmd *cobra.Command, args []string) {
		var con conn
		db , err := bolt.Open(DbFile, 0755, nil)
		checkPanic(err)
		con.db = db
		err = con.Generate()
		checkPanic(err)
		tasks := con.ListTasks()
		for n, task := range tasks {
			fmt.Printf("%3v: %s\n", n+1, task)
		}
	},
}