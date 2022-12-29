package admin_service

import (
	"context"
	"log"
	"net"
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

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestAdminService(t *testing.T) {

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
	groupsrv := NewGroupsServer(Db, &cfg)
	pb.RegisterGroupsServer(s, groupsrv)
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

	client1 := pb.NewProducersClient(conn)
	client2 := pb.NewLanguagesClient(conn)

	t.Run("RP1", func(t *testing.T) {

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

		client2.NewLanguage(ctx, req)
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

		req2 := &pb.RegProducerReq{
			Photos:        photos,
			Localizations: m,
		}
		res, _ := client1.RegisterProducer(ctx, req2)
		if res.Err != "success" {
			t.Errorf("RegisterProducerTest1 failed: %v", res.Err)
		}

	})

	t.Run("GLOP1", func(t *testing.T) {

		req := &pb.EmptyAdminReq{}

		res, _ := client1.GetListOfProducers(ctx, req)
		if res.Err != "success" {
			t.Errorf("GetListOfProducers failed: %v", res.Err)
		}
	})

	if err := SqlDB.Close(); err != nil {
		t.Errorf("Cant close Database: %v", err)
	}

}
