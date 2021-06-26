package cmd

import (
	"encoding/binary"
	"strings"
)

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

func TimestampToByte(timestamp int64) []byte {
	var t = make([]byte, 8)
	binary.BigEndian.PutUint64(t, uint64(timestamp))
	return t
}