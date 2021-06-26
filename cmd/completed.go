package cmd

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

func init() {
	cmd.AddCommand(completedCmd)
}

var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "List completed tasks",
	Run: func(cmd *cobra.Command, args []string) {
		var con conn
		db, err := bolt.Open(DbFile, 0755, nil)
		checkPanic(err)
		con.db = db
		err = con.Generate()
		checkPanic(err)
		tasks := con.ListCompletedTasks()
		for n, task := range tasks {
			fmt.Printf("%3v: %s\n", n+1, task)
		}
	},
}
