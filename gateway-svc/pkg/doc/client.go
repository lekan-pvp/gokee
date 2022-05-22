package doc

import (
	"fmt"
	"github.com/lekan-pvp/gokee/gateway-svc/pkg/config"
	"github.com/lekan-pvp/gokee/gateway-svc/pkg/doc/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.DocServiceClient
}

func InitServiceClient(c *config.Config) pb.DocServiceClient {
	cc, err := grpc.Dial(c.DocSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect: ", err)
	}
	return pb.NewDocServiceClient(cc)
}
