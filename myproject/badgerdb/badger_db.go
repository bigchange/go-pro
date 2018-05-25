package badger


import (
	"higgs.com/inmind/idmg/lieluobo/dynamic_features/utils"
	"github.com/dgraph-io/badger"
)

type DB struct {
	db *badger.DB
}

func InitBadgerDB(config *utils.LLBConfig) {
	// Open the Badger database located in the /tmp/badger directory.
  // It will be created if it doesn't exist.
  opts := badger.DefaultOptions
  opts.Dir = config.BadgerDir
  opts.ValueDir = config.BadgerDir
  db, err := badger.Open(opts)
  if err != nil {
	  log.Fatal(err)
	}
	// DB := &DB {db: db}
}

func(db DB) Close() {
	db.db.Close()
}

func GetDB() *badger.DB {
	return db
}