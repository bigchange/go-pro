package main

import (
	"log"
	"github.com/bigchange/go-pro/myproject/badgerdb"
	"github.com/bigchange/go-pro/myproject/utils"
)

func write_badger_db() {
	
}

func main() {
	// test manage txt by self
	config := &utils.LLBConfig{
		BadgerDir: "../data/db"}
	badgerdb.InitBadgerDB(config)
	db := badgerdb.GetDB()
	txn := db.NewTransaction(true)
	defer db.Close()
	item, err := txn.Get([]byte("answer"))
	if err != nil {
		log.Println("err", err.Error())
	}
	val, err := item.Value()
	if err != nil {
		log.Println("err", err.Error())
	}
	log.Println("value", string(val))
	defer txn.Discard()
}