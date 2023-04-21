package admin

import (
	"encoding/json"
	"fmt"
	"online_shop/admin-svc/pb"
	st "online_shop/status"

	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"

	"github.com/kataras/iris/v12/mvc"

	"log"
	"net"
	"online_shop/api-gw/config"
)

type AdminGroupsController struct {
	Client pb.AdminGroupsClient
	Logger *golog.Logger
	IP     net.IP
	Ctx    iris.Context
}

func SetupAdmin(app *mvc.Application, cfg *config.Config) {
	GroupsClient, err := InitAdminGroupsClient(cfg)
	if err != nil {
		log.Fatalf("Can't initialize user client: %v", err)
	}
	app.Register(GroupsClient)
	app.Handle(new(AdminGroupsController))

	ProductsClient, err := InitAdminProductsClient(cfg)
	if err != nil {
		log.Fatalf("Can't initialize user client: %v", err)
	}
	app.Register(ProductsClient)
	app.Handle(new(AdminProductsController))

	ProducersClient, err := InitAdminProducersClient(cfg)
	if err != nil {
		log.Fatalf("Can't initialize user client: %v", err)
	}
	app.Register(ProducersClient)
	app.Handle(new(AdminProducersController))

	LanguagesClient, err := InitAdminLanguagesClient(cfg)
	if err != nil {
		log.Fatalf("Can't initialize user client: %v", err)
	}
	app.Register(LanguagesClient)
	app.Handle(new(AdminLanguagesController))

	ParametrsClient, err := InitAdminParametrsClient(cfg)
	if err != nil {
		log.Fatalf("Can't initialize user client: %v", err)
	}
	app.Register(ParametrsClient)
	app.Handle(new(AdminParametrsController))

}

type DataTableResponse struct {
	Draw            int              `form:"draw" json:"draw"`
	Recordstotal    int              `form:"recordsTotal" json:"recordsTotal"`
	Recordsfiltered int              `form:"recordsFiltered" json:"recordsFiltered"`
	Data            []map[string]any `form:"data" json:"data"`
	Error           string           `form:"error" json:"error"`
}

// PostLogin godoc
// @Summary Регистрация группы товаров
// @Description Для регистрации главной группы parent_id должен быть 0, для дочерней группы должен присылаться айди группы, в которую хотим добавить дочернюю группу
// @Tags admin groups
// @Param  RegGroupReq body pb.RegGroupReq true " "
// @Produce json
// @Success 200 {object} pb.AdminRes "Всё прошло успешно"
// @Success 209 {object} pb.AdminRes "Прошло успешно, но есть warning, потому что группа с таким названием уже существует"
// @Failure 500 {object} pb.AdminRes "Ошибка возникающая в методах внутри функции или в базе данных, более подробную информацию об ошибке можно получить внутри AdminRes в поле Err"
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
// @Summary Получение списка групп в виде таблицы данных
// @Description В поле Filter необходимо добавить поле Format: Если Format будет равен 0, то метод будет работать как дататэйбл для групп, если format равен 1 - метод будет работать в упрощенном режиме и просто вернет список всех языков в виде мап и проигнорирует все остльные данные в request, главное указать формат, чтобы метод работал в упрощенном режиме; А также необходимо добавить поле lang_id для получения данных на определенном языке, получить список языков можно методом get/list/languages
// @Tags admin groups
// @Param  DataTableReq body pb.DataTableReq true " "
// @Produce json
// @Success 200 {object} pb.DataTableRes
// @Failure 500 {object} pb.DataTableRes "Ошибка возникающая в методах внутри функции или в базе данных, более подробную информацию об ошибке можно получить внутри DataTableRes в поле Err"
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
// @Summary Смена статуса группы
// @Description ---
// @Tags admin groups
// @Param  ChangeStatusReq body pb.ChangeStatusReq true " "
// @Produce json
// @Success 200 {object} pb.ChangeStatusRes
// @Failure 500 {object} pb.ChangeStatusRes "Ошибка возникающая в методах внутри функции или в базе данных, более подробную информацию об ошибке можно получить внутри ChangeStatusRes в поле Err"
// @Router /admin/change/status/groups [post]
// @Security BearerAuth
func (c *AdminGroupsController) PostChangeStatusGroups(ctx iris.Context) *mvc.Response {
	var req pb.ChangeStatusReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		return &mvc.Response{
			Object: &pb.ChangeStatusRes{
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
