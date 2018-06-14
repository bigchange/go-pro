package gotest

import (
	"fmt"
	"testing"
	"github.com/syndtr/goleveldb/leveldb"
)

func TestLevelDB(t *testing.T) {
	db, err := leveldb.OpenFile("../data/leveldb", nil)
	if err != nil {
		fmt.Println("get data error", err.Error())
	}
	// Remember that the contents of the returned slice should not be modified.
	data, err := db.Get([]byte("key"), nil)
	if err == nil {
		fmt.Println("get data:", string(data))
	} else {
		fmt.Println("get data error", err.Error())
	}
	err = db.Put([]byte("key"), []byte("value"), nil)
	data, err = db.Get([]byte("key"), nil)
	if err == nil {
		fmt.Println("get data:", string(data))
	} else {
		fmt.Println("get data error", err.Error())
	}
	err = db.Delete([]byte("key"), nil)
	fmt.Println("get data:", string(data))
}