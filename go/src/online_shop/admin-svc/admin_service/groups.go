package admin_service

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/reform.v1"
	"online_shop/admin-svc/config"
	"online_shop/admin-svc/pb"
	rep "online_shop/repository"
	"online_shop/repository/models"
	st "online_shop/status"
	"strconv"
	"time"
	//"google.golang.org/protobuf/types/known/timestamppb"
)

type AdminGroupsServer struct {
	pb.UnimplementedAdminGroupsServer
	Db  *reform.DB
	Cfg *config.Config
}

func NewAdminGroupsServer(db *reform.DB, cfg *config.Config) *AdminGroupsServer {
	return &AdminGroupsServer{
		Db:  db,
		Cfg: cfg,
	}
}

func (s *AdminGroupsServer) RegisterGroup(_ context.Context, req *pb.RegGroupReq) (*pb.AdminRes, error) {
	tr, err := s.Db.Begin()
	if err != nil {
		return &pb.AdminRes{
			Status: st.StatusInternalServerError,
			Err:    "error in beginning of the transaction: " + fmt.Sprint(err)}, nil
	}

	var parent_id *int32
	if req.ParentId == 0 {
		parent_id = nil
	} else {
		parent_id = &req.ParentId
	}
	fmt.Println(time.Now())
	group := rep.NewGroup(parent_id, &req.Photos, req.Status, req.SortOrder, time.Now(), time.Now())

	err = tr.Insert(group)
	if err != nil {
		fmt.Println(err)
		err2 := tr.Rollback()
		if err2 != nil {
			return &pb.AdminRes{
				Status: st.StatusInternalServerError,
				Err:    "error in inserting data into groups table and closing transaction: " + fmt.Sprint(err) + ": " + fmt.Sprint(err2)}, nil
		}
		return &pb.AdminRes{
			Status: st.StatusInternalServerError,
			Err:    "error in inserting data into groups table: " + fmt.Sprint(err)}, nil
	}

	warn := false

	for key, value := range req.Localizations {
		num, err := strconv.Atoi(key)
		if err != nil {
			return &pb.AdminRes{
				Status: st.StatusInvalidData,
				Err:    "invalid data in request localizations: " + fmt.Sprint(err)}, nil
		}

		_, err = s.Db.SelectOneFrom(models.GroupsLocalizationView, "where title = $1", value.Title)
		if err == nil {
			warn = true
		}

		loc := rep.NewLocalizationForGroups(group.GroupID, int32(num), value.Title, value.Description)
		err = tr.Insert(loc)
		if err != nil {
			err2 := tr.Rollback()
			if err2 != nil {
				return &pb.AdminRes{
					Status: st.StatusInternalServerError,
					Err:    "error in inserting data into groups table and closing transaction: " + fmt.Sprint(err) + ": " + fmt.Sprint(err2)}, nil
			}
			return &pb.AdminRes{
				Status: st.StatusInternalServerError,
				Err:    "error in inserting data into groups localization table: " + fmt.Sprint(err)}, nil
		}
	}

	err = tr.Commit()
	if err != nil {
		return &pb.AdminRes{
			Status: st.StatusInternalServerError,
			Err:    "error in closing transaction: " + fmt.Sprint(err)}, nil
	}
	if warn {
		return &pb.AdminRes{
			Status: st.StatusOkWithWarning,
			Err:    "success, but group with this name already exist"}, nil

	}
	return &pb.AdminRes{
		Status: st.StatusOK,
		Err:    "success"}, nil
}

