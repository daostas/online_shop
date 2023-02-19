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
	"time"

	"github.com/lib/pq"
	"gopkg.in/reform.v1"
)

type LanguagesServer struct {
	pb.UnimplementedLanguagesServer
	Db  *reform.DB
	Cfg *config.Config
}

func NewLanguagesServer(db *reform.DB, cfg *config.Config) *LanguagesServer {
	return &LanguagesServer{
		Db:  db,
		Cfg: cfg,
	}
}

func (s *LanguagesServer) NewLanguage(ctx context.Context, req *pb.NewLangReq) (*pb.AdminRes, error) {
	tr, err := s.Db.Begin()
	if err != nil {
		return &pb.AdminRes{
			Status: st.StatusInternalServerError,
			Err:    "error in begining tranzaction: " + fmt.Sprint(err)}, nil
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

	lang := repository.NewLanguage(req.Language.Code, req.Language.Image, req.Language.Locale, req.Language.LangName, req.Language.SortOrder)

	_, err = s.Db.SelectOneFrom(models.LanguagesTable, "where lang_name = $1", req.Language.LangName)
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
	}

	tr.Commit()
	return &pb.AdminRes{Err: "success"}, nil
}

func (s *LanguagesServer) ChangeLanguageStatus(ctx context.Context, req *pb.ChangeLanguageStatusReq) (*pb.ChangeStatusRes, error) {
	lang, err := s.Db.SelectOneFrom(models.LanguagesTable, "where lang_id = $1", req.LangId)
	if err != nil {
		return &pb.ChangeStatusRes{Err: "error in getting data from laguages table"}, nil
	}

	dl, err := s.Db.SelectOneFrom(models.SettingsTable, "where key = 'DefaultLanguage'")
	if err != nil {
		return &pb.ChangeStatusRes{Err: "error in getting data from settings table"}, nil
	}

	if dl.(*models.Settings).Value == string(req.LangId) {
		return &pb.ChangeStatusRes{Err: "error in saving changes in languages table, you not able to change status of the default language"}, nil
	}

	if lang.(*models.Languages).Status {
		lang.(*models.Languages).Status = false
	} else {
		lang.(*models.Languages).Status = true
	}

	err = s.Db.Update(lang.(*models.Languages))
	if err != nil {
		return &pb.ChangeStatusRes{Err: "error in saving changes in languages table"}, nil
	}

	return &pb.ChangeStatusRes{Err: "success"}, nil

}

func (s *LanguagesServer) GetListOfLanguages1(ctx context.Context, req *pb.EmptyAdminReq) (*pb.GetListOfLanguagesRes, error) {

	tlangs, err := s.Db.SelectAllFrom(models.LanguagesTable, "")
	if err != nil {
		return &pb.GetListOfLanguagesRes{
			Languages: nil,
			Err:       "error in getting data from languages table",
		}, nil
	}
	var langs []*pb.Language
	for _, lang := range tlangs {
		t := &pb.Language{
			LangId:    lang.(*models.Languages).LangID,
			Code:      lang.(*models.Languages).Code,
			Image:     *lang.(*models.Languages).Image,
			Locale:    lang.(*models.Languages).Locale,
			LangName:  lang.(*models.Languages).LangName,
			SortOrder: lang.(*models.Languages).SortOrder,
			Status:    lang.(*models.Languages).Status,
		}
		langs = append(langs, t)

	}
	return &pb.GetListOfLanguagesRes{
		Languages: langs,
		Err:       "success",
	}, nil
}

func (s *LanguagesServer) GetListOfLanguages(ctx context.Context, req *pb.DataTableReq) (*pb.DataTableRes, error) {

	basetail := ""
	if req.Filter != nil {
		basetail += " and  "
		i := 0
		for key, value := range req.Filter {
			basetail += key + "::text = '" + value + "'"
			if i < len(req.Filter)-1 {
				basetail += " and "
			}
			i++
		}
	}

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

	query := fmt.Sprintf(`SELECT count(*) 
						FROM languages l
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

	query = fmt.Sprintf(`SELECT g.group_id, g.parent_id, gl.title, gl.description, g.photos, g.status, g.sort_order, g.created_at, g.updated_at
						FROM groups g, groups_localization gl
						WHERE g.group_id = gl.group_id %s %s`, basetail, tail)

	rows, err = s.Db.Query(query)
	if err != nil {
		return &pb.DataTableRes{
			Status: st.StatusInternalServerError,
			Err:    "error in getting data from groups and groups localization table: " + fmt.Sprint(err)}, nil
	}
	defer rows.Close()

	var data []map[string]any
	for rows.Next() {
		var id int
		var parent_id *int
		var title string
		var description string
		var photos *pq.StringArray
		var status bool
		var sort_order int
		var created_at time.Time
		var updated_at time.Time

		err := rows.Scan(&id, &parent_id, &title, &description, &photos, &status, &sort_order, &created_at, &updated_at)
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
				WHERE g.group_id = gl.group_id and g.group_id = %d`, *parent_id)

				rows2, err := s.Db.Query(query2)
				if err != nil {
					return &pb.DataTableRes{
						Err: "error in getting data from groups and groups localization table: " + fmt.Sprint(err)}, nil
				}
				defer rows.Close()

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

	type Responce struct {
		Draw            int
		Recordstotal    int
		Recordsfiltered int
		Data            []map[string]any
	}

	responce := &Responce{int(req.Draw), counttotal, countfiltered, data}

	res, err := json.Marshal(responce)
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
