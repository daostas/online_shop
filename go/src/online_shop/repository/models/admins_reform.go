// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type adminsTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *adminsTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("admins").
func (v *adminsTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *adminsTableType) Columns() []string {
	return []string{
		"admin_id",
		"login",
		"password",
		"role",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *adminsTableType) NewStruct() reform.Struct {
	return new(Admins)
}

// NewRecord makes a new record for that table.
func (v *adminsTableType) NewRecord() reform.Record {
	return new(Admins)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *adminsTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// AdminsTable represents admins view or table in SQL database.
var AdminsTable = &adminsTableType{
	s: parse.StructInfo{
		Type:    "Admins",
		SQLName: "admins",
		Fields: []parse.FieldInfo{
			{Name: "AdminID", Type: "int32", Column: "admin_id"},
			{Name: "Login", Type: "string", Column: "login"},
			{Name: "Password", Type: "string", Column: "password"},
			{Name: "Role", Type: "string", Column: "role"},
		},
		PKFieldIndex: 0,
	},
	z: new(Admins).Values(),
}

// String returns a string representation of this struct or record.
func (s Admins) String() string {
	res := make([]string, 4)
	res[0] = "AdminID: " + reform.Inspect(s.AdminID, true)
	res[1] = "Login: " + reform.Inspect(s.Login, true)
	res[2] = "Password: " + reform.Inspect(s.Password, true)
	res[3] = "Role: " + reform.Inspect(s.Role, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Admins) Values() []interface{} {
	return []interface{}{
		s.AdminID,
		s.Login,
		s.Password,
		s.Role,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Admins) Pointers() []interface{} {
	return []interface{}{
		&s.AdminID,
		&s.Login,
		&s.Password,
		&s.Role,
	}
}

// View returns View object for that struct.
func (s *Admins) View() reform.View {
	return AdminsTable
}

// Table returns Table object for that record.
func (s *Admins) Table() reform.Table {
	return AdminsTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Admins) PKValue() interface{} {
	return s.AdminID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Admins) PKPointer() interface{} {
	return &s.AdminID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Admins) HasPK() bool {
	return s.AdminID != AdminsTable.z[AdminsTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.AdminID = pk.
func (s *Admins) SetPK(pk interface{}) {
	reform.SetPK(s, pk)
}

// check interfaces
var (
	_ reform.View   = AdminsTable
	_ reform.Struct = (*Admins)(nil)
	_ reform.Table  = AdminsTable
	_ reform.Record = (*Admins)(nil)
	_ fmt.Stringer  = (*Admins)(nil)
)

func init() {
	parse.AssertUpToDate(&AdminsTable.s, new(Admins))
}
