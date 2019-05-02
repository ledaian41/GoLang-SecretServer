package swagger

import (
	"encoding/json"
	"github.com/boltdb/bolt"
	"log"
	"time"
)

func getConnection() (*bolt.DB, error) {
	db, err := bolt.Open("secret.db", 0600, &bolt.Options{Timeout: 5 * time.Second})
	return db, err
}

func saveSecret(secret Secret) {
	db, err := getConnection()
	checkError(err)
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("secrets"))
		if err != nil {
			return err
		}
		encoded, err := json.Marshal(secret)
		if err != nil {
			return err
		}
		return b.Put([]byte(secret.Hash), []byte(encoded))
	})
	checkError(err)
}

func getSecret(hash string) *Secret {
	db, err := getConnection()
	if err != nil {
		log.Println(err)
		return nil
	}
	defer db.Close()
	secret := Secret{}
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("secrets"))
		value := b.Get([]byte(hash))
		err := json.Unmarshal(value, &secret)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Println(err, "Secret not found")
		return nil
	}
	return &secret
}

func updateRemainingViews(secret *Secret) {
	secret.RemainingViews--
	db, err := getConnection()
	checkError(err)
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("secrets"))
		encoded, err := json.Marshal(secret)
		if err != nil {
			return err
		}
		return bucket.Put([]byte(secret.Hash), []byte(encoded))
	})
	checkError(err)
}

func removeSecret(hash string) {
	db, err := getConnection()
	checkError(err)
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("secrets"))
		return bucket.Delete([]byte(hash))
	})
	checkError(err)
}
