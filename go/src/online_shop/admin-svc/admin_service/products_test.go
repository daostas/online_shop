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

func TestProducts(t *testing.T) {

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

	productclient := pb.NewAdminProductsClient(conn)

	t.Run("RegParentProduct", func(t *testing.T) {

		m := make(map[string]*pb.Localization)
		langs, err := paramsrv.Db.SelectAllFrom(models.LanguagesTable, "")
		if err != nil {
			t.Errorf("RegPatametr failed: %v", err)

		}

		for i := range langs {
			m[strconv.Itoa(int(langs[i].(*models.Languages).LangID))] = &pb.Localization{Title: "Phone title " + strconv.Itoa(int(langs[i].(*models.Languages).LangID)) + langs[i].(*models.Languages).LangName, Description: ""}
		}

		req2 := &pb.
			RegProductReq{
			ParentId:        0,
			Model:           "Model1",
			Sku:             "Sku1",
			Upc:             "Upc1",
			Jan:             "Jan1",
			Usbn:            "Usbn1",
			Mpn:             "Mpn1",
			Photos:          nil,
			Amount:          0,
			Rating:          0,
			CurrentDiscount: 0,
			Price:           0,
			Localizations:   m,
		}

		res, _ := productclient.RegisterProduct(ctx, req2)
		if res.Err != "success" && res.Err != "success, but group with this name already exist" {
			t.Errorf("RegPatametr failed: %v", res.Err)
		}

	})

	t.Run("RegChildProduct", func(t *testing.T) {

		m := make(map[string]*pb.Localization)
		langs, err := paramsrv.Db.SelectAllFrom(models.LanguagesTable, "")
		if err != nil {
			t.Errorf("RegPatametr failed: %v", err)

		}

		for i := range langs {
			m[strconv.Itoa(int(langs[i].(*models.Languages).LangID))] = &pb.Localization{Title: "Phone" + strconv.Itoa(int(langs[i].(*models.Languages).LangID)) + langs[i].(*models.Languages).LangName, Description: ""}
		}

		req2 := &pb.
			RegProductReq{
			ParentId:        5,
			Model:           "",
			Sku:             "Sku2",
			Upc:             "",
			Jan:             "",
			Usbn:            "",
			Mpn:             "",
			Photos:          nil,
			Amount:          0,
			Rating:          0,
			CurrentDiscount: 0,
			Price:           100,
			Localizations:   m,
		}

		res, _ := productclient.RegisterProduct(ctx, req2)
		if res.Err != "success" && res.Err != "success, but group with this name already exist" {
			t.Errorf("RegPatametr failed: %v", res.Err)
		}

	})

	if err := SqlDB.Close(); err != nil {
		t.Errorf("Cant close Database: %v", err)
	}

}
