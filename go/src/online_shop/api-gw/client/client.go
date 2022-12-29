package client

import (
	//"fmt"
	"online_shop/api-gw/config"
	"online_shop/client-svc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitUsersClient(c *config.Config) (pb.UsersClient, error) {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.ClientSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return pb.NewUsersClient(cc), nil
}
