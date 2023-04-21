package admin_service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
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

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}
func TestGroups(t *testing.T) {

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

	group_client := pb.NewAdminGroupsClient(conn)

	t.Run("RegisterGroup1", func(t *testing.T) {

		photos := "photo.png"

		m := make(map[string]*pb.Localization)
		langs, err := producersrv.Db.SelectAllFrom(models.LanguagesTable, "where status = true")
		if err != nil {
			t.Errorf("NewLangTest3 failed: %v", err)

		}

		for i := range langs {
			m[strconv.Itoa(int(langs[i].(*models.Languages).LangID))] = &pb.Localization{Title: langs[i].(*models.Languages).LangName, Description: langs[i].(*models.Languages).LangName}
		}

		req2 := &pb.RegGroupReq{
			ParentId:      5,
			SortOrder:     1,
			Photos:        photos,
			Status:        true,
			Localizations: m,
		}

		res, _ := group_client.RegisterGroup(ctx, req2)
		if res.Err != "success" && res.Err != "success, but group with this name already exist" {
			t.Errorf("RegisterGroupTest1 failed: %v", res.Err)
		}

	})

	t.Run("GetListOfGroups", func(t *testing.T) {

		var columns []*pb.DataTableColumns
		column := &pb.DataTableColumns{
			Data:       "group_id",
			Name:       "",
			Searchable: true,
			Orderable:  true,
			Search: &pb.Search{
				Value: "",
				Regex: false,
			},
		}
		columns = append(columns, column)

		column = &pb.DataTableColumns{
			Data:       "title",
			Name:       "",
			Searchable: true,
			Orderable:  true,
			Search: &pb.Search{
				Value: "",
				Regex: false,
			},
		}
		columns = append(columns, column)

		column = &pb.DataTableColumns{
			Data:       "path",
			Name:       "",
			Searchable: false,
			Orderable:  false,
			Search: &pb.Search{
				Value: "",
				Regex: false,
			},
		}
		columns = append(columns, column)

		column = &pb.DataTableColumns{
			Data:       "description",
			Name:       "",
			Searchable: false,
			Orderable:  false,
			Search: &pb.Search{
				Value: "",
				Regex: false,
			},
		}
		columns = append(columns, column)

		column = &pb.DataTableColumns{
			Data:       "photos",
			Name:       "",
			Searchable: false,
			Orderable:  false,
			Search: &pb.Search{
				Value: "",
				Regex: false,
			},
		}
		columns = append(columns, column)

		column = &pb.DataTableColumns{
			Data:       "status",
			Name:       "",
			Searchable: false,
			Orderable:  false,
			Search: &pb.Search{
				Value: "",
				Regex: false,
			},
		}
		columns = append(columns, column)

		column = &pb.DataTableColumns{
			Data:       "status",
			Name:       "",
			Searchable: false,
			Orderable:  false,
			Search: &pb.Search{
				Value: "",
				Regex: false,
			},
		}
		columns = append(columns, column)

		column = &pb.DataTableColumns{
			Data:       "created_at",
			Name:       "",
			Searchable: false,
			Orderable:  false,
			Search: &pb.Search{
				Value: "",
				Regex: false,
			},
		}
		columns = append(columns, column)

		column = &pb.DataTableColumns{
			Data:       "updated_at",
			Name:       "",
			Searchable: false,
			Orderable:  false,
			Search: &pb.Search{
				Value: "",
				Regex: false,
			},
		}
		columns = append(columns, column)

		var orders []*pb.DataTableOrder
		order := &pb.DataTableOrder{
			Column: 0,
			Dir:    "asc",
		}
		orders = append(orders, order)

		search := &pb.Search{
			Value: "",
			Regex: false,
		}

		filter := make(map[string]string)
		filter["lang_id"] = "90"

		req := &pb.DataTableReq{
			Draw:    1,
			Columns: columns,
			Order:   orders,
			Start:   0,
			Length:  10,
			Search:  search,
			Filter:  filter,
		}

		res, _ := group_client.GetListOfGroups(ctx, req)
		if res.Err != "success" {
			t.Errorf("GetListOfGroups test failed: " + res.Err)
		}

		type DataTableResponse struct {
			Draw            int              `form:"draw" json:"draw"`
			Recordstotal    int              `form:"recordsTotal" json:"recordsTotal"`
			Recordsfiltered int              `form:"recordsFiltered" json:"recordsFiltered"`
			Data            []map[string]any `form:"data" json:"data"`
			Error           string           `form:"error" json:"error"`
		}

		var result DataTableResponse
		json.Unmarshal(res.Data, &result)
		fmt.Println(result)
	})

	// t.Run("DeleteAll", func(t *testing.T) {

	// 	groupsrv.Db.DeleteFrom(models.GroupsLocalizationView, "where group_id > 0")
	// 	groupsrv.Db.DeleteFrom(models.GroupsTable, "where group_id > 0")

	// })

	if err := SqlDB.Close(); err != nil {
		t.Errorf("Cant close Database: %v", err)
	}

}
