package models

//go:generate reform

// OrdersProducts represents a row in orders_products table.
//
//reform:orders_products
type OrdersProducts struct {
	OrderID   *int32   `reform:"order_id"`
	ProductID *int32   `reform:"product_id"`
	Amount    *int32   `reform:"amount"`
	Discount  *int32   `reform:"discount"`
	Price     *float64 `reform:"price"`
}
