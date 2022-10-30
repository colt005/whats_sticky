package database

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/colt005/whats_sticky/models"
	bolt "go.etcd.io/bbolt"
)

const (
	DB_FILE_NAME = "database.db"
	BUCKET_NAME  = "users"
)

type database struct {
	*bolt.DB
}

var Instance *database

func Initialize() error {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open(DB_FILE_NAME, 0600, nil)
	if err != nil {
		log.Fatal(err)
		return err
	}
	// defer db.Close()
	Instance = &database{
		DB: db,
	}

	err = Instance.createUsersBucket()

	return err
}

func Close() {
	Instance.Close()
}

func (d *database) createUsersBucket() error {
	return d.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(BUCKET_NAME))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
}

func (d *database) UpsertUser(mobileNo string, user models.CurrentUser) error {
	return d.Update(func(tx *bolt.Tx) error {
		bytes, err := json.Marshal(user)
		if err != nil {
			return err
		}
		b := tx.Bucket([]byte(BUCKET_NAME))
		err = b.Put([]byte(mobileNo), bytes)
		return err
	})
}

func (d *database) GetUser(mobileNo string) (user *models.CurrentUser, err error) {
	d.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(BUCKET_NAME))
		v := b.Get([]byte(mobileNo))
		user = &models.CurrentUser{}
		err = json.Unmarshal(v, &user)

		return err
	})

	return
}
