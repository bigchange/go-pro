package gotest

import (
	"log"
	"context"
	"github.com/bigchange/go-pro/myproject/utils"
	"github.com/bigchange/go-pro/myproject/clients"
	"testing"
)

var dclient *clients.DClient

func init() {
}

func initDgraphClient(address string) {
	dclient, err := clients.NewDClient(address)
	err = dclient.Dial()
	if err != nil {
		utils.GetLogger().Criticalf("can't create dgraph client")
	}
	utils.GetLogger().Info("create dgraph client success")
}

func TestInitDgraphClient(t *testing.T)  {
	// initDgraphClient("")
}
