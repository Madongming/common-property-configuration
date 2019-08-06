package property

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/jsimonetti/berkeleydb"
)

type Storage struct{}

func NewDB() *berkeleydb.Db {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err.Error())
	}

	dir = "/Users/madongming"
	db, err := berkeleydb.NewDB()
	if err != nil {
		panic(err)
	}
	if err := db.Open(dir+"/data.db", berkeleydb.DbHash, berkeleydb.DbCreate); err != nil {
		panic(err)
	}
	return db
}

func (s *Storage) Create(key string, data []byte) error {
	db := NewDB()
	defer db.Close()

	dataStr := string(data) + "\n"
	if err := db.Put(key, dataStr); err != nil {
		return err
	}

	return nil
}

func (s *Storage) Update(key string, data []byte) error {
	db := NewDB()
	defer db.Close()

	dataStr := string(data) + "\n"
	if err := db.Put(key, dataStr); err != nil {
		return err
	}
	return nil
}

func (s *Storage) Read(key string) ([]byte, error) {
	db := NewDB()
	defer db.Close()

	value, err := db.Get(key)
	if err != nil {
		return nil, err
	}

	return []byte(strings.TrimSpace(value)), nil
}

func (s *Storage) Delete(key string) ([]byte, error) {
	db := NewDB()
	defer db.Close()

	value, err := db.Get(key)
	if err != nil {
		return nil, err
	}

	delBytes := []byte(strings.TrimSpace(value))

	if err := db.Delete(key); err != nil {
		return nil, err
	}
	return delBytes, nil
}
