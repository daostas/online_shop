package models

import "github.com/lib/pq"

//go:generate reform

// Products represents a row in products table.
//
//reform:products
type Products struct {
	ProductID        int32          `reform:"product_id,pk"`
	ParentID         *int32         `reform:"parent_id"`
	Model            *string        `reform:"model"`
	Sku              *string        `reform:"sku"`
	Upc              *string        `reform:"upc"`
	Jan              *string        `reform:"jan"`
	Usbn             *string        `reform:"usbn"`
	Mpn              *string        `reform:"mpn"`
	Photos           pq.StringArray `reform:"photos"` // FIXME unhandled database type "ARRAY"
	Amount           *int32         `reform:"amount"`
	Rating           *float64       `reform:"rating"`
	CurreuntDiscount *int32         `reform:"curreunt_discount"`
	Status           bool           `reform:"status"`
}
