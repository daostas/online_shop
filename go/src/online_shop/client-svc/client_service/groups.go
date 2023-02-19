package client_service

import (
	"context"
	"fmt"
	"online_shop/client-svc/config"
	"online_shop/client-svc/pb"
	st "online_shop/status"

	"gopkg.in/reform.v1"
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

func (s *ClientGroupsServer) GetGroups(ctx context.Context, req *pb.GetGroupsReq) (*pb.GetGroupsRes, error) {

	if req.GroupId == 0 {
		query := fmt.Sprintf(`SELECT g.group_id, gl.title, gl.description, g.sort_order
		FROM groups g, groups_localization gl
		WHERE g.group_id = gl.group_id AND g.parent_id is null
		AND gl.lang_id = %d
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
			var sort_order int32

			err := rows.Scan(&group_id, &title, &description, &sort_order)
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

			res = append(res, &g)
		}

		return &pb.GetGroupsRes{
			Groups: res,
			Status: st.StatusOK,
			Err:    "success",
		}, nil
	} else {

		query := fmt.Sprintf(`SELECT g.group_id, gl.title, gl.description, g.sort_order
		FROM groups g, groups_localization gl
		WHERE g.group_id = gl.group_id AND g.parent_id = %d
		AND gl.lang_id = %d
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
			var sort_order int32

			err := rows.Scan(&group_id, &title, &description, &sort_order)
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

			res = append(res, &g)
		}

		return &pb.GetGroupsRes{
			Groups: res,
			Status: st.StatusOK,
			Err:    "success",
		}, nil
	}
}
