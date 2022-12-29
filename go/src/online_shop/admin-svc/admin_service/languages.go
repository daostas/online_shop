package admin_service

import (
	"context"
	"online_shop/admin-svc/config"
	"online_shop/admin-svc/pb"
	"online_shop/repository"
	"online_shop/repository/models"

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
		return &pb.AdminRes{Err: "error in begining tranzaction"}, nil
	}

	var check bool = false
	t, err := s.Db.SelectAllFrom(models.LanguagesTable, "")
	if err != nil {
		tr.Rollback()
		return &pb.AdminRes{Err: "error in getting data from languages table"}, nil

	}
	if t == nil {
		check = true
	}

	lang := repository.NewLanguage(req.Language.Code, req.Language.Image, req.Language.Locale, req.Language.LangName, req.Language.SortOrder)
	_, err = s.Db.SelectOneFrom(models.LanguagesTable, "where lang_name = $1", req.Language.LangName)
	if err == nil {
		tr.Rollback()
		return &pb.AdminRes{Err: "that language already exist"}, nil
	}

	err = tr.Insert(lang)
	if err != nil {
		tr.Rollback()
		return &pb.AdminRes{Err: "error in inserting data into language table"}, nil
	}

	if check {
		dl, err := s.Db.SelectOneFrom(models.SettingsTable, "where key = 'DefaultLanguage'")
		if err != nil {
			tr.Rollback()
			return &pb.AdminRes{Err: "error in getting data from settings table"}, nil

		}

		products, err := s.Db.SelectAllFrom(models.ProductsLocalizationView, "where lang_id = $1", dl.(*models.Settings).Value)
		if err != nil {
			tr.Rollback()
			return &pb.AdminRes{Err: "error in getting data from products localization table"}, nil
		}

		if products != nil {
			for i := range products {
				products[i].(*models.ProductsLocalization).LangID = &lang.LangID
			}

			err = tr.InsertMulti(products...)
			if err != nil {
				tr.Rollback()
				return &pb.AdminRes{Err: "error in inserting data into products localization table"}, nil
			}
		}

		producers, err := s.Db.SelectAllFrom(models.ProducersLocalizationView, "where lang_id = $1", dl.(*models.Settings).Value)
		if err != nil {
			tr.Rollback()
			return &pb.AdminRes{Err: "error in getting data from producers localization table"}, nil
		}

		if producers != nil {
			for i := range producers {
				producers[i].(*models.ProducersLocalization).LangID = &lang.LangID
			}

			err = tr.InsertMulti(producers...)
			if err != nil {
				tr.Rollback()
				return &pb.AdminRes{Err: "error in inserting data into producers localization table"}, nil
			}
		}

		groups, err := s.Db.SelectAllFrom(models.GroupsLocalizationView, "where lang_id = $1", dl.(*models.Settings).Value)
		if err != nil {
			tr.Rollback()
			return &pb.AdminRes{Err: "error in getting data from groups localization table"}, nil
		}

		if groups != nil {
			for i := range groups {
				groups[i].(*models.GroupsLocalization).LangID = &lang.LangID
			}

			err = tr.InsertMulti(groups...)
			if err != nil {
				tr.Rollback()
				return &pb.AdminRes{Err: "error in inserting data into groups localization table"}, nil
			}
		}
	}

	tr.Commit()
	return &pb.AdminRes{Err: "success"}, nil
}

func (s *LanguagesServer) GetListOfLanguages(ctx context.Context, req *pb.EmptyAdminReq) (*pb.GetListOfLanguagesRes, error) {

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

func (s *LanguagesServer) ChangeLanguageStatus(ctx context.Context, req *pb.ChangeLanguageStatusReq) (*pb.AdminRes, error) {
	lang, err := s.Db.SelectOneFrom(models.LanguagesTable, "where lang_id = $1", req.LangId)
	if err != nil {
		return &pb.AdminRes{Err: "error in getting data from laguages table"}, nil
	}

	dl, err := s.Db.SelectOneFrom(models.SettingsTable, "where key = 'DefaultLanguage'")
	if err != nil {
		return &pb.AdminRes{Err: "error in getting data from settings table"}, nil
	}

	if dl.(*models.Settings).Value == string(req.LangId) {
		return &pb.AdminRes{Err: "error in saving changes in languages table, you not able to change status of the default language"}, nil
	}

	if lang.(*models.Languages).Status {
		lang.(*models.Languages).Status = false
	} else {
		lang.(*models.Languages).Status = true
	}

	err = s.Db.Update(lang.(*models.Languages))
	if err != nil {
		return &pb.AdminRes{Err: "error in saving changes in languages table"}, nil
	}

	return &pb.AdminRes{Err: "success"}, nil

}
