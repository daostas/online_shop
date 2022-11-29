package setting_service

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"net"

	"online_shop/repository"
	"online_shop/repository/models"

	//"online_shop/repository/models"
	c "online_shop/prod-svc/config"
	p "online_shop/prod-svc/pb"
	"online_shop/prod-svc/prod_service"
	"online_shop/setting-svc/config"
	"online_shop/setting-svc/pb"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}
func TestSettingService(t *testing.T) {

	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	cfg, err := config.LoadConfig("../config")
	if err != nil {
		fmt.Printf("Error loading loadConfig: %v", err)
		return
	}
	SqlDB, Db, err := repository.Conect_to_DB()
	if err != nil {
		fmt.Printf("Cant connect to Database: %v", err)
	}

	srv := NewSettingServiceServer(Db, &cfg)

	pb.RegisterSettingServiceServer(s, srv)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := pb.NewSettingServiceClient(conn)

	t.Run("FNL1", func(t *testing.T) {
		req := &pb.NewLangReq{
			Language: &pb.Language{
				Code:      "ru-ru",
				Image:     "ru.png",
				Locale:    "ru-Ru",
				LangName:  "Lang1",
				SortOrder: 0,
			},
		}
		res, _ := client.FirstNewLanguage(ctx, req)
		if res.Err != "success" {
			t.Errorf("NewLangTest1 failed: %v", res.Err)
		}
	})

	t.Run("NL1", func(t *testing.T) {
		req := &pb.NewLangReq{
			Language: &pb.Language{
				Code:      "ru-ru",
				Image:     "ru.png",
				Locale:    "ru-Ru",
				LangName:  "Lang1",
				SortOrder: 0,
			},
		}
		res, _ := client.NewLanguage(ctx, req)
		if res.Err != "that language already exist" {
			t.Errorf("NewLangTest1 failed: %v", res)
		}
	})

	t.Run("NL2", func(t *testing.T) {
		req := &pb.NewLangReq{
			Language: &pb.Language{
				Code:      "en-en",
				Image:     "en.png",
				Locale:    "en-en",
				LangName:  "Lang2",
				SortOrder: 0,
			},
		}
		res, _ := client.NewLanguage(ctx, req)
		if res.Err != "success" {
			t.Errorf("NewLangTest2 failed: %v", res)
		}
	})

	t.Run("SDL1", func(t *testing.T) {

		lang, err := srv.Db.SelectOneFrom(models.LanguagesTable, "where lang_name = 'Lang2'")
		if err != nil {
			t.Errorf("Cant get data from Database: %v", err)
		}

		req := &pb.SetDefaultLanguageReq{
			LangId: lang.(*models.Languages).LangID,
		}

		res, err := client.SetDefaultLanguage(ctx, req)
		if res.Err != "success" {
			t.Errorf("SetDefaultLanguageTest1 failed: %v", err)
		}
	})

	t.Run("GLOL1", func(t *testing.T) {

		req := &pb.EmptySettReq{}

		res, err := client.GetListOfLanguages(ctx, req)
		for i := range res.Languages {
			fmt.Println(res.Languages[i].LangName)
		}
		if res.Err != "success" {
			t.Errorf("GetListOfLanguagesTest1 failed: %v", err)
		}
	})

	t.Run("NL3", func(t *testing.T) {
		req := &pb.NewLangReq{
			Language: &pb.Language{
				Code:      "kz-kz",
				Image:     "kz.png",
				Locale:    "kz-kz",
				LangName:  "Lang3",
				SortOrder: 0,
			},
		}

		cfg2, err := c.LoadConfig("../../prod_svc/config")
		if err != nil {
			fmt.Printf("Error loading loadConfig: %v", err)
			return
		}
		srv2 := prod_service.NewProdServiceServer(Db, &cfg2)
		var photos []string
		photos = append(photos, "photo")
		photos = append(photos, "photo2")

		m := make(map[string]*p.Localization)
		langs, err := srv.Db.SelectAllFrom(models.LanguagesTable, "where status = true")
		if err != nil {
			t.Errorf("NewLangTest3 failed: %v", err)

		}

		for i := range langs {
			m[strconv.Itoa(int(langs[i].(*models.Languages).LangID))] = &p.Localization{Title: langs[i].(*models.Languages).LangName, Description: langs[i].(*models.Languages).LangName}
		}

		req2 := &p.RegProducerReq{
			Photos:        photos,
			Localizations: m,
		}

		res2, _ := srv2.RegisterProducer(ctx, req2)
		if res2.Err != "success" {
			t.Errorf("NewLangTest3 failed: %v", res2.Err)
		}

		res, _ := client.NewLanguage(ctx, req)
		if res.Err != "success" {
			t.Errorf("NewLangTest3 failed: %v", res2.Err)
		}
	})

	t.Run("CS1", func(t *testing.T) {
		lang, err := srv.Db.SelectOneFrom(models.LanguagesTable, "where lang_name = 'Lang3'")
		if err != nil {
			t.Errorf("error in getting data from languages table: %v", err)
		}
		req := &pb.ChangeLanguageStatusReq{
			LangId: lang.(*models.Languages).LangID,
		}

		res, err := client.ChangeLanguageStatus(ctx, req)

		if res.Err != "success" {
			t.Errorf("ChangeStatusTest1 failed: %v", err)
		}
	})

	t.Run("Delete", func(t *testing.T) {

		producers, err := srv.Db.SelectAllFrom(models.ProducersLocalizationView, "where title like 'Lang%'")
		if err != nil {
			t.Errorf("error in getting data from producer table: %v", err)
		}

		for i := range producers {
			_, err = srv.Db.DeleteFrom(models.ProducersLocalizationView, "where producer_id = $1", producers[i].(*models.ProducersLocalization).ProducerID)
			if err != nil {
				t.Errorf("error in deleting data from producer localization table: %v", err)
			}

			_, err = srv.Db.DeleteFrom(models.ProducersTable, "where producer_id = $1", producers[i].(*models.ProducersLocalization).ProducerID)
			if err != nil {
				t.Errorf("error in deleting data from producer localization table: %v", err)
			}
		}

		_, err = srv.Db.DeleteFrom(models.LanguagesTable, "where lang_name like 'Lang%'")
		if err != nil {
			t.Errorf("error in deleting data from languages table: %v", err)
		}

	})

	if err := SqlDB.Close(); err != nil {
		t.Errorf("Cant close Database: %v", err)
	}

}
