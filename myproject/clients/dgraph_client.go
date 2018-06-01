/*
 * @Author: Jerry You 
 * @CreatedDate: 2018-05-14 19:25:34 
 * @Last Modified by: Jerry You
 * @Last Modified time: 2018-05-14 20:16:53
 */
package clients

import (
	"log"
	"google.golang.org/grpc"
	"github.com/bgfurfeature/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	// api_go "higgs.com/inmind/idmg/lieluobo/dynamic_features/proto/api_proto_go"
)

type DClient struct {
	Conns	[]*grpc.ClientConn
	Address	[]string
	Stub *dgo.Dgraph
}

func NewDClients(address []string) (client *DClient, err error) {
	client = &DClient{
		Conns: make([]*grpc.ClientConn, len(address)), 
		Address: address, 
		Stub: nil}
	return client, nil
}

func NewDClient(address ...string) (client *DClient, err error) {
	client = &DClient{
		Conns: make([]*grpc.ClientConn, len(address)), 
		Address: address, 
		Stub: nil}
	return client, nil
}

func (client DClient) Dial()(err error) {
	log.Println("address length:", len(client.Address))
	var size = len(client.Address)
	var dc = make([]api.DgraphClient, len(client.Address))
	for i := 0; i < size; i++ {
		conn, err := grpc.Dial(client.Address[i], grpc.WithInsecure())
		client.Conns[i] = conn
		dc[i] = api.NewDgraphClient(client.Conns[i])
		if err != nil {
			log.Fatal("While trying to dial gRPC")
			return err
		}
	}
	client.Stub = dgo.NewDgraphClients(dc)
	return nil
}
func (client DClient) Close() {
	for _, conn := range client.Conns {
		defer conn.Close()
	}
}

