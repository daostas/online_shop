package user

import (
	//"fmt"
	"online_shop/api-gw/config"
	"online_shop/user-svc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitUserClient(c *config.Config) (pb.UserServiceClient, error) {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.UserSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return pb.NewUserServiceClient(cc), nil
}
