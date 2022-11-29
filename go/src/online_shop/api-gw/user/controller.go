package user

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"online_shop/user-svc/pb"
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

type Controller struct {
	Client pb.UserServiceClient
	Logger *golog.Logger
	IP     net.IP
	Ctx    iris.Context
}

func SetupUser(app *mvc.Application, cfg *config.Config) {
	client, err := InitUserClient(cfg)
	if err != nil {
		log.Fatalf("Can't initialize user client: %v", err)
	}
	app.Register(client)
	app.Handle(new(Controller))
}

// PostLogin godoc
// @Summary Регистрация пользователя
// @Description Регистрация пользователя
// @Tags user
// @Param  user body pb.RegUserReq true " "
// @Produce json
// @Success 200 {object} pb.UserRes
// @Failure 500 {object} string "error"
// @Router /user/reguster/user [post]
// @Security BearerAuth
func (c *Controller) PostRegisterUser(ctx iris.Context) *mvc.Response {
	var req pb.RegUserReq
	err := ctx.ReadJSON(&req)
	res, err := c.Client.RegisterUser(ctx, &req)
	if err != nil {
		c.Logger.Errorf("Error registering user: %v", err)
		return &mvc.Response{
			Code:   iris.StatusInternalServerError,
			Object: nil,
			Err:    err,
		}
	}
	return &mvc.Response{
		Object: res,
		Err:    nil,
		Code:   iris.StatusOK,
	}
}

func (c *Controller) Get(ctx iris.Context) *mvc.Response {
	var req pb.SignInReq
	req = pb.SignInReq{
		Login:    "daostas@gmail.com",
		Password: "password",
	}
	res, err := c.Client.SignInUser(ctx, &req)
	if err != nil {
		c.Logger.Errorf("Error authentification user: %v", err)
		return &mvc.Response{
			Code:   iris.StatusInternalServerError,
			Object: nil,
			Err:    err,
		}
	}
	return &mvc.Response{
		Object: res,
		Err:    nil,
		Code:   iris.StatusOK,
	}
}

// PostLogin godoc
// @Summary Аутенфикации пользователя
// @Description Аутенфикации пользователя
// @Tags user
// @Param  user body pb.SignInReq true " "
// @Produce json
// @Success 200 {object} pb.UserRes
// @Failure 500 {object} string "error"
// @Router /user/login [post]
// @Security BearerAuth
func (c *Controller) PostLogin(ctx iris.Context) *mvc.Response {
	var req pb.SignInReq
	err := ctx.ReadJSON(&req)
	res, err := c.Client.SignInUser(ctx, &req)
	if err != nil {
		c.Logger.Errorf("Error authentification user: %v", err)
		return &mvc.Response{
			Code:   iris.StatusInternalServerError,
			Object: nil,
			Err:    err,
		}
	}
	return &mvc.Response{
		Object: res,
		Err:    nil,
		Code:   iris.StatusOK,
	}
}
