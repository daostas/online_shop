package admin_service

import (
	"context"
	"encoding/json"
	"fmt"
	"online_shop/admin-svc/config"
	"online_shop/admin-svc/pb"
	rep "online_shop/repository"
	"online_shop/repository/models"
	st "online_shop/status"
	"time"

	"strconv"

	"github.com/lib/pq"
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

func (s *ProductsServer) RegisterProduct(ctx context.Context, req *pb.RegProductReq) (*pb.AdminRes, error) {
	tr, err := s.Db.Begin()
	if err != nil {
		return &pb.AdminRes{
			Status: st.StatusInternalServerError,
			Err:    "error in begining of the transaction: " + fmt.Sprint(err)}, nil
	}

	var amount, discount int32
	var rating float64
	amount, discount, rating = 0, 0, 0

	var parent_id *int32
	var parent_prod reform.Struct
	if req.ParentId == 0 {
		parent_id = nil
		parent_prod = nil
	} else {
		parent_id = &req.ParentId
		parent_prod, err = s.Db.SelectOneFrom(models.ProductsTable, "where product_id = $1", parent_id)
		if err != nil {
			return &pb.AdminRes{
				Status: st.StatusInternalServerError,
				Err:    "error in getting data from products table: " + fmt.Sprint(err),
			}, nil
		}
	}

	if parent_prod != nil {
		if req.Model == "" {
			req.Model = *parent_prod.(*models.Products).Model
		}
		if req.Upc == "" {
			req.Upc = *parent_prod.(*models.Products).Upc
		}
		if req.Jan == "" {
			req.Jan = *parent_prod.(*models.Products).Jan
		}
		if req.Usbn == "" {
			req.Usbn = *parent_prod.(*models.Products).Usbn
		}
		if req.Mpn == "" {
			req.Mpn = *parent_prod.(*models.Products).Mpn
		}
	}

	product := rep.NewProduct(parent_id, req.Model, req.Sku, req.Upc, req.Jan,
		req.Usbn, req.Mpn, req.Photos, &amount, &rating, &discount, req.Status)

	err = tr.Insert(product)
	if err != nil {
		tr.Rollback()
		return &pb.AdminRes{
			Status: st.StatusInternalServerError,
			Err:    "error in inserting data into products table:" + fmt.Sprint(err)}, nil
	}

	if parent_id != nil {
		for key, value := range req.Localizations {
			num, err := strconv.Atoi(key)
			if err != nil {
				return &pb.AdminRes{Err: "invalid data: " + fmt.Sprint(err)}, nil
			}

			loc := rep.NewLocalizaionForProducts(product.ProductID, int32(num), value.Title, value.Description)
			err = tr.Insert(loc)
			if err != nil {
				tr.Rollback()
				return &pb.AdminRes{Err: "error in inserting data into products localization table: " + fmt.Sprint(err)}, nil
			}
		}
	}

	tr.Commit()
	return &pb.AdminRes{Err: "success"}, nil
}

func (s *ProductsServer) GetListOfProducts(ctx context.Context, req *pb.DataTableReq) (*pb.DataTableRes, error) {
	basetail := ""
	if req.Filter != nil {
		basetail += " and  "
		i := 0
		for key, value := range req.Filter {
			basetail += key + "::text = '" + value + "'"
			if i < len(req.Filter)-1 {
				basetail += " and "
			}
			i++
		}
	}

	if req.Search.Value != "" {
		var sc []*pb.DataTableColumns
		for _, c := range req.Columns {
			if c.Searchable {
				sc = append(sc, c)
			}
		}

		basetail += " and ("

		for i, c := range sc {
			if c.Data == "title" || c.Data == "description" {
				basetail += "pl."
			} else {
				basetail += "p."
			}

			basetail += c.Data
			if i < len(sc)-1 {
				basetail += ", "
			} else {
				basetail += ")::text"
			}
		}

		basetail += " like " + "'%" + req.Search.Value + "%' "
	}

	counttotal, err := s.Db.Count(models.GroupsTable, "")
	if err != nil {
		return &pb.DataTableRes{
			Status: st.StatusInternalServerError,
			Err:    "error in getting data from products table" + fmt.Sprint(err)}, nil
	}

	query := fmt.Sprintf(`SELECT count(*) 
						FROM products p, products_localization pl
						WHERE p.products_id = pl.products_id and p.parent_id is null %s`, basetail)

	rows, err := s.Db.Query(query)
	if err != nil {
		return &pb.DataTableRes{
			Status: st.StatusInternalServerError,
			Err:    "error in getting data from products and products localization table: " + fmt.Sprint(err)}, nil
	}
	defer rows.Close()

	var countfiltered int
	for rows.Next() {
		err := rows.Scan(&countfiltered)
		if err != nil {
			return &pb.DataTableRes{
				Status: st.StatusInternalServerError,
				Err:    "error in getting data from products tables: " + fmt.Sprint(err)}, nil
		}
	}

	if len(req.Order) != 0 {
		basetail += " order by "
		for i, o := range req.Order {
			basetail += req.Columns[o.Column].Data + " " + o.Dir
			if i < len(req.Order)-1 {
				basetail += ", "
			} else {
				basetail += " "
			}
		}
	}

	tail := fmt.Sprintf("LIMIT %d OFFSET %d", req.Length, req.Start)

	query = fmt.Sprintf(`SELECT p.product_id, p.parent_id, pl.title, pl.description, p.photos, p.model, p.sku, p.upc, p.jan, p.usbn, p.mpn, p.status, p.created_at, p.updated_at
						FROM products p, products_localization pl
						WHERE p.product_id = pl.product_id %s %s`, basetail, tail)

	rows, err = s.Db.Query(query)
	if err != nil {
		return &pb.DataTableRes{
			Status: st.StatusInternalServerError,
			Err:    "error in getting data from products and products localization table: " + fmt.Sprint(err)}, nil
	}
	defer rows.Close()

	var data []map[string]any
	for rows.Next() {
		var id int
		var parent_id *int
		var title string
		var description string
		var photos *pq.StringArray
		var model *string
		var sku *string
		var upc *string
		var jan *string
		var usbn *string
		var mpn *string
		var status bool
		var created_at time.Time
		var updated_at time.Time

		err := rows.Scan(&id, &parent_id, &title, &description, &photos, &status, &model, &sku, &upc, &jan, &usbn, &mpn, &created_at, &updated_at)
		if err != nil {
			return &pb.DataTableRes{
				Status: st.StatusInternalServerError,
				Err:    "error in processing data from products tables: " + fmt.Sprint(err)}, nil
		}

		item := map[string]any{}

		item["product_id"] = id
		item["title"] = title
		item["description"] = description
		if photos != nil {
			item["photos"] = *photos
		} else {
			item["photos"] = ""
		}

		if model != nil {
			item["model"] = *model
		} else {
			item["model"] = ""
		}

		if sku != nil {
			item["sku"] = *sku
		} else {
			item["sku"] = ""
		}

		if upc != nil {
			item["upc"] = *upc
		} else {
			item["upc"] = ""
		}

		if jan != nil {
			item["jan"] = *jan
		} else {
			item["jan"] = ""
		}

		if usbn != nil {
			item["usbn"] = *usbn
		} else {
			item["usbn"] = ""
		}

		if mpn != nil {
			item["mpn"] = *mpn
		} else {
			item["mpn"] = ""
		}

		item["status"] = status
		item["created_at"] = created_at
		item["updated_at"] = updated_at

		data = append(data, item)
	}

	type Responce struct {
		Draw            int
		Recordstotal    int
		Recordsfiltered int
		Data            []map[string]any
	}

	responce := &Responce{int(req.Draw), counttotal, countfiltered, data}

	res, err := json.Marshal(responce)
	if err != nil {
		return &pb.DataTableRes{
			Status: st.StatusInternalServerError,
			Err:    "error in processing data"}, nil
	}

	return &pb.DataTableRes{
		Status: st.StatusOK,
		Data:   res,
		Err:    "success"}, nil
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