func (s *AdminGroupsServer) GetListOfGroups(_ context.Context, req *pb.DataTableReq) (*pb.DataTableRes, error) {

	basetail := ""
	check := true
	if req.Filter != nil {

		if req.Filter["format"] == "1" {
			check = false
		}

		if check {
			basetail += " and  "
			i := 0
			for key, value := range req.Filter {
				if key == "format" {
					i++
					continue
				}
				basetail += key + "::text = '" + value + "'"
				if i < len(req.Filter)-1 {
					basetail += " and "
				}
				i++
			}
		}
	}
	if check {
		if req.Search.Value != "" {
			var sc []*pb.DataTableColumns
			for _, c := range req.Columns {
				if c.Searchable {
					sc = append(sc, c)
				}
			}

			basetail += " and ("

			for i, c := range sc {
				if c.Data == "title" || c.Data == "description" {
					basetail += "gl."
				} else {
					basetail += "g."
				}

				basetail += c.Data
				if i < len(sc)-1 {
					basetail += ", "
				} else {
					basetail += ")::text"
				}
			}

			basetail += " like " + "'%" + req.Search.Value + "%' "
		}

		counttotal, err := s.Db.Count(models.GroupsTable, "")
		if err != nil {
			return &pb.DataTableRes{
				Status: st.StatusInternalServerError,
				Err:    "error in getting data from groups table: " + fmt.Sprint(err)}, nil
		}

		query := fmt.Sprintf(`SELECT count(*) 
							FROM groups g, groups_localization gl
							WHERE g.group_id = gl.group_id %s`, basetail)

		rows, err := s.Db.Query(query)
		defer rows.Close()

		if err != nil {
			return &pb.DataTableRes{
				Status: st.StatusInternalServerError,
				Err:    "error in getting data from groups and groups localization table: " + fmt.Sprint(err)}, nil
		}

		var countfiltered int
		for rows.Next() {
			err := rows.Scan(&countfiltered)
			if err != nil {
				return &pb.DataTableRes{
					Status: st.StatusInternalServerError,
					Err:    "error in processing data from groups tables: " + fmt.Sprint(err)}, nil
			}
		}

		if len(req.Order) != 0 {
			basetail += " order by "
			for i, o := range req.Order {
				basetail += req.Columns[o.Column].Data + " " + o.Dir
				if i < len(req.Order)-1 {
					basetail += ", "
				} else {
					basetail += " "
				}
			}
		}

		tail := fmt.Sprintf("LIMIT %d OFFSET %d", req.Length, req.Start)

		query = fmt.Sprintf(`SELECT g.group_id, g.parent_id, gl.title, gl.description, g.photos, g.status, g.sort_order, g.created_at, g.updated_at
							FROM groups g, groups_localization gl
							WHERE g.group_id = gl.group_id %s %s`, basetail, tail)

		rows2, err := s.Db.Query(query)
		if err != nil {
			return &pb.DataTableRes{
				Status: st.StatusInternalServerError,
				Err:    "error in getting data from groups and groups localization table: " + fmt.Sprint(err)}, nil
		}
		defer rows2.Close()

		var data []map[string]any
		for rows2.Next() {
			var id int
			var parent_id *int
			var title string
			var description string
			var photos *string
			var status bool
			var sort_order int
			var created_at time.Time
			var updated_at time.Time

			err := rows2.Scan(&id, &parent_id, &title, &description, &photos, &status, &sort_order, &created_at, &updated_at)
			if err != nil {
				return &pb.DataTableRes{
					Status: st.StatusInternalServerError,
					Err:    "error in processing data from groups tables: " + fmt.Sprint(err)}, nil
			}

			check := true

			path := ""
			if parent_id != nil {
				for i := 0; check; i++ {
					query2 := fmt.Sprintf(`SELECT g.group_id, g.parent_id, gl.title
					FROM groups g, groups_localization gl
					WHERE g.group_id = gl.group_id and g.group_id = %d and gl.lang_id = %s`, *parent_id, req.Filter["lang_id"])

					rows2, err := s.Db.Query(query2)
					if err != nil {
						return &pb.DataTableRes{
							Err: "error in getting data from groups and groups localization table: " + fmt.Sprint(err)}, nil
					}

					for rows2.Next() {
						var id2 int
						var parent_id2 *int
						var title2 string

						err := rows2.Scan(&id2, &parent_id2, &title2)
						if err != nil {
							return &pb.DataTableRes{
								Status: st.StatusInternalServerError,
								Err:    "error in processing data from groups tables: " + fmt.Sprint(err)}, nil
						}

						if i == 0 {
							path = title2
						} else {
							path = title2 + "->" + path
						}

						if parent_id2 != nil {
							parent_id = parent_id2
						} else {
							check = false
						}
					}
				}

			}
			item := map[string]any{}

			item["group_id"] = id
			item["path"] = path
			item["title"] = title
			item["description"] = description
			if photos != nil {
				item["photos"] = *photos
			} else {
				item["photos"] = ""
			}
			item["status"] = status
			item["sort_order"] = sort_order
			item["created_at"] = created_at
			item["updated_at"] = updated_at

			data = append(data, item)
		}

		type Response struct {
			Draw            int
			Recordstotal    int
			Recordsfiltered int
			Data            []map[string]any
		}

		response := &Response{int(req.Draw), counttotal, countfiltered, data}

		res, err := json.Marshal(response)
		if err != nil {
			return &pb.DataTableRes{
				Status: st.StatusInternalServerError,
				Err:    "error in processing data: " + fmt.Sprint(err)}, nil
		}

		return &pb.DataTableRes{
			Status: st.StatusOK,
			Data:   res,
			Err:    "success"}, nil
	} else {

		query := fmt.Sprintf(`SELECT g.group_id, g.parent_id, gl.title, gl.description, g.photos, g.status, g.sort_order, g.created_at, g.updated_at
							FROM groups g, groups_localization gl
							WHERE g.group_id = gl.group_id AND gl.lang_id = %s`, req.Filter["lang_id"])

		rows2, err := s.Db.Query(query)
		if err != nil {
			return &pb.DataTableRes{
				Status: st.StatusInternalServerError,
				Err:    "error in getting data from groups and groups localization table: " + fmt.Sprint(err)}, nil
		}
		defer rows2.Close()

		var data []map[string]any
		for rows2.Next() {
			var id int
			var parent_id *int
			var title string
			var description string
			var photos *string
			var status bool
			var sort_order int
			var created_at time.Time
			var updated_at time.Time

			err := rows2.Scan(&id, &parent_id, &title, &description, &photos, &status, &sort_order, &created_at, &updated_at)
			if err != nil {
				return &pb.DataTableRes{
					Status: st.StatusInternalServerError,
					Err:    "error in processing data from groups tables: " + fmt.Sprint(err)}, nil
			}

			check := true

			path := ""
			if parent_id != nil {
				for i := 0; check; i++ {
					query2 := fmt.Sprintf(`SELECT g.group_id, g.parent_id, gl.title
					FROM groups g, groups_localization gl
					WHERE g.group_id = gl.group_id and g.group_id = %d and gl.lang_id = %s`, *parent_id, req.Filter["lang_id"])

					rows2, err := s.Db.Query(query2)
					if err != nil {
						return &pb.DataTableRes{
							Err: "error in getting data from groups and groups localization table: " + fmt.Sprint(err)}, nil
					}

					for rows2.Next() {
						var id2 int
						var parent_id2 *int
						var title2 string

						err := rows2.Scan(&id2, &parent_id2, &title2)
						if err != nil {
							return &pb.DataTableRes{
								Status: st.StatusInternalServerError,
								Err:    "error in processing data from groups tables: " + fmt.Sprint(err)}, nil
						}

						if i == 0 {
							path = title2
						} else {
							path = title2 + "->" + path
						}

						if parent_id2 != nil {
							parent_id = parent_id2
						} else {
							check = false
						}
					}
				}

			}
			item := map[string]any{}

			item["group_id"] = id
			item["path"] = path
			item["title"] = title
			item["description"] = description
			if photos != nil {
				item["photos"] = *photos
			} else {
				item["photos"] = ""
			}
			item["status"] = status
			item["sort_order"] = sort_order
			item["created_at"] = created_at
			item["updated_at"] = updated_at

			data = append(data, item)
		}

		type Response struct {
			Draw            int
			Recordstotal    int
			Recordsfiltered int
			Data            []map[string]any
		}

		response := &Response{0, 0, 0, data}

		res, err := json.Marshal(response)
		if err != nil {
			return &pb.DataTableRes{
				Status: st.StatusInternalServerError,
				Err:    "error in processing data: " + fmt.Sprint(err)}, nil
		}

		return &pb.DataTableRes{
			Status: st.StatusOK,
			Data:   res,
			Err:    "success"}, nil
	}

}

func (s *AdminGroupsServer) ChangeGroupStatus(ctx context.Context, req *pb.ChangeStatusReq) (*pb.ChangeStatusRes, error) {
	group, err := s.Db.SelectOneFrom(models.GroupsTable, "where group_id = $1", req.Id)
	if err != nil {
		return &pb.ChangeStatusRes{
			Status: st.StatusInternalServerError,
			Err:    "error in getting data from group table: " + fmt.Sprint(err)}, nil
	}

	if group.(*models.Groups).Status {
		group.(*models.Groups).Status = false
	} else {
		group.(*models.Groups).Status = true
	}

	err = s.Db.Save(group.(*models.Groups))
	if err != nil {
		return &pb.ChangeStatusRes{
			Status: st.StatusInternalServerError,
			Err:    "error in saving changes in groups table: " + fmt.Sprint(err)}, nil
	}

	return &pb.ChangeStatusRes{
		Status: st.StatusOK,
		Err:    "success",
		Object: group.(*models.Groups).Status,
	}, nil

}
