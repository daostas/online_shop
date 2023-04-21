package admin_service

import (
	"context"
	"encoding/json"
	"fmt"
	"online_shop/admin-svc/config"
	"online_shop/admin-svc/pb"
	rep "online_shop/repository"
	"online_shop/repository/models"
	st "online_shop/status"
	"strconv"
	"time"

	//"google.golang.org/protobuf/types/known/timestamppb"

	"gopkg.in/reform.v1"
)

type AdminParametrsServer struct {
	pb.UnimplementedAdminParametrsServer
	Db  *reform.DB
	Cfg *config.Config
}

func NewAdminParametrsServer(db *reform.DB, cfg *config.Config) *AdminParametrsServer {
	return &AdminParametrsServer{
		Db:  db,
		Cfg: cfg,
	}
}

func (s *AdminParametrsServer) RegisterParametr(ctx context.Context, req *pb.RegParametrReq) (*pb.AdminRes, error) {
	tr, err := s.Db.Begin()
	if err != nil {
		return &pb.AdminRes{
			Status: st.StatusInternalServerError,
			Err:    "error in beginning of the transaction: " + fmt.Sprint(err)}, nil
	}

	parametr := rep.NewParametr()

	err = tr.Insert(parametr)
	if err != nil {
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

		_, err = s.Db.SelectOneFrom(models.ParametrsLocalizationView, "where title = $1", value.Title)
		if err == nil {
			warn = true
		}

		loc := rep.NewLocalizationForParametr(parametr.ParametrID, int32(num), value.Title)
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

func (s *AdminGroupsServer) GetListOfParametrs(_ context.Context, req *pb.DataTableReq) (*pb.DataTableRes, error) {

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
				if c.Data == "title" {
					basetail += "pl."
				} else {
					basetail += "p."
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

		counttotal, err := s.Db.Count(models.ParametrsTable, "")
		if err != nil {
			return &pb.DataTableRes{
				Status: st.StatusInternalServerError,
				Err:    "error in getting data from parametrs table: " + fmt.Sprint(err)}, nil
		}

		query := fmt.Sprintf(`SELECT count(*) 
							FROM parametrs p, parametrs_localization pl
							WHERE p.parametr_id = pl.parametr_id %s`, basetail)

		rows, err := s.Db.Query(query)
		defer rows.Close()

		if err != nil {
			return &pb.DataTableRes{
				Status: st.StatusInternalServerError,
				Err:    "error in getting data from parametrs and parametrs localization table: " + fmt.Sprint(err)}, nil
		}

		var countfiltered int
		for rows.Next() {
			err := rows.Scan(&countfiltered)
			if err != nil {
				return &pb.DataTableRes{
					Status: st.StatusInternalServerError,
					Err:    "error in processing data from parametrs tables: " + fmt.Sprint(err)}, nil
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

		query = fmt.Sprintf(`SELECT p.parametr_id, pl.title, p.status, p.created_at, p.updated_at
							FROM parametrs p, parametrs_localization pl
							WHERE p.parametr_id = pl.parametr_id %s %s`, basetail, tail)

		rows2, err := s.Db.Query(query)
		if err != nil {
			return &pb.DataTableRes{
				Status: st.StatusInternalServerError,
				Err:    "error in getting data from parametrs and parametrs localization table: " + fmt.Sprint(err)}, nil
		}
		defer rows2.Close()

		var data []map[string]any
		for rows2.Next() {
			var id int
			var title string
			var status bool
			var created_at time.Time
			var updated_at time.Time

			err := rows2.Scan(&id, &title, &status, &created_at, &updated_at)
			if err != nil {
				return &pb.DataTableRes{
					Status: st.StatusInternalServerError,
					Err:    "error in processing data from parametrs tables: " + fmt.Sprint(err)}, nil
			}

			item := map[string]any{}

			item["group_id"] = id
			item["title"] = title
			item["status"] = status
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

		query := fmt.Sprintf(`SELECT p.parametr_id, pl.title, p.status, p.created_at, p.updated_at
							FROM parametrs p, parametrs_localization pl
							WHERE p.parametr_id = pl.parametr_id AND pl.lang_id = %s`, req.Filter["lang_id"])

		rows2, err := s.Db.Query(query)
		if err != nil {
			return &pb.DataTableRes{
				Status: st.StatusInternalServerError,
				Err:    "error in getting data from parametrs and parametrs localization table: " + fmt.Sprint(err)}, nil
		}
		defer rows2.Close()

		var data []map[string]any
		for rows2.Next() {
			var id int
			var title string
			var status bool
			var created_at time.Time
			var updated_at time.Time

			err := rows2.Scan(&id, &title, &status, &created_at, &updated_at)
			if err != nil {
				return &pb.DataTableRes{
					Status: st.StatusInternalServerError,
					Err:    "error in processing data from parametrs tables: " + fmt.Sprint(err)}, nil
			}

			item := map[string]any{}

			item["group_id"] = id
			item["title"] = title
			item["status"] = status
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

func (s *AdminParametrsServer) AddParametrToGroup(ctx context.Context, req *pb.AddParametrToGroupReq) (*pb.AdminRes, error) {

	tr, err := s.Db.Begin()
	if err != nil {
		return &pb.AdminRes{
			Status: st.StatusInternalServerError,
			Err:    "error in beginning of the transaction: " + fmt.Sprint(err),
		}, nil
	}

	products, err := tr.SelectAllFrom(models.GroupsProductsView, "where group_id = $1", req.GroupId)
	if err != nil {
		tr.Rollback()
		return &pb.AdminRes{
			Status: st.StatusInternalServerError,
			Err:    "error in getting data from groups products table: " + fmt.Sprint(err)}, nil
	}

	if products != nil {
		check := false
		var nums []int32
		for _, prod := range products {
			ok := false
			parametrs, err := tr.SelectAllFrom(models.ParametrsProductsView, "where product_id = $1", prod.(*models.GroupsProducts).ProductID)
			if err != nil {
				tr.Rollback()
				return &pb.AdminRes{
					Status: st.StatusInternalServerError,
					Err:    "error in getting data from parametrs products table: " + fmt.Sprint(err)}, nil
			}

			for _, parametr := range parametrs {
				if parametr.(*models.ParametrsProducts).ParametrID == &req.GroupId {
					ok = true
				}
			}

			if !ok {
				nums = append(nums, *prod.(*models.GroupsProducts).ProductID)
				check = true
			}
		}

		if check {
			tail := ""
			for i, num := range nums {
				tail += strconv.Itoa(int(num))
				if i != len(nums)-1 {
					tail += ", "
				}
			}
			tr.Rollback()
			return &pb.AdminRes{
				Status: st.StatusForbidden,
				Err:    "not all products it this groups have that parametr. miss parametr in the products with this id: " + tail,
			}, nil
		}

	}

	_, err = tr.SelectOneFrom(models.ParametrsGroupsView, "where parametr_id = $1 and group_id = $2", req.ParametrId, req.GroupId)
	if err == nil {
		tr.Rollback()
		return &pb.AdminRes{
			Status: st.StatusAlreadyExist,
			Err:    "that parametr already exist in this group",
		}, nil
	}
	rec := rep.NewParametrsGroups(req.ParametrId, req.GroupId)

	err = tr.Insert(rec)
	if err != nil {
		tr.Rollback()
		return &pb.AdminRes{
			Status: st.StatusInternalServerError,
			Err:    "error in inserting data into parametrd groups table: " + fmt.Sprint(err),
		}, nil
	}

	tr.Commit()
	return &pb.AdminRes{
		Status: st.StatusOK,
		Err:    "success",
	}, nil
}

func (s *AdminParametrsServer) AddParametrToProduct(ctx context.Context, req *pb.AddParametrToProductReq) (*pb.AdminRes, error) {

	_, err := s.Db.SelectOneFrom(models.ParametrsProductsView, "where parametr_id = $1 and product_id = $2", req.ParametrId, req.ProductId)
	if err == nil {
		return &pb.AdminRes{
			Status: st.StatusAlreadyExist,
			Err:    "that parametr already exist in this product",
		}, nil
	}
	rec := rep.NewParametrsProducts(req.ParametrId, req.ProductId, req.Value)

	err = s.Db.Insert(rec)

	if err != nil {
		return &pb.AdminRes{
			Status: st.StatusInternalServerError,
			Err:    "error in inserting data into parametrd groups table: " + fmt.Sprint(err),
		}, nil
	}

	return &pb.AdminRes{
		Status: st.StatusOK,
		Err:    "success",
	}, nil
}
