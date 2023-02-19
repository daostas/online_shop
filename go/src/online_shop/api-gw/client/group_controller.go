package client

import (
	"fmt"
	"online_shop/client-svc/pb"

	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"

	//"github.com/kataras/iris/v12/context"
	//"github.com/kataras/iris/v12/middleware/jwt"
	"github.com/kataras/iris/v12/mvc"
	//"github.com/kataras/iris/v12/x/errors"
	//"google.golang.org/grpc/codes"
	"net"
	st "online_shop/status"
	//"net/http"
	//"strconv"
	//"time"
)

type ClientGroupsController struct {
	Client pb.ClientGroupsClient
	Logger *golog.Logger
	IP     net.IP
	Ctx    iris.Context
}

// PostLogin godoc
// @Summary Получение списка групп
// @Description Чтобы получить список главных групп, поле group_id должно быть равно нулю. Для получение подгрупп тебе нужно отправить в group_id айди нужной группы
// @Tags client groups
// @Param  GetGroupsReq body pb.GetGroupsReq true " "
// @Produce json
// @Success 200 {object} pb.GetGroupsRes "Всё прошло успешно"
// @Failure 500 {object} pb.GetGroupsRes "Ошибка возникающая в методах внутри функции или в базе данных, более подробную информацию об ошибке можно получить внутри GetGroupsRes в поле Err"
// @Router /client/get/groups [post]
// @Security BearerAuth
func (c *ClientGroupsController) PostGetGroups(ctx iris.Context) *mvc.Response {
	var req pb.GetGroupsReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		return &mvc.Response{
			Object: &pb.GetGroupsRes{
				Status: st.StatusInternalServerError,
				Err:    "error in reading data from context"},
			Err:  err,
			Code: st.StatusInternalServerError,
		}
	}
	res, err := c.Client.GetGroups(ctx, &req)
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
