package admin

import (
	"online_shop/admin-svc/pb"
	//"fmt"
	"online_shop/api-gw/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitAdminGroupsClient(c *config.Config) (pb.AdminGroupsClient, error) {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.AdminSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return pb.NewAdminGroupsClient(cc), nil

}

func InitProductsClient(c *config.Config) (pb.ProductsClient, error) {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.AdminSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return pb.NewProductsClient(cc), nil

}

func InitProducersClient(c *config.Config) (pb.ProducersClient, error) {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.AdminSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return pb.NewProducersClient(cc), nil

}
