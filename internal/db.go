package internal

import (
	"encoding/binary"

	"go.etcd.io/bbolt"
)

var taskBucket = []byte("tasks")

func OpenDB(path string) (*bbolt.DB, error) {
	db, err := bbolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}

	// Ensure the bucket for tasks exists
	db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})

	return db, nil
}

func AddTask(db *bbolt.DB, task string) error {
	return db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id, _ := b.NextSequence()
		return b.Put(itob(id), []byte(task))
	})
}

func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

