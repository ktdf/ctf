package cmd

import (
	"github.com/boltdb/bolt"
	"time"
)

type conn struct {
	db *bolt.DB
}

func (db *conn) Generate() error {
	if err := db.db.Batch(func(tx *bolt.Tx) error {
		for _, name := range []string{"active", "completed"} {
			_, err := tx.CreateBucketIfNotExists([]byte(name))
			if err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (db *conn) AddTask(task []string) error {
	timestamp := TimestampToByte(time.Now().Unix())
	if err := db.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("active"))
		err := b.Put(timestamp, []byte(sliceToString(task)))
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
