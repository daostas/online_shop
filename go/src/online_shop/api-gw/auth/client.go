package auth

import (
	//"fmt"
	"online_shop/api-gw/config"
	"online_shop/auth-svc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitAuthClient(c *config.Config) (pb.AuthClient, error) {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.ClientSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return pb.NewAuthClient(cc), nil
}
