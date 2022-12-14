// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type producersProductsViewType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *producersProductsViewType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("producers_products").
func (v *producersProductsViewType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *producersProductsViewType) Columns() []string {
	return []string{
		"producer_id",
		"product_id",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *producersProductsViewType) NewStruct() reform.Struct {
	return new(ProducersProducts)
}

// ProducersProductsView represents producers_products view or table in SQL database.
var ProducersProductsView = &producersProductsViewType{
	s: parse.StructInfo{
		Type:    "ProducersProducts",
		SQLName: "producers_products",
		Fields: []parse.FieldInfo{
			{Name: "ProducerID", Type: "*int32", Column: "producer_id"},
			{Name: "ProductID", Type: "*int32", Column: "product_id"},
		},
		PKFieldIndex: -1,
	},
	z: new(ProducersProducts).Values(),
}

// String returns a string representation of this struct or record.
func (s ProducersProducts) String() string {
	res := make([]string, 2)
	res[0] = "ProducerID: " + reform.Inspect(s.ProducerID, true)
	res[1] = "ProductID: " + reform.Inspect(s.ProductID, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *ProducersProducts) Values() []interface{} {
	return []interface{}{
		s.ProducerID,
		s.ProductID,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *ProducersProducts) Pointers() []interface{} {
	return []interface{}{
		&s.ProducerID,
		&s.ProductID,
	}
}

// View returns View object for that struct.
func (s *ProducersProducts) View() reform.View {
	return ProducersProductsView
}

// check interfaces
var (
	_ reform.View   = ProducersProductsView
	_ reform.Struct = (*ProducersProducts)(nil)
	_ fmt.Stringer  = (*ProducersProducts)(nil)
)

func init() {
	parse.AssertUpToDate(&ProducersProductsView.s, new(ProducersProducts))
}
