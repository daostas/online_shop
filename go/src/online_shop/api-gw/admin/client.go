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

func InitAdminProductsClient(c *config.Config) (pb.AdminProductsClient, error) {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.AdminSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return pb.NewAdminProductsClient(cc), nil

}

func InitAdminProducersClient(c *config.Config) (pb.AdminProducersClient, error) {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.AdminSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return pb.NewAdminProducersClient(cc), nil

}

func InitAdminLanguagesClient(c *config.Config) (pb.AdminLanguagesClient, error) {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.AdminSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return pb.NewAdminLanguagesClient(cc), nil

}

func InitAdminParametrsClient(c *config.Config) (pb.AdminParametrsClient, error) {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.AdminSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return pb.NewAdminParametrsClient(cc), nil

}
