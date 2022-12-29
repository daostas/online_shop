package admin_service

import (
	"context"
	"fmt"
	"online_shop/admin-svc/config"
	"online_shop/admin-svc/pb"
	rep "online_shop/repository"
	"online_shop/repository/models"
	"strconv"
	"strings"

	"gopkg.in/reform.v1"
)

type ProductsServer struct {
	pb.UnimplementedProductsServer
	Db  *reform.DB
	Cfg *config.Config
}

func NewProductsServer(db *reform.DB, cfg *config.Config) *ProductsServer {
	return &ProductsServer{
		Db:  db,
		Cfg: cfg,
	}
}

func (s *ProductsServer) RegisterParentProduct(ctx context.Context, req *pb.RegParentProductReq) (*pb.AdminRes, error) {
	tr, err := s.Db.Begin()
	if err != nil {
		return &pb.AdminRes{Err: "error in begining of the transaction"}, nil
	}

	var amount, discount int32
	var rating float64
	amount, discount, rating = 0, 0, 0
	product := rep.NewProduct(nil, req.Model, req.Sku, req.Upc, req.Jan, req.Usbn, req.Mpn, nil,
		&amount, &rating, &discount, req.Status)

	err = tr.Insert(product)
	if err != nil {
		tr.Rollback()
		return &pb.AdminRes{Err: "error in inserting data into products table"}, nil
	}

	var warn bool = false

	for key, value := range req.Localizations {
		num, err := strconv.Atoi(key)
		if err != nil {
			return &pb.AdminRes{Err: "invalid data"}, nil
		}

		_, err = s.Db.SelectOneFrom(models.ProductsLocalizationView, "where title = $1", value.Title)
		if err == nil {
			warn = true
		}

		loc := rep.NewLocalizaionForProducts(product.ProductID, int32(num), value.Title, value.Description)
		err = tr.Insert(loc)
		if err != nil {
			tr.Rollback()
			return &pb.AdminRes{Err: "error in inserting data into products localization table"}, nil
		}
	}

	tr.Commit()
	if warn {
		return &pb.AdminRes{Err: "success, but product with this name already exist"}, nil

	}
	return &pb.AdminRes{Err: "success"}, nil
}

func (s *ProductsServer) RegisterProduct(ctx context.Context, req *pb.RegProductReq) (*pb.AdminRes, error) {
	tr, err := s.Db.Begin()
	if err != nil {
		return &pb.AdminRes{Err: "error in begining of the transaction"}, nil
	}

	var amount, discount int32
	var rating float64
	amount, discount, rating = 0, 0, 0
	product := rep.NewProduct(&req.ParentId, req.Model, req.Sku, req.Upc, req.Jan,
		req.Usbn, req.Mpn, req.Photos, &amount, &rating, &discount, req.Status)

	err = tr.Insert(product)
	if err != nil {
		tr.Rollback()
		return &pb.AdminRes{Err: "error in inserting data into products table"}, nil
	}

	for key, value := range req.Localizations {
		num, err := strconv.Atoi(key)
		if err != nil {
			return &pb.AdminRes{Err: "invalid data"}, nil
		}

		loc := rep.NewLocalizaionForProducts(product.ProductID, int32(num), value.Title, "")
		err = tr.Insert(loc)
		if err != nil {
			tr.Rollback()
			return &pb.AdminRes{Err: "error in inserting data into products localization table"}, nil
		}
	}

	tr.Commit()
	return &pb.AdminRes{Err: "success"}, nil
}

