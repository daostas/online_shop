package models

//go:generate reform

// ProducersProducts represents a row in producers_products table.
//
//reform:producers_products
type ProducersProducts struct {
	ProducerID *int32 `reform:"producer_id"`
	ProductID  *int32 `reform:"product_id"`
}
