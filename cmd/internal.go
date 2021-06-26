package cmd

import "strings"

func checkPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func sliceToString(s []string) string {
	var db strings.Builder
	for _, word := range s {
		db.WriteString(word)
	}
	return db.String()
}