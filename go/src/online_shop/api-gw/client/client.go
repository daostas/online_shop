package client

import (
	//"fmt"
	"online_shop/api-gw/config"
	"online_shop/client-svc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitClientsClient(c *config.Config) (pb.ClientsClient, error) {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.ClientSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return pb.NewClientsClient(cc), nil
}

func InitClientGroupsClient(c *config.Config) (pb.ClientGroupsClient, error) {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.ClientSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return pb.NewClientGroupsClient(cc), nil
}

func InitClientLanguagesClient(c *config.Config) (pb.ClientLanguagesClient, error) {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.ClientSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return pb.NewClientLanguagesClient(cc), nil
}
