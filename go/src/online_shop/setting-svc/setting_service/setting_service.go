package setting_service

import (
	"context"
	rep "online_shop/repository"
	"online_shop/setting-svc/config"
	"online_shop/setting-svc/pb"
	"strconv"

	"gopkg.in/reform.v1"
)

type SettingServiceServer struct {
	pb.UnimplementedSettingServiceServer
	Db  *reform.DB
	Cfg *config.Config
}

func NewSettingServiceServer(db *reform.DB, cfg *config.Config) *SettingServiceServer {
	return &SettingServiceServer{
		Db:  db,
		Cfg: cfg,
	}
}

func (s *SettingServiceServer) SetDefaultLanguage(ctx context.Context, req *pb.SetDefaultLanguageReq) (*pb.SettRes, error) {
	dl := rep.NewSetting("DefaultLanguage", strconv.Itoa(int(req.LangId)))
	err := s.Db.Save(dl)
	if err != nil {
		return &pb.SettRes{Err: "error in saving changes in settings table"}, nil
	}
	return &pb.SettRes{Err: "success"}, nil
}

//func (s *SettingServiceServer) SetChangingParentStatus(ctx context.Context, req *pb.EmptySettReq) (*pb.SettRes, error) {
//	sett, err := s.Db.SelectAllFrom(models.SettingsTable, "where key = 'ChangingParentStatus'")
//	if err != nil {
//		return &pb.SettRes{
//			Err: "error in getting data from settings table",
//		}, err
//	}
//
//	if sett != nil {
//		if sett[0].(*models.Settings).Value == "true" {
//			sett[0].(*models.Settings).Value = "false"
//		} else {
//			sett[0].(*models.Settings).Value = "true"
//		}
//	} else {
//		nsett := rep.NewSetting("ChangingParentStatus", "true")
//		return &pb.SettRes{
//			Err: "no such column in database",
//		}, nil
//	}
//
//}
