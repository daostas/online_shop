package auth

import (
	"errors"
	"fmt"
	"net/http"
	"online_shop/auth-svc/pb"
	st "online_shop/status"
	"strconv"
	"time"

	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"google.golang.org/grpc/codes"

	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/jwt"
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

type AuthController struct {
	Client pb.AuthClient
	Logger *golog.Logger
	IP     net.IP
	Ctx    iris.Context
}

type ClaimsUser struct {
	ID   int32  `json:"id"`
	Role string `json:"role"`
}

func SetupAuth(app *mvc.Application, cfg *config.Config) {
	client, err := InitAuthClient(cfg)
	if err != nil {
		log.Fatalf("Can't initialize user client: %v", err)
	}
	app.Register(client)
	app.Handle(new(AuthController))
}

func (u *ClaimsUser) GetRaw() (interface{}, error) {
	return u, nil
}
func (u *ClaimsUser) GetAuthorization() (string, error) {
	return "Basic", nil
}
func (u *ClaimsUser) GetAuthorizedAt() (time.Time, error) {
	return time.Now(), context.ErrNotSupported
}
func (u *ClaimsUser) GetID() (string, error) {
	return strconv.Itoa(int(u.ID)), nil
}
func (u *ClaimsUser) GetUsername() (string, error) {
	return "", context.ErrNotSupported
}
func (u *ClaimsUser) GetPassword() (string, error) {
	return "", context.ErrNotSupported
}
func (u *ClaimsUser) GetEmail() (string, error) {
	return "", context.ErrNotSupported
}
func (u *ClaimsUser) GetRoles() ([]string, error) {
	var roles []string
	roles = append(roles, u.Role)
	return roles, nil
}
func (u *ClaimsUser) GetToken() ([]byte, error) {
	return nil, context.ErrNotSupported
}
func (u *ClaimsUser) GetField(_ string) (interface{}, error) {
	return nil, context.ErrNotSupported
}

func InitAuthMiddleware(cfg *config.Config, logger *golog.Logger) context.Handler {
	client, err := InitAuthClient(cfg)
	if err != nil {
		log.Fatalf("Can't initialize auth client: %v", err)
	}
	return func(ctx *context.Context) {
		token := jwt.FromHeader(ctx)
		if token == "" {
			if isAnonymous(ctx.Method(), ctx.RequestPath(false)) {
				logger.Printf("Access granted to anonymous on method %s %s\n", ctx.Method(), ctx.RequestPath(false))
				ctx.Next()
			} else {
				logger.Printf("Access denied to anonymous on method %s %s\n", ctx.Method(), ctx.RequestPath(false))
				ctx.StopWithError(http.StatusUnauthorized, errors.New("Access denied. Need admin permission"))
			}
			return
		}
		res, err := client.Validate(ctx, &pb.ValidateReq{Token: token})
		if err != nil {
			ctx.StopWithError(http.StatusInternalServerError, err)
			return
		}
		if res.Status == int32(codes.OK) {
			claims := ClaimsUser{
				ID:   res.UserId,
				Role: res.Role,
			}
			if !IsAccessGranted(ctx.Method(), ctx.RequestPath(false), &claims) {
				logger.Printf("Access denied for user %s on method %s %s\n", claims.ID, ctx.Method(), ctx.RequestPath(false))
				ctx.StopWithError(http.StatusForbidden, err)
				return
			}
			err = ctx.SetUser(&claims)
			if err != nil {
				ctx.StopWithError(http.StatusInternalServerError, err)
				return
			}
			logger.Printf("Access granted for user %s on method %s %s\n", claims.ID, ctx.Method(), ctx.RequestPath(false))
			ctx.Next()
		}
	}
}

// PostLogin godoc
// @Summary Регистрация пользователя
// @Description Регистрация пользователя
// @Tags auth
// @Param  user body pb.RegReq true " "
// @Produce json
// @Success 200 {object} pb.AuthRes
// @Failure 432 {object} pb.AuthRes "Пользователь с таким логином уже существует"
// @Failure 500 {object} pb.AuthRes "Ошибка возникающая в методах внутри функции или в базе данных, более подробную информацию об ошибке можно получить внутри AuthRes в поле Err"
// @Failure 433 {object} pb.AuthRes "Ошибка возникающая если пользователь ввел ни почту и ни номер"
// @Router /auth/register/user [post]
// @Security BearerAuth
func (c *AuthController) PostRegisterUser(ctx iris.Context) *mvc.Response {
	var req pb.RegReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		fmt.Println(err)
		return &mvc.Response{
			Object: &pb.AuthRes{
				Status: st.StatusInternalServerError,
				Err:    "error in reading data from context"},
			Err:  err,
			Code: st.StatusInternalServerError,
		}
	}

	res, err := c.Client.RegisterUser(ctx, &req)
	if err != nil {
		c.Logger.Errorf("Error registering user: %v", err)
		fmt.Println(err)
		return &mvc.Response{
			Code:   st.StatusInternalServerError,
			Object: nil,
			Err:    err,
		}
	}
	if res.Err != "success" {
		return &mvc.Response{
			Object: res,
			Err:    fmt.Errorf(res.Err),
			Code:   int(res.Status),
		}
	}

	return &mvc.Response{
		Object: res,
		Err:    nil,
		Code:   st.StatusOK,
	}
}

