package admin

import (
	"encoding/json"
	"fmt"
	"online_shop/admin-svc/pb"
	st "online_shop/status"

	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"

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

type AdminGroupsController struct {
	Client pb.AdminGroupsClient
	Logger *golog.Logger
	IP     net.IP
	Ctx    iris.Context
}

func SetupAdmin(app *mvc.Application, cfg *config.Config) {
	GroupsСlient, err := InitAdminGroupsClient(cfg)
	if err != nil {
		log.Fatalf("Can't initialize user client: %v", err)
	}
	app.Register(GroupsСlient)
	app.Handle(new(AdminGroupsController))

	ProductsСlient, err := InitProductsClient(cfg)
	if err != nil {
		log.Fatalf("Can't initialize user client: %v", err)
	}
	app.Register(ProductsСlient)
	app.Handle(new(ProductsController))

	ProducersСlient, err := InitProducersClient(cfg)
	if err != nil {
		log.Fatalf("Can't initialize user client: %v", err)
	}
	app.Register(ProducersСlient)
	app.Handle(new(ProducersController))

}

type DataTableResponse struct {
	Draw            int              `form:"draw" json:"draw"`
	Recordstotal    int              `form:"recordsTotal" json:"recordsTotal"`
	Recordsfiltered int              `form:"recordsFiltered" json:"recordsFiltered"`
	Data            []map[string]any `form:"data" json:"data"`
	Error           string           `form:"error" json:"error"`
}

// PostLogin godoc
// @Summary Регистрация пользователя
// @Description Регистрация пользователя
// @Tags admin groups
// @Param  admin body pb.RegGroupReq true " "
// @Produce json
// @Success 200 {object} pb.AdminRes "Всё прошло успешно"
// @Success 209 {object} pb.AdminRes "Прошло успешно, но есть warning, потому что группа с таким названием уже существует"
// @Failure 500 {object} pb.AdminRes "Ошибка возникающая в методах внутри функции или в базе данных, более подробную информацию об ошибке можно получить внутри SignInRes в поле Err"
// @Failure 433 {object} pb.AdminRes "Ошибка возникающая при передаче неправильных данных в localizations"
// @Router /admin/register/groups [post]
// @Security BearerAuth
func (c *AdminGroupsController) PostRegisterGroups(ctx iris.Context) *mvc.Response {
	var req pb.RegGroupReq
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
	res, err := c.Client.RegisterGroup(ctx, &req)
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
// @Tags admin groups
// @Param  admin body pb.DataTableReq true " "
// @Produce json
// @Success 200 {object} pb.DataTableRes
// @Failure 500 {object} pb.DataTableRes "Ошибка возникающая в методах внутри функции или в базе данных, более подробную информацию об ошибке можно получить внутри SignInRes в поле Err"
// @Router /admin/get/list/groups [post]
// @Security BearerAuth
func (c *AdminGroupsController) PostGetListGroups(ctx iris.Context) *DataTableResponse {
	var req pb.DataTableReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		return &DataTableResponse{
			Error: fmt.Sprint(err),
		}
	}
	res, err := c.Client.GetListOfGroups(ctx, &req)
	if err != nil {
		c.Logger.Errorf("Error registering user: %v", err)
		return &DataTableResponse{Error: fmt.Sprint(err)}

	}
	if res.Err != "success" {
		return &DataTableResponse{Error: res.Err}
	}
	var result DataTableResponse
	err = json.Unmarshal(res.Data, &result)
	if err != nil {
		return &DataTableResponse{
			Error: fmt.Sprint(err),
		}
	}
	result.Error = ""
	return &result
}

// PostLogin godoc
// @Summary Регистрация пользователя
// @Description Регистрация пользователя
// @Tags user
// @Param  user body pb.DataTableReq true " "
// @Produce json
// @Success 200 {object} pb.ChangeStatusReq
// @Failure 500 {object} pb.ChangeStatusRes "Ошибка возникающая в методах внутри функции или в базе данных, более подробную информацию об ошибке можно получить внутри SignInRes в поле Err"
// @Router /admin/groups/change/status [post]
// @Security BearerAuth
func (c *AdminGroupsController) PostGroupsChangeStatus(ctx iris.Context) *mvc.Response {
	var req pb.ChangeStatusReq
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
	res, err := c.Client.ChangeGroupsStatus(ctx, &req)
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
