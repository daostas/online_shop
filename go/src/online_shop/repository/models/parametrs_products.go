package models

//go:generate reform

// ParametrsProducts represents a row in parametrs_products table.
//
//reform:parametrs_products
type ParametrsProducts struct {
	ProductID  *int32 `reform:"product_id"`
	ParametrID *int32 `reform:"parametr_id"`
	Value      string `reform:"value"`
}
