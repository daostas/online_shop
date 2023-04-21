package client_service

import (
	"context"
	"fmt"
	"online_shop/client-svc/config"
	"online_shop/client-svc/pb"
	"online_shop/repository/models"
	st "online_shop/status"

	"gopkg.in/reform.v1"
)

type ClientLanguagesServer struct {
	pb.UnimplementedClientLanguagesServer
	Db  *reform.DB
	Cfg *config.Config
}

func NewClientLanguagesServer(db *reform.DB, cfg *config.Config) *ClientLanguagesServer {
	return &ClientLanguagesServer{
		Db:  db,
		Cfg: cfg,
	}
}

func (s *ClientLanguagesServer) GetLanguages(_ context.Context, req *pb.GetLanguagesReq) (*pb.GetLanguagesRes, error) {

	langs, err := s.Db.SelectAllFrom(models.LanguagesTable, "where status = true")
	if err != nil {
		return &pb.GetLanguagesRes{
			Languages: nil,
			Status:    st.StatusInternalServerError,
			Err:       "error in getting data from languages table: " + fmt.Sprint(err),
		}, nil
	}

	var res []*pb.GetLanguagesRes_Language
	for _, l := range langs {
		var lang pb.GetLanguagesRes_Language
		lang.LangId = l.(*models.Languages).LangID
		lang.Code = l.(*models.Languages).Code
		lang.Image = *l.(*models.Languages).Image
		lang.Locale = l.(*models.Languages).Locale
		lang.Name = l.(*models.Languages).LangName

		res = append(res, &lang)
	}

	return &pb.GetLanguagesRes{
		Languages: res,
		Status:    st.StatusOK,
		Err:       "success",
	}, nil
}
