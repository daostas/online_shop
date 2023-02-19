package admin

import (
	"fmt"
	"online_shop/admin-svc/pb"
	st "online_shop/status"

	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"

	"github.com/kataras/iris/v12/mvc"

	//"github.com/kataras/iris/v12/x/errors"
	//"google.golang.org/grpc/codes"
	"net"
	//"net/http"
	//"strconv"
	//"time"
)

type ProducersController struct {
	Client pb.ProducersClient
	Logger *golog.Logger
	IP     net.IP
	Ctx    iris.Context
}

// PostLogin godoc
// @Summary Регистрация пользователя
// @Description Регистрация пользователя
// @Tags admin producers
// @Param  admin body pb.RegProducerReq true " "
// @Produce json
// @Success 200 {object} pb.AdminRes "Всё прошло успешно"
// @Success 209 {object} pb.AdminRes "Прошло успешно, но есть warning, потому что группа с таким названием уже существует"
// @Failure 500 {object} pb.AdminRes "Ошибка возникающая в методах внутри функции или в базе данных, более подробную информацию об ошибке можно получить внутри SignInRes в поле Err"
// @Failure 433 {object} pb.AdminRes "Ошибка возникающая при передаче неправильных данных в localizations"
// @Router /admin/register/producers [post]
// @Security BearerAuth
func (c *ProducersController) PostRegisterProducers(ctx iris.Context) *mvc.Response {
	var req pb.RegProducerReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		return &mvc.Response{
			Object: &pb.AdminRes{
				Status: st.StatusInternalServerError,
				Err:    "error in reading data from context"},
			Err:  err,
			Code: st.StatusInternalServerError,
		}
	}
	res, err := c.Client.RegisterProducer(ctx, &req)
	if err != nil {
		c.Logger.Errorf("Error registering user: %v", err)
		return &mvc.Response{
			Code:   iris.StatusInternalServerError,
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
		Code:   iris.StatusOK,
	}
}

// PostLogin godoc
// @Summary Регистрация пользователя
// @Description Регистрация пользователя
// @Tags admin producers
// @Param  admin body pb.DataTableReq true " "
// @Produce json
// @Success 200 {object} pb.DataTableReq
// @Failure 500 {object} pb.DataTableReq "Ошибка возникающая в методах внутри функции или в базе данных, более подробную информацию об ошибке можно получить внутри SignInRes в поле Err"
// @Router /admin/get/list/producers [post]
// @Security BearerAuth
// func (c *ProducersController) PostGetListProducers(ctx iris.Context) *DataTableResponse {
// 	var req pb.DataTableReq
// 	err := ctx.ReadJSON(&req)
// 	if err != nil {
// 		return &DataTableResponse{
// 			Error: fmt.Sprint(err),
// 		}
// 	}
// 	res, err := c.Client.GetListOfProducers(ctx, &req)
// 	if err != nil {
// 		c.Logger.Errorf("Error registering user: %v", err)
// 		return &DataTableResponse{Error: fmt.Sprint(err)}

// 	}
// 	if res.Err != "success" {
// 		return &DataTableResponse{Error: res.Err}
// 	}
// 	var result DataTableResponse
// 	err = json.Unmarshal(res.Data, &result)
// 	if err != nil {
// 		return &DataTableResponse{
// 			Error: fmt.Sprint(err),
// 		}
// 	}
// 	result.Error = ""
// 	return &result
// }