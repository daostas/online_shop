package client_service

import (
	"context"
	"fmt"
	"online_shop/client-svc/config"
	"online_shop/client-svc/pb"
	"online_shop/repository/models"
	st "online_shop/status"

	"gopkg.in/reform.v1"
	"strconv"
)

type ClientGroupsServer struct {
	pb.UnimplementedClientGroupsServer
	Db  *reform.DB
	Cfg *config.Config
}

func NewClientGroupsServer(db *reform.DB, cfg *config.Config) *ClientGroupsServer {
	return &ClientGroupsServer{
		Db:  db,
		Cfg: cfg,
	}
}

func (s *ClientGroupsServer) GetGroups(_ context.Context, req *pb.GetGroupsReq) (*pb.GetGroupsRes, error) {

	if req.LanguageId == 0 {
		lang, err := s.Db.SelectOneFrom(models.SettingsTable, "where key = 'DefaultLanguage'")
		if err != nil {
			return &pb.GetGroupsRes{
				Groups: nil,
				Status: st.StatusInternalServerError,
				Err:    "error in getting data from settings table: " + fmt.Sprint(err),
			}, nil
		}

		num, err := strconv.Atoi(lang.(*models.Settings).Value)
		if err != nil {
			return &pb.GetGroupsRes{
				Groups: nil,
				Status: st.StatusInternalServerError,
				Err:    "error in proccessing data from settings table: " + fmt.Sprint(err),
			}, nil
		}

		req.LanguageId = int32(num)
	}
	if req.GroupId == 0 {
		query := fmt.Sprintf(`SELECT g.group_id, gl.title, gl.description, g.photos
		FROM groups g, groups_localization gl
		WHERE g.group_id = gl.group_id AND g.parent_id is null
		AND gl.lang_id = %d AND g.status = true
		ORDER BY g.sort_order`, req.LanguageId)

		rows, err := s.Db.Query(query)
		if err != nil {
			return &pb.GetGroupsRes{
				Groups: nil,
				Status: st.StatusInternalServerError,
				Err:    "error in getting data from groups and groups localization table: " + fmt.Sprint(err),
			}, nil
		}

		var res []*pb.GetGroupsRes_Group

		for rows.Next() {
			var group_id int32
			var title string
			var description string
			var photos *string

			err := rows.Scan(&group_id, &title, &description, &photos)
			if err != nil {
				return &pb.GetGroupsRes{
					Groups: nil,
					Status: st.StatusInternalServerError,
					Err:    "error in getting data from groups and groups localization table: " + fmt.Sprint(err),
				}, nil
			}

			var g pb.GetGroupsRes_Group

			g.GroupId = group_id
			g.Title = title
			g.Description = description
			if photos != nil {
				g.Photos = *photos
			} else {
				g.Photos = ""
			}

			res = append(res, &g)
		}

		return &pb.GetGroupsRes{
			Groups: res,
			Status: st.StatusOK,
			Err:    "success",
		}, nil

	} else {

		query := fmt.Sprintf(`SELECT g.group_id, gl.title, gl.description, g.photos
		FROM groups g, groups_localization gl
		WHERE g.group_id = gl.group_id AND g.parent_id = %d
		AND gl.lang_id = %d AND g.status = true
		ORDER BY g.sort_order`, req.GroupId, req.LanguageId)

		rows, err := s.Db.Query(query)
		if err != nil {
			return &pb.GetGroupsRes{
				Groups: nil,
				Status: st.StatusInternalServerError,
				Err:    "error in getting data from groups and groups localization table: " + fmt.Sprint(err),
			}, nil
		}

		var res []*pb.GetGroupsRes_Group

		for rows.Next() {
			var group_id int32
			var title string
			var description string
			var photos *string

			err := rows.Scan(&group_id, &title, &description, &photos)
			if err != nil {
				return &pb.GetGroupsRes{
					Groups: nil,
					Status: st.StatusInternalServerError,
					Err:    "error in getting data from groups and groups localization table: " + fmt.Sprint(err),
				}, nil
			}

			var g pb.GetGroupsRes_Group

			g.GroupId = group_id
			g.Title = title
			g.Description = description
			if photos != nil {
				g.Photos = *photos
			} else {
				g.Photos = ""
			}

			res = append(res, &g)
		}

		return &pb.GetGroupsRes{
			Groups: res,
			Status: st.StatusOK,
			Err:    "success",
		}, nil
	}
}
