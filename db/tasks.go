package db

import (
	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")

var db *bolt.DB

type Task struct {
	Key   int
	Value string
}

// Init creates db file and create appropriate bucket if
// not already exists. This is required for inserts later on.
func Init(dbPath string) error {
	var err error
	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
	// defer db.Close()
}
