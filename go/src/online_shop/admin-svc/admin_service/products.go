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

type AdminProductsServer struct {
	pb.UnimplementedAdminProductsServer
	Db  *reform.DB
	Cfg *config.Config
}

func NewAdminProductsServer(db *reform.DB, cfg *config.Config) *AdminProductsServer {
	return &AdminProductsServer{
		Db:  db,
		Cfg: cfg,
	}
}

func (s *AdminProductsServer) RegisterProduct(ctx context.Context, req *pb.RegProductReq) (*pb.AdminRes, error) {
	tr, err := s.Db.Begin()
	if err != nil {
		return &pb.AdminRes{
			Status: st.StatusInternalServerError,
			Err:    "error in beginning of the transaction: " + fmt.Sprint(err)}, nil
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
			tr.Rollback()
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
		req.Usbn, req.Mpn, req.Photos, &amount, &rating, &discount)

	err = tr.Insert(product)
	if err != nil {
		tr.Rollback()
		return &pb.AdminRes{
			Status: st.StatusInternalServerError,
			Err:    "error in inserting data into products table:" + fmt.Sprint(err)}, nil
	}

	warn := false
	for key, value := range req.Localizations {
		num, err := strconv.Atoi(key)
		if err != nil {
			tr.Rollback()
			return &pb.AdminRes{
				Status: st.StatusInvalidData,
				Err:    "invalid data in request localizations: " + fmt.Sprint(err)}, nil
		}

		if parent_id == nil {
			_, err = s.Db.SelectOneFrom(models.ProductsLocalizationView, "where title = $1", value.Title)
			if err == nil {
				warn = true
			}
		}

		loc := rep.NewLocalizationForProducts(product.ProductID, int32(num), value.Title, value.Description)
		err = tr.Insert(loc)
		if err != nil {
			tr.Rollback()
			return &pb.AdminRes{
				Status: st.StatusInternalServerError,
				Err:    "error in inserting data into products localization table: " + fmt.Sprint(err)}, nil
		}
	}

	if parent_id != nil {
		price := rep.NewProductPrices(product.ProductID, 0, req.Price)
		err := tr.Insert(price)
		if err != nil {
			tr.Rollback()
			return &pb.AdminRes{
				Status: st.StatusInternalServerError,
				Err:    "error in inserting fata in prices table: " + fmt.Sprint(err),
			}, nil
		}
	}
	tr.Commit()
	if warn {
		return &pb.AdminRes{
			Status: st.StatusOkWithWarning,
			Err:    "success, but product with this name already exist"}, nil

	}
	return &pb.AdminRes{
		Status: st.StatusOK,
		Err:    "success"}, nil
}

