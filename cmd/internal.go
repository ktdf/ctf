package cmd

import (
	"encoding/binary"
	"strings"
)

var DbFile string

func init() {
	DbFile = "ctm.db"
}

func checkPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func sliceToString(s []string) string {
	var db strings.Builder
	for n, word := range s {
		if n != 0 {
			db.WriteString(" ")
		}
		db.WriteString(word)
	}
	return db.String()
}

func TimestampToByte(timestamp int64) []byte {
	var t = make([]byte, 8)
	binary.BigEndian.PutUint64(t, uint64(timestamp))
	return t
}