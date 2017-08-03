package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql" //import postgres
	"bigchange.com/go-pro/db"
)

//DB ...
type DB struct {
	*sql.DB
}

const (
	//DbUser ...
	DbUser = "casem"
	//DbPassword ...
	DbPassword = "Casem123@"
	//DbName ...
	DbName = "casem"

	DbHost = "172.16.52.52"
)

var db *gorp.DbMap

//Init ...
func Init() {

	// dbinfo := fmt.Sprintf("tcp:%s:3306*%s/%s/%s",
	// 	DbHost, DbName, DbUser, DbPassword)

	dbinfo := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s",
		DbUser, DbPassword, DbHost, DbName)

	var err error
	db, err = ConnectDB(dbinfo)
	if err != nil {
		utils.GetLogger().Fatal(err)
	}

}

//ConnectDB ...
func ConnectDB(dataSourceName string) (*gorp.DbMap, error) {
	utils.GetLogger().Infof("connect to :%v ", dataSourceName)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "casem:", log.Lmicroseconds)) //Trace database requests
	return dbmap, nil
}

//GetDB ...
func GetDB() *gorp.DbMap {
	return db
}