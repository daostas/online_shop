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

	"github.com/lib/pq"
	"gopkg.in/reform.v1"
)

type AdminProducersServer struct {
	pb.UnimplementedAdminProducersServer
	Db  *reform.DB
	Cfg *config.Config
}

func NewAdminProducersServer(db *reform.DB, cfg *config.Config) *AdminProducersServer {
	return &AdminProducersServer{
		Db:  db,
		Cfg: cfg,
	}
}

func (s *AdminProducersServer) RegisterProducer(_ context.Context, req *pb.RegProducerReq) (*pb.AdminRes, error) {
	tr, err := s.Db.Begin()
	if err != nil {
		return &pb.AdminRes{
			Status: st.StatusInternalServerError,
			Err:    "error in beginning of the transaction: " + fmt.Sprint(err)}, nil
	}

	producer := rep.NewProducer(req.Photos, req.Status)

	err = tr.Insert(producer)
	if err != nil {
		tr.Rollback()
		return &pb.AdminRes{
			Status: st.StatusInternalServerError,
			Err:    "error in inserting data into producers table: " + fmt.Sprint(err)}, nil
	}

	var warn bool = false

	for key, value := range req.Localizations {
		num, err := strconv.Atoi(key)
		if err != nil {
			return &pb.AdminRes{
				Status: st.StatusInvalidData,
				Err:    "invalid data: " + fmt.Sprint(err)}, nil
		}

		_, err = s.Db.SelectOneFrom(models.ProducersLocalizationView, "where title = $1", value.Title)
		if err == nil {
			warn = true
		}

		loc := rep.NewLocalizationForProducer(producer.ProducerID, int32(num), value.Title, value.Description)
		err = tr.Insert(loc)
		if err != nil {
			tr.Rollback()
			return &pb.AdminRes{
				Status: st.StatusInternalServerError,
				Err:    "error in inserting data into producers localization table: " + fmt.Sprint(err)}, nil
		}
	}

	tr.Commit()
	if warn {
		return &pb.AdminRes{
			Status: st.StatusOkWithWarning,
			Err:    "success, but producer with this name already exist"}, nil

	}
	return &pb.AdminRes{
		Status: st.StatusOK,
		Err:    "success"}, nil
}

func (s *AdminProducersServer) GetListOfProducers(_ context.Context, req *pb.DataTableReq) (*pb.DataTableRes, error) {
	basetail := ""
	check := true
	if req.Filter != nil {
		if req.Filter["format"] == "1" {
			check = false
		}

		if check {
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

		counttotal, err := s.Db.Count(models.ProducersTable, "")
		if err != nil {
			return &pb.DataTableRes{
				Status: st.StatusInternalServerError,
				Err:    "error in getting data from producers table: " + fmt.Sprint(err)}, nil
		}

		query := fmt.Sprintf(`SELECT count(*) 
							FROM producers p, producers_localization pl 
							WHERE p.producer_id = pl.producer_id AND %s`, basetail)

		rows, err := s.Db.Query(query)
		if err != nil {
			return &pb.DataTableRes{
				Status: st.StatusInternalServerError,
				Err:    "error in getting data from producers table: " + fmt.Sprint(err)}, nil
		}
		defer rows.Close()

		var countfiltered int
		for rows.Next() {
			err := rows.Scan(&countfiltered)
			if err != nil {
				return &pb.DataTableRes{
					Status: st.StatusInternalServerError,
					Err:    "error in processing data from producers tables: " + fmt.Sprint(err)}, nil
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

		query = fmt.Sprintf(`SELECT p.producer_id, pl.title, pl.description, p.photos, p.status, p.created_at, p.updated_at
						FROM producers p, producers_localization pl
						WHERE p.producer_id = pl.producer_id and %s %s`, basetail, tail)

		rows, err = s.Db.Query(query)
		if err != nil {
			return &pb.DataTableRes{
				Status: st.StatusInternalServerError,
				Err:    "error in getting data from producers table: " + fmt.Sprint(err)}, nil
		}
		defer rows.Close()

		var data []map[string]any
		for rows.Next() {
			var id int
			var title string
			var description string
			var photos *pq.StringArray
			var status bool
			var created_at time.Time
			var updated_at time.Time

			err := rows.Scan(&id, &title, &description, &photos, &status, &created_at, &updated_at)
			if err != nil {
				return &pb.DataTableRes{
					Status: st.StatusInternalServerError,
					Err:    "error in processing data from producers tables: " + fmt.Sprint(err)}, nil
			}

			item := map[string]any{}

			item["producer_id"] = id
			item["title"] = title
			item["description"] = description
			if photos != nil {
				item["photos"] = *photos
			} else {
				var ph []string
				ph = append(ph, "")
				item["photos"] = ph
			}
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
		query := fmt.Sprintf(`SELECT p.producer_id, pl.title, pl.description, p.photos, p.status, p.created_at, p.updated_at
						FROM producers p, producers_localization pl
						WHERE p.producer_id = pl.producer_id and pl.lang_id = %s`, req.Filter["lang_id"])

		rows, err := s.Db.Query(query)
		if err != nil {
			return &pb.DataTableRes{
				Status: st.StatusInternalServerError,
				Err:    "error in getting data from producers table: " + fmt.Sprint(err)}, nil
		}
		defer rows.Close()

		var data []map[string]any
		for rows.Next() {
			var id int
			var title string
			var description string
			var photos *pq.StringArray
			var status bool
			var created_at time.Time
			var updated_at time.Time

			err := rows.Scan(&id, &title, &description, &photos, &status, &created_at, &updated_at)
			if err != nil {
				return &pb.DataTableRes{
					Status: st.StatusInternalServerError,
					Err:    "error in processing data from producers tables: " + fmt.Sprint(err)}, nil
			}

			item := map[string]any{}

			item["producer_id"] = id
			item["title"] = title
			item["description"] = description
			if photos != nil {
				item["photos"] = *photos
			} else {
				var ph []string
				ph = append(ph, "")
				item["photos"] = ph
			}
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

func (s *AdminProducersServer) ChangeProducerStatus(ctx context.Context, req *pb.ChangeStatusReq) (*pb.ChangeStatusRes, error) {
	producer, err := s.Db.SelectOneFrom(models.ProducersTable, "where producer_id = $1", req.Id)
	if err != nil {
		return &pb.ChangeStatusRes{
			Status: st.StatusInternalServerError,
			Err:    "error in getting data from producers table: " + fmt.Sprint(err)}, nil
	}

	if producer.(*models.Producers).Status {
		producer.(*models.Producers).Status = false
	} else {
		producer.(*models.Producers).Status = true
	}

	err = s.Db.Save(producer.(*models.Producers))
	if err != nil {
		return &pb.ChangeStatusRes{
			Status: st.StatusInternalServerError,
			Err:    "error in saving changes in producers table: " + fmt.Sprint(err)}, nil
	}

	return &pb.ChangeStatusRes{
		Status: st.StatusOK,
		Err:    "success",
		Object: producer.(*models.Producers).Status,
	}, nil

}
