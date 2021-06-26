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
	err := db.addRecord(timestamp, []byte(sliceToString(task)), []byte("active"))
	if err != nil {
		return err
	}
	return nil
}

func (db *conn) ListTasks() []string {
	var ret []string
	db.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("active"))
		b.ForEach(func(_, v []byte) error {
			ret = append(ret, string(v))
			return nil
		})
		return nil
	})
	return ret
}

func (db *conn) DoTask(i int) (ret string, err error) {
	var neededK, neededV []byte
	err = db.db.Batch(func(tx *bolt.Tx) error {
		aBucket := tx.Bucket([]byte("active"))
		c := aBucket.Cursor()
		k, v := c.First()
		for n := 0; k != nil; k, v = c.Next() {
			n++
			if n == i {
				ret = string(v)
				neededK, neededV = k, v
				c.Delete()
				break
			}
		}
		return nil
	})
	err = db.addRecord(neededK, neededV, []byte("completed"))
	if err != nil {
		ret = ""
		return
	}
	return
}

func (db *conn) addRecord(timestamp, task, bucket []byte) error {
	if err := db.db.Batch(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		err := b.Put(timestamp, task)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