func (s *ProductsServer) GetListOfProducts(ctx context.Context, req *pb.EmptyAdminReq) (*pb.GetListOfProductsRes, error) {
	dl, err := s.Db.SelectOneFrom(models.SettingsTable, "where key = 'DefaultLanguage'")
	if err != nil {
		return &pb.GetListOfProductsRes{
			Productslist: nil,
			Err:          "error in getting data from settings table"}, nil
	}

	num, err := strconv.Atoi(dl.(*models.Settings).Value)
	if err != nil {
		return &pb.GetListOfProductsRes{
			Productslist: nil,
			Err:          "invalid data in setting table"}, nil
	}

	query := fmt.Sprintf(`select p.product_id, pl.title
						from products p, products_localization pl
						where p.product_id = pl.product_id and pl.lang_id = %d
						and p.parent_id is null;`, num)

	rows, err := s.Db.Query(query)
	if err != nil {
		return &pb.GetListOfProductsRes{
			Productslist: nil,
			Err:          "error in getting data from products and products localization table"}, nil
	}
	defer rows.Close()

	var products []*pb.GetListOfProductsResResult

	for rows.Next() {
		var id int32
		var title string

		if err = rows.Scan(&id, &title); err != nil {
			fmt.Println(err)
		}

		query2 := fmt.Sprintf(`select photos
						from products
						where parent_id = %d;`, id)

		rows2, err := s.Db.Query(query2)
		if err != nil {
			return &pb.GetListOfProductsRes{
				Productslist: nil,
				Err:          "error in getting data from products table"}, nil
		}
		defer rows2.Close()

		var photos []string
		for rows2.Next() {
			var tphotos []uint8

			if err = rows2.Scan(&tphotos); err != nil {
				fmt.Println(err)
			}

			var str string
			for i := 1; i < len(tphotos)-1; i++ {
				str += string(tphotos[i])
			}
			ss := strings.Split(str, ",")

			photos = append(photos, ss...)
		}

		t := &pb.GetListOfProductsResResult{
			ProductId: id,
			Title:     title,
			Photos:    photos,
		}

		products = append(products, t)
	}
	if err = rows.Err(); err != nil {
		return &pb.GetListOfProductsRes{
			Productslist: nil,
			Err:          "error in processing of the data"}, nil
	}

	return &pb.GetListOfProductsRes{
		Productslist: products,
		Err:          "success"}, nil
}

func (s *ProductsServer) ChangeProductStatus(ctx context.Context, req *pb.ChangeStatusReq) (*pb.AdminRes, error) {
	product, err := s.Db.SelectOneFrom(models.ProducersTable, "where product_id = $1", req.Id)
	if err != nil {
		return &pb.AdminRes{Err: "error in getting data from product table"}, nil
	}

	if product.(*models.Products).Status {
		product.(*models.Products).Status = false
	} else {
		product.(*models.Products).Status = true
	}

	err = s.Db.Save(product.(*models.Products))
	if err != nil {
		return &pb.AdminRes{Err: "error in saving changes in products table"}, nil
	}
	return &pb.AdminRes{Err: "success"}, nil
}
func (s *ProductsServer) ChangeParentProductsStatus(ctx context.Context, req *pb.ChangeStatusReq) (*pb.AdminRes, error) {
	tr, err := s.Db.Begin()
	if err != nil {
		return &pb.AdminRes{Err: "error in begining of the transaction"}, nil
	}

	product, err := s.Db.SelectOneFrom(models.ProductsTable, "where product_id = $1", req.Id)
	if err != nil {
		tr.Rollback()
		return &pb.AdminRes{Err: "error in getting data from product table"}, nil
	}

	if product.(*models.Products).Status {
		product.(*models.Products).Status = false
	} else {
		product.(*models.Products).Status = true
	}

	err = tr.Save(product.(*models.Products))
	if err != nil {
		tr.Rollback()
		return &pb.AdminRes{Err: "error in saving changes in products table"}, nil
	}

	products, err := s.Db.SelectAllFrom(models.ProducersTable, "where parent_id = $1", req.Id)
	if err != nil {
		tr.Rollback()
		return &pb.AdminRes{Err: "error in getting data from product table"}, nil
	}

	for i := range products {
		if product.(*models.Products).Status {
			products[i].(*models.Products).Status = true
		} else {
			products[i].(*models.Products).Status = false
		}

		err = tr.Save(products[i].(*models.Products))
		if err != nil {
			tr.Rollback()
			return &pb.AdminRes{Err: "error in saving changes in products table"}, nil
		}
	}

	return &pb.AdminRes{Err: "success"}, nil
}
