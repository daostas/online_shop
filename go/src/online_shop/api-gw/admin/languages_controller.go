package admin

import (
	"encoding/json"
	"fmt"
	"online_shop/admin-svc/pb"
	st "online_shop/status"

	// st "online_shop/status"

	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

	// "github.com/kataras/iris/v12/mvc"

	//"github.com/kataras/iris/v12/x/errors"
	//"google.golang.org/grpc/codes"
	"net"
	//"net/http"
	//"strconv"
	//"time"
)

type AdminLanguagesController struct {
	Client pb.AdminLanguagesClient
	Logger *golog.Logger
	IP     net.IP
	Ctx    iris.Context
}

// PostLogin godoc
// @Summary Получение списка языков в виде таблицы данных
// @Description В поле Filter необходимо добавить поле Format: Если Format будет равен 0, то метод будет работать как дататэйбл для языков, если format равен 1 - метод будет работать в упрощенном режиме и просто вернет список всех языков в виде мап и проигнорирует все остльные данные в request, главное указать формат, чтобы метод работал в упрощенном режиме
// @Tags admin languages
// @Param  DataTableReq body pb.DataTableReq true " "
// @Produce json
// @Success 200 {object} pb.DataTableRes
// @Failure 500 {object} pb.DataTableRes "Ошибка возникающая в методах внутри функции или в базе данных, более подробную информацию об ошибке можно получить внутри DataTableRes в поле Err"
// @Router /admin/get/list/languages [post]
// @Security BearerAuth
func (c *AdminLanguagesController) PostGetListLanguages(ctx iris.Context) *DataTableResponse {
	var req pb.DataTableReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		return &DataTableResponse{
			Error: fmt.Sprint(err),
		}
	}
	res, err := c.Client.GetListOfLanguages(ctx, &req)
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
// @Summary Регистрация языка
// @Description Для регистрации главной группы parent_id должен быть 0, для дочерней группы должен присылаться айди группы, в которую хотим добавить дочернюю группу
// @Tags admin languages
// @Param  NewLangReq body pb.NewLangReq true " "
// @Produce json
// @Success 200 {object} pb.AdminRes "Всё прошло успешно"
// @Failure 432 {object} pb.AdminRes "Ошибка возникающая, если такой язык уже сущесвует"
// @Failure 500 {object} pb.AdminRes "Ошибка возникающая в методах внутри функции или в базе данных, более подробную информацию об ошибке можно получить внутри AdminRes в поле Err"
// @Router /admin/register/languages [post]
// @Security BearerAuth
func (c *AdminLanguagesController) PostRegisterLanguages(ctx iris.Context) *mvc.Response {
	var req pb.NewLangReq
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
	res, err := c.Client.NewLanguage(ctx, &req)
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
// @Summary Смена статуса языка
// @Description ---
// @Tags admin languages
// @Param  ChangeStatusReq body pb.ChangeStatusReq true " "
// @Produce json
// @Success 200 {object} pb.ChangeStatusRes
// @Failure 403 {object} pb.ChangeStatusRes "Ошибка возникающая если админ попробовал выключить язык, который поставлен дефолтным(главным) в админке"
// @Failure 500 {object} pb.ChangeStatusRes "Ошибка возникающая в методах внутри функции или в базе данных, более подробную информацию об ошибке можно получить внутри ChangeStatusRes в поле Err"
// @Router /admin/change/status/languages [post]
// @Security BearerAuth
func (c *AdminLanguagesController) PostChangeStatusLanguages(ctx iris.Context) *mvc.Response {
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
	res, err := c.Client.ChangeLanguageStatus(ctx, &req)
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
