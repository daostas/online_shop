package admin_service

import (
	"context"
	"fmt"
	"log"

	"online_shop/repository"
	"online_shop/repository/models"

	//"online_shop/repository/models"
	"online_shop/admin-svc/config"
	"online_shop/admin-svc/pb"

	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func TestLanguages(t *testing.T) {

	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	cfg, err := config.LoadConfig("../config")
	if err != nil {
		fmt.Printf("Error loading loadConfig: %v", err)
		return
	}
	SqlDB, Db, err := repository.ConnectToDb()
	if err != nil {
		fmt.Printf("Cant connect to Database: %v", err)
	}

	langsrv := NewAdminLanguagesServer(Db, &cfg)

	pb.RegisterAdminLanguagesServer(s, langsrv)
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

	client := pb.NewAdminLanguagesClient(conn)

	t.Run("NewLanguage1", func(t *testing.T) {

		req := &pb.NewLangReq{
			Code:      "ru-ru",
			Image:     "ru.png",
			Locale:    "ru-Ru",
			LangName:  "Lang",
			SortOrder: 0,
		}

		res, _ := client.NewLanguage(ctx, req)

		lang, err := langsrv.Db.SelectOneFrom(models.LanguagesTable, "where lang_name = $1", req.LangName)
		if err != nil {
			t.Errorf("NewLangTest1 failed: %v", err)
		}

		lang.(*models.Languages).Code = fmt.Sprintf("lg%d-lg%d", lang.(*models.Languages).LangID, lang.(*models.Languages).LangID)
		*lang.(*models.Languages).Image = fmt.Sprintf("lang%d.png", lang.(*models.Languages).LangID)
		lang.(*models.Languages).Locale = fmt.Sprintf("lg%d-Lg%d", lang.(*models.Languages).LangID, lang.(*models.Languages).LangID)
		lang.(*models.Languages).LangName = fmt.Sprintf("Lang%d", lang.(*models.Languages).LangID)

		err = langsrv.Db.Save(lang.(*models.Languages))
		if err != nil {
			t.Errorf("NewLangTest1 failed: %v", err)
		}

		if res.Err != "success" {
			t.Errorf("NewLangTest1 failed: %v", res)
		}

	})

	// t.Run("GetListOfLanguages1", func(t *testing.T) {

	// 	req := &pb.EmptyAdminReq{}

	// 	res, err := client.GetListOfLanguages(ctx, req)
	// 	for i := range res.Languages {
	// 		fmt.Println(res.Languages[i].LangName)
	// 	}
	// 	if res.Err != "success" {
	// 		t.Errorf("GetListOfLanguagesTest1 failed: %v", err)
	// 	}
	// })

	t.Run("ChangeStatus1", func(t *testing.T) {

		lang, err := langsrv.Db.SelectAllFrom(models.LanguagesTable, "")
		if err != nil {
			t.Errorf("error in getting data from languages table: %v", err)
		}
		req := &pb.ChangeStatusReq{
			Id: lang[0].(*models.Languages).LangID,
		}

		res, err := client.ChangeLanguageStatus(ctx, req)

		if res.Err != "success" {
			t.Errorf("ChangeStatusTest1 failed: %v", err)
		}

		res, err = client.ChangeLanguageStatus(ctx, req)
		if res.Err != "success" {
			t.Errorf("returning of the previous status failed: %v", err)
		}
	})

	// t.Run("Delete", func(t *testing.T) {

	// 	producers, err := langsrv.Db.SelectAllFrom(models.ProducersLocalizationView, "where title like 'Lang%'")
	// 	if err != nil {
	// 		t.Errorf("error in getting data from producer table: %v", err)
	// 	}

	// 	for i := range producers {
	// 		_, err = langsrv.Db.DeleteFrom(models.ProducersLocalizationView, "where producer_id = $1", producers[i].(*models.ProducersLocalization).ProducerID)
	// 		if err != nil {
	// 			t.Errorf("error in deleting data from producer localization table: %v", err)
	// 		}

	// 		_, err = langsrv.Db.DeleteFrom(models.ProducersTable, "where producer_id = $1", producers[i].(*models.ProducersLocalization).ProducerID)
	// 		if err != nil {
	// 			t.Errorf("error in deleting data from producer localization table: %v", err)
	// 		}
	// 	}

	// 	_, err = langsrv.Db.DeleteFrom(models.LanguagesTable, "where lang_name like 'Lang%'")
	// 	if err != nil {
	// 		t.Errorf("error in deleting data from languages table: %v", err)
	// 	}

	// })

	if err := SqlDB.Close(); err != nil {
		t.Errorf("Cant close Database: %v", err)
	}

}
