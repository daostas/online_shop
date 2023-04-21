package admin_service

import (
	"context"
	"encoding/json"
	"fmt"
	"online_shop/admin-svc/config"
	"online_shop/admin-svc/pb"
	"online_shop/repository"
	"online_shop/repository/models"
	st "online_shop/status"
	"strconv"
	"time"

	"gopkg.in/reform.v1"
)

type AdminLanguagesServer struct {
	pb.UnimplementedAdminLanguagesServer
	Db  *reform.DB
	Cfg *config.Config
}

func NewAdminLanguagesServer(db *reform.DB, cfg *config.Config) *AdminLanguagesServer {
	return &AdminLanguagesServer{
		Db:  db,
		Cfg: cfg,
	}
}

func (s *AdminLanguagesServer) NewLanguage(_ context.Context, req *pb.NewLangReq) (*pb.AdminRes, error) {
	tr, err := s.Db.Begin()
	if err != nil {
		return &pb.AdminRes{
			Status: st.StatusInternalServerError,
			Err:    "error in beginning transaction: " + fmt.Sprint(err)}, nil
	}

	var check bool = false
	t, err := s.Db.SelectAllFrom(models.LanguagesTable, "")
	if err != nil {
		tr.Rollback()
		return &pb.AdminRes{
			Status: st.StatusInternalServerError,
			Err:    "error in getting data from languages table: " + fmt.Sprint(err)}, nil

	}

	if t == nil {
		check = true
	}

	lang := repository.NewLanguage(req.Code, req.Image, req.Locale, req.LangName, req.SortOrder)

	_, err = s.Db.SelectOneFrom(models.LanguagesTable, "where lang_name = $1", req.LangName)
	if err == nil {
		tr.Rollback()
		return &pb.AdminRes{
			Status: st.StatusAlreadyExist,
			Err:    "that language already exist"}, nil
	}

	err = tr.Insert(lang)
	if err != nil {
		tr.Rollback()
		return &pb.AdminRes{
			Status: st.StatusInternalServerError,
			Err:    "error in inserting data into language table: " + fmt.Sprint(err)}, nil
	}

	if !check {
		dl, err := s.Db.SelectOneFrom(models.SettingsTable, "where key = 'DefaultLanguage'")
		if err != nil {
			tr.Rollback()
			return &pb.AdminRes{
				Status: st.StatusInternalServerError,
				Err:    "error in getting data from settings table: " + fmt.Sprint(err)}, nil

		}

		products, err := s.Db.SelectAllFrom(models.ProductsLocalizationView, "where lang_id = $1", dl.(*models.Settings).Value)
		if err != nil {
			tr.Rollback()
			return &pb.AdminRes{
				Status: st.StatusInternalServerError,
				Err:    "error in getting data from products localization table: " + fmt.Sprint(err)}, nil
		}

		if products != nil {
			for i := range products {
				products[i].(*models.ProductsLocalization).LangID = &lang.LangID
			}

			err = tr.InsertMulti(products...)
			if err != nil {
				tr.Rollback()
				return &pb.AdminRes{
					Status: st.StatusInternalServerError,
					Err:    "error in inserting data into products localization table: " + fmt.Sprint(err)}, nil
			}
		}

		producers, err := s.Db.SelectAllFrom(models.ProducersLocalizationView, "where lang_id = $1", dl.(*models.Settings).Value)
		if err != nil {
			tr.Rollback()
			return &pb.AdminRes{
				Status: st.StatusInternalServerError,
				Err:    "error in getting data from producers localization table: " + fmt.Sprint(err)}, nil
		}

		if producers != nil {
			for i := range producers {
				producers[i].(*models.ProducersLocalization).LangID = &lang.LangID
			}

			err = tr.InsertMulti(producers...)
			if err != nil {
				tr.Rollback()
				return &pb.AdminRes{
					Status: st.StatusInternalServerError,
					Err:    "error in inserting data into producers localization table: " + fmt.Sprint(err)}, nil
			}
		}

		groups, err := s.Db.SelectAllFrom(models.GroupsLocalizationView, "where lang_id = $1", dl.(*models.Settings).Value)
		if err != nil {
			tr.Rollback()
			return &pb.AdminRes{
				Status: st.StatusInternalServerError,
				Err:    "error in getting data from groups localization table: " + fmt.Sprint(err)}, nil
		}

		if groups != nil {
			for i := range groups {
				groups[i].(*models.GroupsLocalization).LangID = &lang.LangID
			}

			err = tr.InsertMulti(groups...)
			if err != nil {
				tr.Rollback()
				return &pb.AdminRes{
					Status: st.StatusInternalServerError,
					Err:    "error in inserting data into groups localization table: " + fmt.Sprint(err)}, nil
			}
		}
	} else {
		l, err := tr.SelectOneFrom(models.LanguagesTable, "where lang_name = $1", req.LangName)
		if err != nil {
			tr.Rollback()
			return &pb.AdminRes{
				Status: st.StatusInternalServerError,
				Err:    "error in getting data from languages table: " + fmt.Sprint(err)}, nil

		}
		dl := repository.NewSetting("DefaultLanguage", strconv.Itoa(int(l.(*models.Languages).LangID)))

		err = tr.Save(dl)
		if err != nil {
			tr.Rollback()
			return &pb.AdminRes{
				Status: st.StatusInternalServerError,
				Err:    "error in saving data into settings table: " + fmt.Sprint(err)}, nil

		}
	}
	tr.Commit()
	return &pb.AdminRes{
		Status: st.StatusOK,
		Err:    "success"}, nil
}

