package repository

import (
	"database/sql"
	"log"
	"online_shop/repository/config"
	"online_shop/repository/models"
	"os"
	"time"

	"github.com/lib/pq"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
)

// reform-db -db-driver=postgres -db-source=postgresql://daostas:St_031028As@185.102.75.212:5432/online_shop -debug init

// Подключение к бд
func Conect_to_DB() (*sql.DB, *reform.DB, error) {
	cfg := config.New_Config()

	SqlDB, err := sql.Open("postgres", cfg.DbAddr)
	if err != nil {
		log.Fatal(err)
	}

	logger := log.New(os.Stderr, "SQL: ", log.Flags())

	Db := reform.NewDB(SqlDB, postgresql.Dialect, reform.NewPrintfLogger(logger.Printf))

	return SqlDB, Db, nil
}

func NewUsers(name *string, number, email *string, dob *time.Time, address *string, password string) *models.Users {
	return &models.Users{
		UserName:     name,
		Number:       number,
		Email:        email,
		Dob:          dob,
		Address:      address,
		UserPassword: password,
	}
}

func NewBasket(user_id int32) *models.Baskets {
	return &models.Baskets{UserID: user_id}
}

func NewFavourite(user_id int32) *models.Favourites {
	return &models.Favourites{UserID: user_id}
}

// func NewOrder(user_id int32, payment_method, payment_details, order_address string, sum float64) *models.Orders {
// 	return &models.Orders{
// 		UserID:     		user_id,
// 		OrderCreatedAt:  	time.Now(),
// 		OrderUpdatedAt:		time.Now(),
// 		Status: 	 		"",
// 		PaymentMethod: payment_method,
// 		PaymentDetails: payment_details,
// 		OrderAddress: order_address,
// 		Sum: sum,
// 	}
// }

func NewProducer(photos pq.StringArray, status bool) *models.Producers {
	return &models.Producers{
		Photos:    photos,
		Status:    status,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func NewBasketProduct(basket_id int32, product_id int32) *models.BasketsProducts {
	return &models.BasketsProducts{
		BasketID:  &basket_id,
		ProductID: &product_id,
	}
}

func NewLanguage(code, image, locale, lang_name string, sort_order int32) *models.Languages {
	return &models.Languages{
		Code:      code,
		Image:     &image,
		Locale:    locale,
		LangName:  lang_name,
		SortOrder: sort_order,
		Status:    true,
	}
}

func NewSetting(key, value string) *models.Settings {
	return &models.Settings{
		Key:   key,
		Value: value,
	}
}

func NewAdmin(login, password, role string) *models.Admins {
	return &models.Admins{
		Login:    login,
		Password: password,
		Role:     role,
	}
}
func NewProduct(parent_id *int32, model, sku, upc, jan, usbn, mpn string, photos pq.StringArray, amount *int32, rating *float64, discount *int32, status bool) *models.Products {
	return &models.Products{
		ParentID:         parent_id,
		Model:            &model,
		Sku:              &sku,
		Upc:              &upc,
		Jan:              &jan,
		Usbn:             &usbn,
		Mpn:              &mpn,
		Photos:           photos,
		Amount:           amount,
		Rating:           rating,
		CurreuntDiscount: discount,
		Status:           status,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
}

func NewLocalizaionForProducer(producer_id int32, lang_id int32, title, description string) *models.ProducersLocalization {
	return &models.ProducersLocalization{
		ProducerID:  &producer_id,
		LangID:      &lang_id,
		Title:       title,
		Description: &description,
	}
}

func NewLocalizaionForProducts(product_id int32, lang_id int32, title, description string) *models.ProductsLocalization {
	return &models.ProductsLocalization{
		ProductID:   &product_id,
		LangID:      &lang_id,
		Title:       title,
		Description: &description,
	}
}

func NewLocalizaionForGroups(group_id int32, lang_id int32, title, description string) *models.GroupsLocalization {
	return &models.GroupsLocalization{
		GroupID:     &group_id,
		LangID:      &lang_id,
		Title:       title,
		Description: &description,
	}
}

func NewGroup(parent_id *int32, photos pq.StringArray, status bool, sort_order int32, created_at time.Time, updated_at time.Time) *models.Groups {
	return &models.Groups{
		ParentID:  parent_id,
		Photos:    photos,
		Status:    status,
		SortOrder: sort_order,
		CreatedAt: created_at,
		UpdatedAt: updated_at,
	}
}
