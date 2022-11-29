package prod_service

import (
	"context"
	"fmt"
	"log"
	"net"
	"online_shop/prod-svc/config"
	"online_shop/prod-svc/pb"
	"online_shop/repository"
	"online_shop/repository/models"
	"strconv"
	"testing"

	c "online_shop/setting-svc/config"
	p "online_shop/setting-svc/pb"
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

func TestProdService(t *testing.T) {

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

	srv := NewProdServiceServer(Db, &cfg)

	pb.RegisterProdServiceServer(s, srv)
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

	client := pb.NewProdServiceClient(conn)

	t.Run("RP1", func(t *testing.T) {

		cfg2, err := c.LoadConfig("../../prod_svc/config")
		if err != nil {
			fmt.Printf("Error loading loadConfig: %v", err)
			return
		}
		srv2 := setting_service.NewSettingServiceServer(Db, &cfg2)

		req := &p.NewLangReq{
			Language: &p.Language{
				Code:      "ru-ru",
				Image:     "ru.png",
				Locale:    "ru-Ru",
				LangName:  "Lang1",
				SortOrder: 0,
			},
		}

		srv2.FirstNewLanguage(ctx, req)

		var photos []string
		photos = append(photos, "photo")
		photos = append(photos, "photo2")

		m := make(map[string]*pb.Localization)
		langs, err := srv.Db.SelectAllFrom(models.LanguagesTable, "where status = true")
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
		res, _ := client.RegisterProducer(ctx, req2)
		if res.Err != "success" {
			t.Errorf("RegisterProducerTest1 failed: %v", res.Err)
		}

	})

	t.Run("GLOP1", func(t *testing.T) {

		req := &pb.EmptyProdReq{}

		res, _ := client.GetListOfProducers(ctx, req)
		if res.Err != "success" {
			t.Errorf("GetListOfProducers failed: %v", res.Err)
		}
	})

	if err := SqlDB.Close(); err != nil {
		t.Errorf("Cant close Database: %v", err)
	}

}
