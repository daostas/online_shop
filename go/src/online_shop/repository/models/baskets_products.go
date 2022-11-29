package models

//go:generate reform

// BasketsProducts represents a row in baskets_products table.
//
//reform:baskets_products
type BasketsProducts struct {
	BasketID  *int32 `reform:"basket_id"`
	ProductID *int32 `reform:"product_id"`
}
