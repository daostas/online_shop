// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type reviewsTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *reviewsTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("reviews").
func (v *reviewsTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *reviewsTableType) Columns() []string {
	return []string{
		"review_id",
		"user_id",
		"product_id",
		"content",
		"photos",
		"rating",
		"created_at",
		"updated_at",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *reviewsTableType) NewStruct() reform.Struct {
	return new(Reviews)
}

// NewRecord makes a new record for that table.
func (v *reviewsTableType) NewRecord() reform.Record {
	return new(Reviews)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *reviewsTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// ReviewsTable represents reviews view or table in SQL database.
var ReviewsTable = &reviewsTableType{
	s: parse.StructInfo{
		Type:    "Reviews",
		SQLName: "reviews",
		Fields: []parse.FieldInfo{
			{Name: "ReviewID", Type: "int32", Column: "review_id"},
			{Name: "UserID", Type: "*int32", Column: "user_id"},
			{Name: "ProductID", Type: "*int32", Column: "product_id"},
			{Name: "Content", Type: "*string", Column: "content"},
			{Name: "Photos", Type: "[]uint8", Column: "photos"},
			{Name: "Rating", Type: "float64", Column: "rating"},
			{Name: "CreatedAt", Type: "[]uint8", Column: "created_at"},
			{Name: "UpdatedAt", Type: "[]uint8", Column: "updated_at"},
		},
		PKFieldIndex: 0,
	},
	z: new(Reviews).Values(),
}

// String returns a string representation of this struct or record.
func (s Reviews) String() string {
	res := make([]string, 8)
	res[0] = "ReviewID: " + reform.Inspect(s.ReviewID, true)
	res[1] = "UserID: " + reform.Inspect(s.UserID, true)
	res[2] = "ProductID: " + reform.Inspect(s.ProductID, true)
	res[3] = "Content: " + reform.Inspect(s.Content, true)
	res[4] = "Photos: " + reform.Inspect(s.Photos, true)
	res[5] = "Rating: " + reform.Inspect(s.Rating, true)
	res[6] = "CreatedAt: " + reform.Inspect(s.CreatedAt, true)
	res[7] = "UpdatedAt: " + reform.Inspect(s.UpdatedAt, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Reviews) Values() []interface{} {
	return []interface{}{
		s.ReviewID,
		s.UserID,
		s.ProductID,
		s.Content,
		s.Photos,
		s.Rating,
		s.CreatedAt,
		s.UpdatedAt,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Reviews) Pointers() []interface{} {
	return []interface{}{
		&s.ReviewID,
		&s.UserID,
		&s.ProductID,
		&s.Content,
		&s.Photos,
		&s.Rating,
		&s.CreatedAt,
		&s.UpdatedAt,
	}
}

// View returns View object for that struct.
func (s *Reviews) View() reform.View {
	return ReviewsTable
}

// Table returns Table object for that record.
func (s *Reviews) Table() reform.Table {
	return ReviewsTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Reviews) PKValue() interface{} {
	return s.ReviewID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Reviews) PKPointer() interface{} {
	return &s.ReviewID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Reviews) HasPK() bool {
	return s.ReviewID != ReviewsTable.z[ReviewsTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.ReviewID = pk.
func (s *Reviews) SetPK(pk interface{}) {
	reform.SetPK(s, pk)
}

// check interfaces
var (
	_ reform.View   = ReviewsTable
	_ reform.Struct = (*Reviews)(nil)
	_ reform.Table  = ReviewsTable
	_ reform.Record = (*Reviews)(nil)
	_ fmt.Stringer  = (*Reviews)(nil)
)

func init() {
	parse.AssertUpToDate(&ReviewsTable.s, new(Reviews))
}
