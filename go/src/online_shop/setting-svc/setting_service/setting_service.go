package setting_service

import (
	"context"
	"fmt"
	"online_shop/repository"
	"online_shop/repository/models"
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

func (s *SettingServiceServer) FirstNewLanguage(ctx context.Context, req *pb.NewLangReq) (*pb.SettRes, error) {
	tr, err := s.Db.Begin()
	if err != nil {
		return &pb.SettRes{Err: "error in begining tranzaction"}, nil
	}

	lang := repository.NewLanguage(req.Language.Code, req.Language.Image, req.Language.Locale, req.Language.LangName, req.Language.SortOrder)

	err = tr.Insert(lang)
	if err != nil {
		tr.Rollback()
		return &pb.SettRes{Err: "error in inserting data into languages table"}, nil
	}
	dl := repository.NewSetting("DefaultLanguage", strconv.Itoa(int(lang.LangID)))
	err = tr.Save(dl)
	if err != nil {
		tr.Rollback()
		return &pb.SettRes{Err: "error in inserting data into settings table"}, nil
	}
	fmt.Println(dl.Key, dl.Value, string(lang.LangID))
	tr.Commit()
	return &pb.SettRes{Err: "success"}, nil
}

func (s *SettingServiceServer) NewLanguage(ctx context.Context, req *pb.NewLangReq) (*pb.SettRes, error) {
	tr, err := s.Db.Begin()
	if err != nil {
		return &pb.SettRes{Err: "error in begining tranzaction"}, nil
	}

	dl, err := s.Db.SelectOneFrom(models.SettingsTable, "where key = 'DefaultLanguage'")
	if err != nil {
		tr.Rollback()
		return &pb.SettRes{Err: "error in getting data from settings table"}, nil

	}

	lang := repository.NewLanguage(req.Language.Code, req.Language.Image, req.Language.Locale, req.Language.LangName, req.Language.SortOrder)
	_, err = s.Db.SelectOneFrom(models.LanguagesTable, "where lang_name = $1", req.Language.LangName)
	if err == nil {
		tr.Rollback()
		return &pb.SettRes{Err: "that language already exist"}, nil
	}

	err = tr.Insert(lang)
	if err != nil {
		tr.Rollback()
		return &pb.SettRes{Err: "error in inserting data into language table"}, nil
	}

	products, err := s.Db.SelectAllFrom(models.ProductsLocalizationView, "where lang_id = $1", dl.(*models.Settings).Value)
	if err != nil {
		tr.Rollback()
		return &pb.SettRes{Err: "error in getting data from products localization table"}, nil
	}

	if products != nil {
		for i := range products {
			products[i].(*models.ProductsLocalization).LangID = &lang.LangID
		}

		err = tr.InsertMulti(products...)
		if err != nil {
			tr.Rollback()
			return &pb.SettRes{Err: "error in inserting data into products localization table"}, nil
		}
	}

	producers, err := s.Db.SelectAllFrom(models.ProducersLocalizationView, "where lang_id = $1", dl.(*models.Settings).Value)
	if err != nil {
		tr.Rollback()
		return &pb.SettRes{Err: "error in getting data from producers localization table"}, nil
	}

	if producers != nil {
		for i := range producers {
			producers[i].(*models.ProducersLocalization).LangID = &lang.LangID
		}

		err = tr.InsertMulti(producers...)
		if err != nil {
			tr.Rollback()
			return &pb.SettRes{Err: "error in inserting data into producers localization table"}, nil
		}
	}

	groups, err := s.Db.SelectAllFrom(models.GroupsLocalizationView, "where lang_id = $1", dl.(*models.Settings).Value)
	if err != nil {
		tr.Rollback()
		return &pb.SettRes{Err: "error in getting data from groups localization table"}, nil
	}

	if groups != nil {
		for i := range groups {
			groups[i].(*models.GroupsLocalization).LangID = &lang.LangID
		}

		err = tr.InsertMulti(groups...)
		if err != nil {
			tr.Rollback()
			return &pb.SettRes{Err: "error in inserting data into groups localization table"}, nil
		}
	}

	tr.Commit()
	return &pb.SettRes{Err: "success"}, nil
}
func (s *SettingServiceServer) GetListOfLanguages(ctx context.Context, req *pb.EmptySettReq) (*pb.GetListOfLanguagesRes, error) {

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
		}
		langs = append(langs, t)

	}
	return &pb.GetListOfLanguagesRes{
		Languages: langs,
		Err:       "success",
	}, nil
}

func (s *SettingServiceServer) SetDefaultLanguage(ctx context.Context, req *pb.SetDefaultLanguageReq) (*pb.SettRes, error) {
	dl, err := s.Db.SelectOneFrom(models.SettingsTable, "where key = 'DefaultLanguage'")
	if err != nil {
		return &pb.SettRes{Err: "error in getting data from settings table"}, nil
	}

	dl.(*models.Settings).Value = strconv.Itoa(int(req.LangId))
	err = s.Db.Update(dl.(*models.Settings))
	if err != nil {
		return &pb.SettRes{Err: "error in saving changes in settings table"}, nil
	}

	return &pb.SettRes{Err: "success"}, nil
}

func (s *SettingServiceServer) ChangeLanguageStatus(ctx context.Context, req *pb.ChangeLanguageStatusReq) (*pb.SettRes, error) {
	lang, err := s.Db.SelectOneFrom(models.LanguagesTable, "where lang_id = $1", req.LangId)
	if err != nil {
		return &pb.SettRes{Err: "error in getting data from laguages table"}, nil
	}

	dl, err := s.Db.SelectOneFrom(models.SettingsTable, "where key = 'DefaultLanguage'")
	if err != nil {
		return &pb.SettRes{Err: "error in getting data from settings table"}, nil
	}

	if dl.(*models.Settings).Value == string(req.LangId) {
		return &pb.SettRes{Err: "error in saving changes in languages table, you not able to change status of the default language"}, nil
	}

	if lang.(*models.Languages).Status {
		lang.(*models.Languages).Status = false
	} else {
		lang.(*models.Languages).Status = true
	}

	err = s.Db.Update(lang.(*models.Languages))
	if err != nil {
		return &pb.SettRes{Err: "error in saving changes in languages table"}, nil
	}

	return &pb.SettRes{Err: "success"}, nil

}
