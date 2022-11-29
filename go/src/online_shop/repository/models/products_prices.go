package models

//go:generate reform

// ProductsPrices represents a row in products_prices table.
//
//reform:products_prices
type ProductsPrices struct {
	ProductID *int32  `reform:"product_id"`
	Discount  int32   `reform:"discount"`
	Price     float64 `reform:"price"`
}
