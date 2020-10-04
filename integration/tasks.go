// Package db holds functions related to database abstraction.
package integration

import (
	"encoding/binary"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")

var db *bolt.DB

// Task has a numerical ID called 'Key' and Value which comprises text.
type Task struct {
	Key   int
	Value string
}

// Init creates db file and create appropriate bucket if
// not already exists. This is required for inserts later on.
func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, nil)
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

// CreateTask Puts a new entry in the bucket.
// For autoincrement we use the Bucket.NextSequence API.
// https://pkg.go.dev/github.com/boltdb/bolt?tab=doc#Bucket.NextSequence
func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)
		return b.Put(key, []byte(task))
	})
	if err != nil {
		// expose error to the caller
		return -1, err
	}
	return id, nil
}

// AllTasks iterates over all the entires in the task bucket
// and returns a slice of all the tasks.
func AllTasks() ([]Task, error) {
	var tasks []Task

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		// we need cursor for iteration
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		// should return nil to complete the transaction
		return nil
	})
	if err != nil {
		// expose error to the caller
		return nil, err
	}

	return tasks, nil
}

// DeleteTask looks for entry by 'key' and deletes it.
// If the key does not exist then nothing is done and a nil error is returned.
// Returns an error if the bucket was created from a read-only transaction.
func DeleteTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})
}

// itob converts from int to Bolt friendly binary
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// btoi converts from byte to int
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
