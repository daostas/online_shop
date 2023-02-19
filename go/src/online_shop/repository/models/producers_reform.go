// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type producersTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *producersTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("producers").
func (v *producersTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *producersTableType) Columns() []string {
	return []string{
		"producer_id",
		"photos",
		"status",
		"created_at",
		"updated_at",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *producersTableType) NewStruct() reform.Struct {
	return new(Producers)
}

// NewRecord makes a new record for that table.
func (v *producersTableType) NewRecord() reform.Record {
	return new(Producers)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *producersTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// ProducersTable represents producers view or table in SQL database.
var ProducersTable = &producersTableType{
	s: parse.StructInfo{
		Type:    "Producers",
		SQLName: "producers",
		Fields: []parse.FieldInfo{
			{Name: "ProducerID", Type: "int32", Column: "producer_id"},
			{Name: "Photos", Type: "pq.StringArray", Column: "photos"},
			{Name: "Status", Type: "bool", Column: "status"},
			{Name: "CreatedAt", Type: "time.Time", Column: "created_at"},
			{Name: "UpdatedAt", Type: "time.Time", Column: "updated_at"},
		},
		PKFieldIndex: 0,
	},
	z: new(Producers).Values(),
}

// String returns a string representation of this struct or record.
func (s Producers) String() string {
	res := make([]string, 5)
	res[0] = "ProducerID: " + reform.Inspect(s.ProducerID, true)
	res[1] = "Photos: " + reform.Inspect(s.Photos, true)
	res[2] = "Status: " + reform.Inspect(s.Status, true)
	res[3] = "CreatedAt: " + reform.Inspect(s.CreatedAt, true)
	res[4] = "UpdatedAt: " + reform.Inspect(s.UpdatedAt, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Producers) Values() []interface{} {
	return []interface{}{
		s.ProducerID,
		s.Photos,
		s.Status,
		s.CreatedAt,
		s.UpdatedAt,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Producers) Pointers() []interface{} {
	return []interface{}{
		&s.ProducerID,
		&s.Photos,
		&s.Status,
		&s.CreatedAt,
		&s.UpdatedAt,
	}
}

// View returns View object for that struct.
func (s *Producers) View() reform.View {
	return ProducersTable
}

// Table returns Table object for that record.
func (s *Producers) Table() reform.Table {
	return ProducersTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Producers) PKValue() interface{} {
	return s.ProducerID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Producers) PKPointer() interface{} {
	return &s.ProducerID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Producers) HasPK() bool {
	return s.ProducerID != ProducersTable.z[ProducersTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.ProducerID = pk.
func (s *Producers) SetPK(pk interface{}) {
	reform.SetPK(s, pk)
}

// check interfaces
var (
	_ reform.View   = ProducersTable
	_ reform.Struct = (*Producers)(nil)
	_ reform.Table  = ProducersTable
	_ reform.Record = (*Producers)(nil)
	_ fmt.Stringer  = (*Producers)(nil)
)

func init() {
	parse.AssertUpToDate(&ProducersTable.s, new(Producers))
}
