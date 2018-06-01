package gotest

import (
	"github.com/bigchange/go-pro/myproject/utils"
	"github.com/bigchange/go-pro/myproject/clients"
	"testing"
)


func initDgraphClient(address string) {
	dclients, err := clients.NewDClient(address)
	err = dclients.Dial()
	if err != nil {
		utils.GetLogger().Criticalf("can't create dgraph client")
	}
	utils.GetLogger().Info("create dgraph client success")
	defer dclients.Close()
}


func TestInitDgraphClient(t *testing.T)  {
	initDgraphClient("172.20.0.8:9080")
}
