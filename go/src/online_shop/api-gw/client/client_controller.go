package client

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"online_shop/client-svc/pb"
	//"github.com/kataras/iris/v12/context"
	//"github.com/kataras/iris/v12/middleware/jwt"
	"github.com/kataras/iris/v12/mvc"
	//"github.com/kataras/iris/v12/x/errors"
	"online_shop/api-gw/config"
	//"google.golang.org/grpc/codes"
	"log"
	"net"
	//"net/http"
	//"strconv"
	//"time"
)

type ClientController struct {
	Client pb.ClientsClient
	Logger *golog.Logger
	IP     net.IP
	Ctx    iris.Context
}

func SetupClient(app *mvc.Application, cfg *config.Config) {
	client, err := InitClientsClient(cfg)
	if err != nil {
		log.Fatalf("Can't initialize user client: %v", err)
	}
	app.Register(client)
	app.Handle(new(ClientController))

	groupsclient, err := InitClientGroupsClient(cfg)
	if err != nil {
		log.Fatalf("Can't initialize user client: %v", err)
	}
	app.Register(groupsclient)
	app.Handle(new(ClientGroupsController))
}
