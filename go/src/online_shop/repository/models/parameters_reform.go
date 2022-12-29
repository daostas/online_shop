// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type parametersTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *parametersTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("parameters").
func (v *parametersTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *parametersTableType) Columns() []string {
	return []string{
		"parameter",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *parametersTableType) NewStruct() reform.Struct {
	return new(Parameters)
}

// NewRecord makes a new record for that table.
func (v *parametersTableType) NewRecord() reform.Record {
	return new(Parameters)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *parametersTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// ParametersTable represents parameters view or table in SQL database.
var ParametersTable = &parametersTableType{
	s: parse.StructInfo{
		Type:    "Parameters",
		SQLName: "parameters",
		Fields: []parse.FieldInfo{
			{Name: "Parameter", Type: "string", Column: "parameter"},
		},
		PKFieldIndex: 0,
	},
	z: new(Parameters).Values(),
}

// String returns a string representation of this struct or record.
func (s Parameters) String() string {
	res := make([]string, 1)
	res[0] = "Parameter: " + reform.Inspect(s.Parameter, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Parameters) Values() []interface{} {
	return []interface{}{
		s.Parameter,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Parameters) Pointers() []interface{} {
	return []interface{}{
		&s.Parameter,
	}
}

// View returns View object for that struct.
func (s *Parameters) View() reform.View {
	return ParametersTable
}

// Table returns Table object for that record.
func (s *Parameters) Table() reform.Table {
	return ParametersTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Parameters) PKValue() interface{} {
	return s.Parameter
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Parameters) PKPointer() interface{} {
	return &s.Parameter
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Parameters) HasPK() bool {
	return s.Parameter != ParametersTable.z[ParametersTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.Parameter = pk.
func (s *Parameters) SetPK(pk interface{}) {
	reform.SetPK(s, pk)
}

// check interfaces
var (
	_ reform.View   = ParametersTable
	_ reform.Struct = (*Parameters)(nil)
	_ reform.Table  = ParametersTable
	_ reform.Record = (*Parameters)(nil)
	_ fmt.Stringer  = (*Parameters)(nil)
)

func init() {
	parse.AssertUpToDate(&ParametersTable.s, new(Parameters))
}