package badgerdb


import (
	"github.com/bigchange/go-pro/myproject/utils"
	"github.com/dgraph-io/badger"
	"log"
)

var db *badger.DB

func InitBadgerDB(config *utils.LLBConfig) {
	// Open the Badger database located in the specific directory.
  // It will be created if it doesn't exist.
  opts := badger.DefaultOptions
  opts.Dir = config.BadgerDir
	opts.ValueDir = config.BadgerDir
	var err error
	// just open once before run db.Close()
  db, err = badger.Open(opts)
  if err != nil {
	  log.Printf("open db error:", err.Error())
	}
}
func GetDB() *badger.DB {
	return db
}