func (s *AdminProductsServer) GetListOfProducts(ctx context.Context, req *pb.DataTableReq) (*pb.DataTableRes, error) {
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
						WHERE p.product_id = pl.product_id and p.parent_id isnull %s`, basetail)

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

	query = fmt.Sprintf(`SELECT p.product_id, pl.title, pl.description, p.photos, p.model, p.sku, p.upc, p.jan, p.usbn, p.mpn, p.amount, p.rating, p.status, p.created_at, p.updated_at
						FROM products p, products_localization pl
						WHERE p.product_id = pl.product_id AND p.parent_id is null %s %s`, basetail, tail)

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
		var title string
		var description *string
		var photos *pq.StringArray
		var model *string
		var sku *string
		var upc *string
		var jan *string
		var usbn *string
		var mpn *string
		var amount int
		var rating float32
		var status bool
		var created_at time.Time
		var updated_at time.Time

		err := rows.Scan(&id, &title, &description, &photos, &status, &model, &sku, &upc, &jan, &usbn, &mpn, &amount, &rating, &created_at, &updated_at)
		if err != nil {
			return &pb.DataTableRes{
				Status: st.StatusInternalServerError,
				Err:    "error in processing data from products tables: " + fmt.Sprint(err)}, nil
		}

		item := map[string]any{}

		item["product_id"] = id
		item["title"] = title
		if description != nil {
			item["description"] = *description
		} else {
			item["description"] = ""
		}
		if photos != nil {
			item["photos"] = *photos
		} else {
			var ph []string
			ph = append(ph, "")
			item["photos"] = ph
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

		item["amount"] = amount
		item["rating"] = rating
		item["status"] = status
		item["created_at"] = created_at
		item["updated_at"] = updated_at

		data = append(data, item)
	}

	type Response struct {
		Draw            int
		Recordstotal    int
		Recordsfiltered int
		Data            []map[string]any
	}

	response := &Response{int(req.Draw), counttotal, countfiltered, data}

	res, err := json.Marshal(response)
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

func (s *AdminProductsServer) GetProduct(ctx context.Context, req *pb.GetProductReq) (*pb.GetProductRes, error) {

	query := fmt.Sprintf(`SELECT p.product_id, pl.title, pl.description, p.photos, p.model, p.sku, p.upc, p.jan, p.usbn, p.mpn, p.amount, p.rating, p.status, p.created_at, p.updated_at
						FROM products p, products_localization pl
						WHERE p.product_id = pl.product_id AND pl.lang_id = %d`, req.LangId)

	rows, err := s.Db.Query(query)
	if err != nil {
		return &pb.GetProductRes{
			Status: st.StatusInternalServerError,
			Err:    "error in getting data from products and products localization table: " + fmt.Sprint(err)}, nil
	}
	defer rows.Close()

	var res *pb.GetProductRes
	var product *pb.GetProductRes_Product
	for rows.Next() {
		var parent_id *int32
		var title string
		var description *string
		var photos *pq.StringArray
		var model *string
		var sku *string
		var upc *string
		var jan *string
		var usbn *string
		var mpn *string
		var amount int32
		var rating float64
		var status bool
		var created_at time.Time
		var updated_at time.Time

		err := rows.Scan(&parent_id, &title, &description, &photos, &model, &sku, &upc, &jan, &usbn, &mpn, &amount, &rating, &status, &created_at, &updated_at)
		if err != nil {
			return &pb.GetProductRes{
				Status: st.StatusInternalServerError,
				Err:    "error in processing data from products tables: " + fmt.Sprint(err)}, nil
		}

		if parent_id != nil {
			product.ParentId = *parent_id
		} else {
			product.ParentId = 0
		}
		product.Title = title
		if description != nil {
			product.Description = *description
		} else {
			product.Description = ""
		}
		if photos != nil {
			product.Photos = *photos
		} else {
			var ph []string
			ph = append(ph, "")
			product.Photos = ph
		}

		if model != nil {
			product.Model = *model
		} else {
			product.Model = ""
		}
		if sku != nil {
			product.Sku = *sku
		} else {
			product.Sku = ""
		}
		if upc != nil {
			product.Upc = *upc
		} else {
			product.Upc = ""
		}
		if jan != nil {
			product.Jan = *jan
		} else {
			product.Jan = ""
		}
		if usbn != nil {
			product.Usbn = *usbn
		} else {
			product.Usbn = ""
		}
		if mpn != nil {
			product.Mpn = *mpn
		} else {
			product.Mpn = ""
		}

		product.Amount = amount
		product.Rating = rating

		res.Product = product
	}

	if product.ParentId == 0 {
		kits_query := fmt.Sprintf(`SELECT p.product_id, pl.title 
		FROM product p, products_localization pl
		WHERE p.parent_id = %d AND pl.lang_id = %d`, req.ProductId, req.LangId)

		kits_rows, err := s.Db.Query(kits_query)
		if err != nil {
			return &pb.GetProductRes{
				Status: st.StatusInternalServerError,
				Err:    "error in getting data from products and products localization table: " + fmt.Sprint(err)}, nil
		}
		defer kits_rows.Close()

		var kits []*pb.GetProductRes_Kit
		for kits_rows.Next() {
			var product_id int32
			var title string

			err := kits_rows.Scan(&product_id, &title)
			if err != nil {
				return &pb.GetProductRes{
					Status: st.StatusInternalServerError,
					Err:    "error in processing data from products tables: " + fmt.Sprint(err)}, nil
			}

			var kit *pb.GetProductRes_Kit
			kit.ProductId = product_id
			kit.Title = title
			kits = append(kits, kit)
		}

		res.Kits = kits
	} else {

		parametrs_query := fmt.Sprintf(`SELECT p.param_id, pl.title, p.value
		FROM parametrs_products p, parametrs_localization pl
		WHERE p.product_id = %d AND pl.lang_id = %d`, req.ProductId, req.LangId)

		parametrs_rows, err := s.Db.Query(parametrs_query)
		if err != nil {
			return &pb.GetProductRes{
				Status: st.StatusInternalServerError,
				Err:    "error in getting data from products and products localization table: " + fmt.Sprint(err)}, nil
		}
		defer parametrs_rows.Close()

		var parametrs []*pb.GetProductRes_Parametr
		for parametrs_rows.Next() {

			var param_id int32
			var title string
			var value string

			err := parametrs_rows.Scan(&param_id, &title, &value)
			if err != nil {
				return &pb.GetProductRes{
					Status: st.StatusInternalServerError,
					Err:    "error in processing data from products tables: " + fmt.Sprint(err)}, nil
			}

			var parametr *pb.GetProductRes_Parametr
			parametr.ParametrId = param_id
			parametr.Title = title
			parametr.Value = value
			parametrs = append(parametrs, parametr)

		}

		res.Parametrs = parametrs
	}

	res.Status = st.StatusOK
	res.Err = "success"
	return res, nil

}

func (s *AdminProductsServer) ChangeProductStatus(ctx context.Context, req *pb.ChangeStatusReq) (*pb.ChangeStatusRes, error) {
	product, err := s.Db.SelectOneFrom(models.ProductsTable, "where product_id = $1", req.Id)
	if err != nil {
		return &pb.ChangeStatusRes{
			Status: st.StatusInternalServerError,
			Err:    "error in getting data from products table: " + fmt.Sprint(err)}, nil
	}

	if product.(*models.Products).Status {
		product.(*models.Products).Status = false
	} else {
		product.(*models.Products).Status = true
	}

	err = s.Db.Save(product.(*models.Products))
	if err != nil {
		return &pb.ChangeStatusRes{
			Status: st.StatusInternalServerError,
			Err:    "error in saving changes in products table: " + fmt.Sprint(err)}, nil
	}

	return &pb.ChangeStatusRes{
		Status: st.StatusOK,
		Err:    "success",
		Object: product.(*models.Products).Status,
	}, nil

}

func (s *AdminProductsServer) AddProductToGroup(ctx context.Context, req *pb.AddToGroupReq) (*pb.AdminRes, error) {

	gp := rep.NewGroupsProducts(req.ProductId, req.GroupId)

	err := s.Db.Insert(gp)
	if err != nil {
		return &pb.AdminRes{
			Status: st.StatusInternalServerError,
			Err:    "error in inserting data in groups products table: " + fmt.Sprint(err)}, nil
	}

	return &pb.AdminRes{
		Status: st.StatusOK,
		Err:    "success"}, nil
}

func (s *AdminProductsServer) AddProductToProducer(ctx context.Context, req *pb.AddToProducerReq) (*pb.AdminRes, error) {

	pp := rep.NewProducersProducts(req.ProductId, req.ProducerId)

	err := s.Db.Insert(pp)
	if err != nil {
		return &pb.AdminRes{
			Status: st.StatusInternalServerError,
			Err:    "error in inserting data in producers products table: " + fmt.Sprint(err)}, nil
	}

	return &pb.AdminRes{
		Status: st.StatusOK,
		Err:    "success"}, nil
}