func (s *AdminLanguagesServer) ChangeLanguageStatus(_ context.Context, req *pb.ChangeStatusReq) (*pb.ChangeStatusRes, error) {
	lang, err := s.Db.SelectOneFrom(models.LanguagesTable, "where lang_id = $1", req.Id)
	if err != nil {
		return &pb.ChangeStatusRes{
			Status: st.StatusInternalServerError,
			Err:    "error in getting data from languages table: " + fmt.Sprint(err)}, nil
	}

	dl, err := s.Db.SelectOneFrom(models.SettingsTable, "where key = 'DefaultLanguage'")
	if err != nil {
		return &pb.ChangeStatusRes{
			Status: st.StatusInternalServerError,
			Err:    "error in getting data from settings table: " + fmt.Sprint(err)}, nil
	}

	if dl.(*models.Settings).Value == string(req.Id) {
		return &pb.ChangeStatusRes{
			Status: st.StatusForbidden,
			Err:    "error in saving changes in languages table, you not able to change status of the default language"}, nil
	}

	if lang.(*models.Languages).Status {
		lang.(*models.Languages).Status = false
	} else {
		lang.(*models.Languages).Status = true
	}

	err = s.Db.Update(lang.(*models.Languages))
	if err != nil {
		return &pb.ChangeStatusRes{
			Status: st.StatusInternalServerError,
			Err:    "error in saving changes in languages table: " + fmt.Sprint(err)}, nil
	}

	return &pb.ChangeStatusRes{
		Status: st.StatusOK,
		Err:    "success",
		Object: lang.(*models.Languages).Status}, nil

}

func (s *AdminLanguagesServer) GetListOfLanguages(_ context.Context, req *pb.DataTableReq) (*pb.DataTableRes, error) {

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
				basetail += c.Data
				if i < len(sc)-1 {
					basetail += ", "
				} else {
					basetail += ")::text"
				}
			}

			basetail += " like " + "'%" + req.Search.Value + "%' "
		}

		counttotal, err := s.Db.Count(models.LanguagesTable, "")
		if err != nil {
			return &pb.DataTableRes{
				Status: st.StatusInternalServerError,
				Err:    "error in getting data from languages table: " + fmt.Sprint(err)}, nil
		}

		var query string
		if basetail == "" {
			basetail = " lang_id > 0"
		}
		query = fmt.Sprintf(`SELECT count(*) 
						FROM languages
						WHERE %s`, basetail)

		rows, err := s.Db.Query(query)
		if err != nil {
			return &pb.DataTableRes{
				Status: st.StatusInternalServerError,
				Err:    "error in getting data from languages table: " + fmt.Sprint(err)}, nil
		}
		defer rows.Close()

		var countfiltered int
		for rows.Next() {
			err := rows.Scan(&countfiltered)
			if err != nil {
				return &pb.DataTableRes{
					Status: st.StatusInternalServerError,
					Err:    "error in processing data from languages tables: " + fmt.Sprint(err)}, nil
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

		query = fmt.Sprintf(`SELECT lang_id, code, image, locale, lang_name, sort_order, status, created_at, updated_at
							FROM languages
							WHERE %s %s`, basetail, tail)

		rows, err = s.Db.Query(query)
		if err != nil {
			return &pb.DataTableRes{
				Status: st.StatusInternalServerError,
				Err:    "error in getting data from languages table: " + fmt.Sprint(err)}, nil
		}
		defer rows.Close()

		var data []map[string]any
		for rows.Next() {
			var lang_id int
			var code string
			var image *string
			var locale string
			var lang_name string
			var sort_order int
			var status bool
			var created_at time.Time
			var updated_at time.Time

			err := rows.Scan(&lang_id, &code, &image, &locale, &lang_name, &sort_order, &status, &created_at, &updated_at)
			if err != nil {
				return &pb.DataTableRes{
					Status: st.StatusInternalServerError,
					Err:    "error in processing data from languages tables: " + fmt.Sprint(err)}, nil
			}

			item := map[string]any{}

			item["lang_id"] = lang_id
			item["code"] = code
			if image != nil {
				item["image"] = *image
			} else {
				item["image"] = ""
			}
			item["locale"] = locale
			item["lang_name"] = lang_name
			item["sort_order"] = sort_order
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

		langs, err := s.Db.SelectAllFrom(models.LanguagesTable, "")
		if err != nil {
			return &pb.DataTableRes{
				Status: st.StatusInternalServerError,
				Data:   nil,
				Err:    "error in getting data from languages table: " + fmt.Sprint(err)}, nil
		}

		var data []map[string]any

		for _, l := range langs {
			item := map[string]any{}

			item["lang_id"] = l.(*models.Languages).LangID
			item["code"] = l.(*models.Languages).Code
			if l.(*models.Languages).Image != nil {
				item["image"] = *l.(*models.Languages).Image
			} else {
				item["image"] = ""
			}
			item["locale"] = l.(*models.Languages).Locale
			item["lang_name"] = l.(*models.Languages).LangName
			item["sort_order"] = l.(*models.Languages).SortOrder
			item["status"] = l.(*models.Languages).Status
			item["created_at"] = l.(*models.Languages).CreatedAt
			item["updated_at"] = l.(*models.Languages).UpdatedAt

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
				Data:   nil,
				Err:    "error in processing data: " + fmt.Sprint(err)}, nil
		}

		return &pb.DataTableRes{
			Status: st.StatusOK,
			Data:   res,
			Err:    "success"}, nil
	}

}
