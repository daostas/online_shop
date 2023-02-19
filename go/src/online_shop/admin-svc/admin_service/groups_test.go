package admin_service

import (
	"context"
	"log"
	"online_shop/admin-svc/config"
	"online_shop/admin-svc/pb"
	"online_shop/repository"
	"online_shop/repository/models"
	"strconv"
	"testing"

	settcfg "online_shop/setting-svc/config"
	settpb "online_shop/setting-svc/pb"
	"online_shop/setting-svc/setting_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func TestGroups(t *testing.T) {

	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	cfg, err := config.LoadConfig("../config")
	if err != nil {
		t.Errorf("Error loading loadConfig: %v", err)
		return
	}

	SqlDB, Db, err := repository.Conect_to_DB()
	if err != nil {
		t.Errorf("Cant connect to Database: %v", err)
	}

	producersrv := NewProducersServer(Db, &cfg)
	pb.RegisterProducersServer(s, producersrv)
	productsrv := NewProductsServer(Db, &cfg)
	pb.RegisterProductsServer(s, productsrv)
	groupsrv := NewAdminGroupsServer(Db, &cfg)
	pb.RegisterAdminGroupsServer(s, groupsrv)
	langsrv := NewLanguagesServer(Db, &cfg)
	pb.RegisterLanguagesServer(s, langsrv)

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

	language_client := pb.NewLanguagesClient(conn)
	group_client := pb.NewAdminGroupsClient(conn)

	t.Run("RegisterGroup1", func(t *testing.T) {

		settcfg, err := settcfg.LoadConfig("../../setting-svc/config")
		if err != nil {
			t.Errorf("Error loading loadConfig: %v", err)
			return
		}

		settsrv := setting_service.NewSettingServiceServer(Db, &settcfg)

		req := &pb.NewLangReq{
			Language: &pb.Language{
				Code:      "ru-ru",
				Image:     "ru.png",
				Locale:    "ru-Ru",
				LangName:  "Lang1",
				SortOrder: 0,
			},
		}

		language_client.NewLanguage(ctx, req)
		nl, _ := settsrv.Db.SelectOneFrom(models.LanguagesTable, "where lang_name = $1", req.Language.LangName)
		settreq := &settpb.SetDefaultLanguageReq{
			LangId: nl.(*models.Languages).LangID,
		}
		settsrv.SetDefaultLanguage(ctx, settreq)

		var photos []string
		photos = append(photos, "photo")
		photos = append(photos, "photo2")

		m := make(map[string]*pb.Localization)
		langs, err := producersrv.Db.SelectAllFrom(models.LanguagesTable, "where status = true")
		if err != nil {
			t.Errorf("NewLangTest3 failed: %v", err)

		}

		for i := range langs {
			m[strconv.Itoa(int(langs[i].(*models.Languages).LangID))] = &pb.Localization{Title: langs[i].(*models.Languages).LangName, Description: langs[i].(*models.Languages).LangName}
		}

		req2 := &pb.RegGroupReq{
			SortOrder:     1,
			Photos:        photos,
			Status:        true,
			Localizations: m,
		}

		res, _ := group_client.RegisterGroup(ctx, req2)
		if res.Err != "success" {
			t.Errorf("RegisterGroupTest1 failed: %v", res.Err)
		}

	})

	t.Run("GetListOfGroups", func(t *testing.T) {

		type Search struct {
			Value string
			Regex bool
		}
		type Column struct {
			Data       string
			Name       string
			Searchable bool
			Orderable  bool
			Search     Search
		}

	})

	t.Run("DeleteAll", func(t *testing.T) {

		groupsrv.Db.DeleteFrom(models.GroupsLocalizationView, "where group_id > 0")
		groupsrv.Db.DeleteFrom(models.GroupsTable, "where group_id > 0")

	})

	if err := SqlDB.Close(); err != nil {
		t.Errorf("Cant close Database: %v", err)
	}

}
