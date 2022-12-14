// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type basketsProductsViewType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *basketsProductsViewType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("baskets_products").
func (v *basketsProductsViewType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *basketsProductsViewType) Columns() []string {
	return []string{
		"basket_id",
		"product_id",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *basketsProductsViewType) NewStruct() reform.Struct {
	return new(BasketsProducts)
}

// BasketsProductsView represents baskets_products view or table in SQL database.
var BasketsProductsView = &basketsProductsViewType{
	s: parse.StructInfo{
		Type:    "BasketsProducts",
		SQLName: "baskets_products",
		Fields: []parse.FieldInfo{
			{Name: "BasketID", Type: "*int32", Column: "basket_id"},
			{Name: "ProductID", Type: "*int32", Column: "product_id"},
		},
		PKFieldIndex: -1,
	},
	z: new(BasketsProducts).Values(),
}

// String returns a string representation of this struct or record.
func (s BasketsProducts) String() string {
	res := make([]string, 2)
	res[0] = "BasketID: " + reform.Inspect(s.BasketID, true)
	res[1] = "ProductID: " + reform.Inspect(s.ProductID, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *BasketsProducts) Values() []interface{} {
	return []interface{}{
		s.BasketID,
		s.ProductID,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *BasketsProducts) Pointers() []interface{} {
	return []interface{}{
		&s.BasketID,
		&s.ProductID,
	}
}

// View returns View object for that struct.
func (s *BasketsProducts) View() reform.View {
	return BasketsProductsView
}

// check interfaces
var (
	_ reform.View   = BasketsProductsView
	_ reform.Struct = (*BasketsProducts)(nil)
	_ fmt.Stringer  = (*BasketsProducts)(nil)
)

func init() {
	parse.AssertUpToDate(&BasketsProductsView.s, new(BasketsProducts))
}
