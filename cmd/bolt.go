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
