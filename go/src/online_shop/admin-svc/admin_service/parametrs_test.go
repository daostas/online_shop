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

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func TestParametrs(t *testing.T) {

	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	cfg, err := config.LoadConfig("../config")
	if err != nil {
		t.Errorf("Error loading loadConfig: %v", err)
		return
	}

	SqlDB, Db, err := repository.ConnectToDb()
	if err != nil {
		t.Errorf("Cant connect to Database: %v", err)
	}

	producersrv := NewAdminProducersServer(Db, &cfg)
	pb.RegisterAdminProducersServer(s, producersrv)
	productsrv := NewAdminProductsServer(Db, &cfg)
	pb.RegisterAdminProductsServer(s, productsrv)
	groupsrv := NewAdminGroupsServer(Db, &cfg)
	pb.RegisterAdminGroupsServer(s, groupsrv)
	langsrv := NewAdminLanguagesServer(Db, &cfg)
	pb.RegisterAdminLanguagesServer(s, langsrv)
	settsrv := NewAdminSettingServiceServer(Db, &cfg)
	pb.RegisterAdminSettingServiceServer(s, settsrv)
	paramsrv := NewAdminParametrsServer(Db, &cfg)
	pb.RegisterAdminParametrsServer(s, paramsrv)

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

	param_client := pb.NewAdminParametrsClient(conn)

	t.Run("RegPatametr", func(t *testing.T) {

		m := make(map[string]*pb.Localization)
		langs, err := paramsrv.Db.SelectAllFrom(models.LanguagesTable, "")
		if err != nil {
			t.Errorf("RegPatametr failed: %v", err)

		}

		for i := range langs {
			m[strconv.Itoa(int(langs[i].(*models.Languages).LangID))] = &pb.Localization{Title: "parametr" + strconv.Itoa(int(langs[i].(*models.Languages).LangID)) + langs[i].(*models.Languages).LangName, Description: ""}
		}

		req2 := &pb.RegParametrReq{
			Localizations: m,
		}

		res, _ := param_client.RegisterParametr(ctx, req2)
		if res.Err != "success" && res.Err != "success, but group with this name already exist" {
			t.Errorf("RegPatametr failed: %v", res.Err)
		}

	})

	t.Run("AddToGroup", func(t *testing.T) {

		req2 := &pb.AddParametrToGroupReq{
			ParametrId: 2,
			GroupId:    1,
		}

		res, _ := param_client.AddParametrToGroup(ctx, req2)
		if res.Err != "success" && res.Err != "success, but group with this name already exist" {
			t.Errorf("RegPatametr failed: %v", res.Err)
		}

	})

	t.Run("AddToProduct", func(t *testing.T) {

		req2 := &pb.AddParametrToProductReq{
			ParametrId: 1,
			ProductId:  5,
		}

		res, _ := param_client.AddParametrToProduct(ctx, req2)
		if res.Err != "success" && res.Err != "success, but group with this name already exist" {
			t.Errorf("RegPatametr failed: %v", res.Err)
		}

	})

	if err := SqlDB.Close(); err != nil {
		t.Errorf("Cant close Database: %v", err)
	}

}
