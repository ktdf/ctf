package cmd

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
	"strconv"
)

func init() {
	cmd.AddCommand(addDo)
}

var addDo = &cobra.Command{
	Use: "do",
	Short: "Marks task as complete",
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
		ret, err := con.DoTask(i)
		checkPanic(err)
		if ret == "" {
			fmt.Println("It looks like nothing was deleted")
		} else {
			fmt.Println("\"" + ret + "\" has been completed")
		}
	},
}