// PostLogin godoc
// @Summary Аутенфикации пользователя
// @Description Аутенфикации пользователя
// @Tags auth
// @Param  user body pb.SignInReq true " "
// @Produce json
// @Success 200 {object} pb.SignInRes
// @Failure 500 {object} pb.SignInRes "Ошибка возникающая в методах внутри функции или в базе данных, более подробную информацию об ошибке можно получить внутри SignInRes в поле Err"
// @Failure 433 {object} pb.SignInRes "Ошибка возникающая,если пользователь ввел ни почту и ни номер"
// @Failure 435 {object} pb.SignInRes "Ошибка возникающая, если пользователь с таким логином не найден"
// @Failure 434 {object} pb.SignInRes "Ошибка возникающая, если пользователь ввел не верный пароль"
// @Router /auth/login/user [post]
// @Security BearerAuth
func (c *AuthController) PostLoginUser(ctx iris.Context) *mvc.Response {
	var req pb.SignInReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		fmt.Println(err)
		return &mvc.Response{
			Object: &pb.SignInRes{
				Status: st.StatusInternalServerError,
				Err:    "error in reading data from context"},
			Err:  err,
			Code: st.StatusInternalServerError,
		}
	}
	res, err := c.Client.SignInUser(ctx, &req)

	if err != nil {
		c.Logger.Errorf("Error authentification user: %v", err)
		fmt.Println(err)
		fmt.Println(err)
		return &mvc.Response{
			Code:   st.StatusInternalServerError,
			Object: nil,
			Err:    err,
		}
	}
	if res.Err != "success" {
		return &mvc.Response{
			Object: res,
			Err:    errors.New(res.Err),
			Code:   int(res.Status),
		}
	}
	return &mvc.Response{
		Object: res,
		Err:    nil,
		Code:   st.StatusOK,
	}
}

// PostLogin godoc
// @Summary Регистрация админа
// @Description Регистрация админа
// @Tags auth
// @Param  admin body pb.RegReq true " "
// @Produce json
// @Success 200 {object} pb.AuthRes
// @Failure 500 {object} pb.AuthRes "Ошибка возникающая в методах внутри функции или в базе данных, более подробную информацию об ошибке можно получить внутри AuthRes в поле Err"
// @Failure 432 {object} pb.AuthRes "Ошибка возникающая, если админ с таким логином уже существует"
// @Failure 433 {object} pb.AuthRes "Ошибка возникающая, если пользователь ввел ни почту и ни пароль"
// @Router /auth/register/admin [post]
// @Security BearerAuth
func (c *AuthController) PostRegisterAdmin(ctx iris.Context) *mvc.Response {
	var req pb.RegReqAdmin
	err := ctx.ReadJSON(&req)
	if err != nil {
		fmt.Println(err)
		return &mvc.Response{
			Object: &pb.AuthRes{
				Status: st.StatusInternalServerError,
				Err:    "error in reading data from context"},
			Err:  err,
			Code: st.StatusInternalServerError,
		}
	}
	res, err := c.Client.RegisterAdmin(ctx, &req)
	if err != nil {
		c.Logger.Errorf("Error registering user: %v", err)
		return &mvc.Response{
			Code:   st.StatusInternalServerError,
			Object: nil,
			Err:    err,
		}
	}
	if res.Err != "success" {
		return &mvc.Response{
			Object: res,
			Err:    fmt.Errorf(res.Err),
			Code:   int(res.Status),
		}
	}

	return &mvc.Response{
		Object: res,
		Err:    nil,
		Code:   st.StatusOK,
	}
}

func (c *AuthController) PostRegisterMadmin(ctx iris.Context) *mvc.Response {
	var req pb.RegReqAdmin
	err := ctx.ReadJSON(&req)
	if err != nil {
		return &mvc.Response{
			Object: &pb.AuthRes{
				Status: st.StatusInternalServerError,
				Err:    "error in reading data from context"},
			Err:  err,
			Code: st.StatusInternalServerError,
		}
	}

	res, err := c.Client.RegisterMainAdmin(ctx, &req)
	if err != nil {
		c.Logger.Errorf("Error registering user: %v", err)
		return &mvc.Response{
			Code:   st.StatusInternalServerError,
			Object: nil,
			Err:    err,
		}
	}

	if res.Err != "success" {
		return &mvc.Response{
			Object: res,
			Err:    fmt.Errorf(res.Err),
			Code:   int(res.Status),
		}
	}

	return &mvc.Response{
		Object: res,
		Err:    nil,
		Code:   st.StatusOK,
	}
}

// PostLogin godoc
// @Summary Аутенфикации админа
// @Description Аутенфикации админа
// @Tags auth
// @Param  user body pb.SignInReq true " "
// @Produce json
// @Success 200 {object} pb.SignInRes
// @Failure 500 {object} pb.SignInRes "Ошибка возникающая в методах внутри функции или в базе данных, более подробную информацию об ошибке можно получить внутри SignInRes в поле Err"
// @Failure 435 {object} pb.SignInRes "Ошибка возникающая, если админ с таким логином не найден"
// @Failure 434 {object} pb.SignInRes "Ошибка возникающая, если пользователь ввел не верный пароль"
// @Router /auth/login/admin [post]
// @Security BearerAuth
func (c *AuthController) PostLoginAdmin(ctx iris.Context) *mvc.Response {
	var req pb.SignInReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		return &mvc.Response{
			Object: &pb.SignInRes{
				Status: st.StatusInternalServerError,
				Err:    "error in reading data from context"},
			Err:  err,
			Code: st.StatusInternalServerError,
		}
	}
	res, err := c.Client.SignInAdmin(ctx, &req)
	if err != nil {
		c.Logger.Errorf("Error authentification user: %v", err)
		return &mvc.Response{
			Code:   st.StatusInternalServerError,
			Object: nil,
			Err:    err,
		}
	}

	if res.Err != "success" {
		return &mvc.Response{
			Object: res,
			Err:    fmt.Errorf(res.Err),
			Code:   int(res.Status),
		}
	}

	return &mvc.Response{
		Object: res,
		Err:    nil,
		Code:   st.StatusOK,
	}
}
