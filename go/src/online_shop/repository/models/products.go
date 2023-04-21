package models

import (
	"github.com/lib/pq"
	"time"
)

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
	CreatedAt        time.Time      `reform:"created_at"` // FIXME unhandled database type "timestamp without time zone"
	UpdatedAt        time.Time      `reform:"updated_at"` // FIXME unhandled database type "timestamp without time zone"
}
