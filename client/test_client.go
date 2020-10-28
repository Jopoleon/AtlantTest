package client

import (
	"context"
	"fmt"
	"time"

	"github.com/k0kubun/pp"

	"github.com/Jopoleon/AtlantTest/server/proto_server"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type AtlantTestClient struct {
	proto_server.ProductServiceClient
}

func NewTestClient(host, port string) *AtlantTestClient {
	target := fmt.Sprintf("%s:%s", host, port)
	conn, err := grpc.Dial(
		target,
		grpc.WithInsecure(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                (time.Duration(10) * time.Second),
			Timeout:             (time.Duration(7) * time.Second),
			PermitWithoutStream: true,
		}),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1000000000)),
	)
	if err != nil {
		logrus.Fatal(err)
	}

	c := proto_server.NewProductServiceClient(conn)
	return &AtlantTestClient{
		c,
	}
}

func Test() {
	c := NewTestClient("localhost", "4000")
	c.TestFetch()
	c.TestList()
}

func (a *AtlantTestClient) TestFetch() {
	reply, err := a.Fetch(context.TODO(), &proto_server.FetchRequest{
		Url: "http://localhost:8888",
	})
	if err != nil {
		logrus.Fatal(err)
	}
	pp.Println(reply.Code)
	pp.Println(reply.Message)
}

func (a *AtlantTestClient) TestList() {
	lr := &proto_server.ListRequest{
		Page:  1,
		Limit: 10,
		Sort:  2,
		Order: -1,
	}
	reply, err := a.List(context.TODO(), lr)
	if err != nil {
		logrus.Fatal(err)
	}
	if len(reply.Products) != int(lr.Limit) {
		pp.Println("WRONG NUMBER OF PRODUCTS")
	}

	for _, p := range reply.Products {
		pp.Println(p.Name)
		pp.Println(p.LastPrice)
		pp.Println(p.PriceChanges)
		pp.Println(p.LastUpdated)
		fmt.Println("----")
	}

}
