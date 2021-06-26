package cmd

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
	"strconv"
)

func init() {
	cmd.AddCommand(rmCmd)
}

var rmCmd = &cobra.Command{
	Use: "rm",
	Short: "Removes task from the list",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Only one integer argument should be provided.")
			return
		}
		var con conn
		db , err := bolt.Open(DbFile, 0755, nil)
		checkPanic(err)
		con.db = db
		err = con.Generate()
		checkPanic(err)
		i, err := strconv.Atoi(args[0])
		checkPanic(err)
		ret, err := con.RmTask(i)
		checkPanic(err)
		if ret == "" {
			fmt.Println("There is no such task")
		} else {
			fmt.Println("\"" + ret + "\" has been deleted")
		}
	},
}