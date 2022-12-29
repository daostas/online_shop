package setting_service

import (
	"context"
	"fmt"
	"log"
	"net"

	"online_shop/repository"
	"online_shop/repository/models"

	//"online_shop/repository/models"
	admin_service "online_shop/admin-svc/admin_service"
	admincfg "online_shop/admin-svc/config"
	adminpb "online_shop/admin-svc/pb"
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

	t.Run("SDL1", func(t *testing.T) {
		admincfg, err := admincfg.LoadConfig("../../setting-svc/config")
		if err != nil {
			t.Errorf("Error loading loadConfig: %v", err)
			return
		}

		langsrv := admin_service.NewLanguagesServer(Db, &admincfg)
		langreq := &adminpb.NewLangReq{
			Language: &adminpb.Language{
				Code:      "ru-ru",
				Image:     "ru.png",
				Locale:    "ru-Ru",
				LangName:  "Lang1",
				SortOrder: 0,
			},
		}

		langsrv.NewLanguage(ctx, langreq)

		lang, err := srv.Db.SelectOneFrom(models.LanguagesTable, "where lang_name = 'Lang1'")
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

	if err := SqlDB.Close(); err != nil {
		t.Errorf("Cant close Database: %v", err)
	}

}
