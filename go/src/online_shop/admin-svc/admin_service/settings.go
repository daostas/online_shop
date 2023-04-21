package admin_service

import (
	"context"
	"fmt"
	"online_shop/admin-svc/config"
	"online_shop/admin-svc/pb"
	rep "online_shop/repository"
	st "online_shop/status"
	"strconv"

	//"google.golang.org/protobuf/types/known/timestamppb"

	"gopkg.in/reform.v1"
)

type AdminSettingServiceServer struct {
	pb.UnimplementedAdminSettingServiceServer
	Db  *reform.DB
	Cfg *config.Config
}

func NewAdminSettingServiceServer(db *reform.DB, cfg *config.Config) *AdminSettingServiceServer {
	return &AdminSettingServiceServer{
		Db:  db,
		Cfg: cfg,
	}
}

func (s *AdminSettingServiceServer) SetDefaultLanguage(ctx context.Context, req *pb.SetDefaultLanguageReq) (*pb.AdminRes, error) {
	dl := rep.NewSetting("DefaultLanguage", strconv.Itoa(int(req.LangId)))
	err := s.Db.Save(dl)
	if err != nil {
		return &pb.AdminRes{
			Status: st.StatusInternalServerError,
			Err:    "error in saving changes in settings table: " + fmt.Sprint(err)}, nil
	}
	return &pb.AdminRes{
		Status: st.StatusOK,
		Err:    "success"}, nil
}
