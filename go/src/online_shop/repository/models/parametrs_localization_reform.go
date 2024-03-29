// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type parametrsLocalizationViewType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *parametrsLocalizationViewType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("parametrs_localization").
func (v *parametrsLocalizationViewType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *parametrsLocalizationViewType) Columns() []string {
	return []string{
		"parametr_id",
		"lang_id",
		"title",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *parametrsLocalizationViewType) NewStruct() reform.Struct {
	return new(ParametrsLocalization)
}

// ParametrsLocalizationView represents parametrs_localization view or table in SQL database.
var ParametrsLocalizationView = &parametrsLocalizationViewType{
	s: parse.StructInfo{
		Type:    "ParametrsLocalization",
		SQLName: "parametrs_localization",
		Fields: []parse.FieldInfo{
			{Name: "ParametrID", Type: "*int32", Column: "parametr_id"},
			{Name: "LangID", Type: "*int32", Column: "lang_id"},
			{Name: "Title", Type: "string", Column: "title"},
		},
		PKFieldIndex: -1,
	},
	z: new(ParametrsLocalization).Values(),
}

// String returns a string representation of this struct or record.
func (s ParametrsLocalization) String() string {
	res := make([]string, 3)
	res[0] = "ParametrID: " + reform.Inspect(s.ParametrID, true)
	res[1] = "LangID: " + reform.Inspect(s.LangID, true)
	res[2] = "Title: " + reform.Inspect(s.Title, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *ParametrsLocalization) Values() []interface{} {
	return []interface{}{
		s.ParametrID,
		s.LangID,
		s.Title,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *ParametrsLocalization) Pointers() []interface{} {
	return []interface{}{
		&s.ParametrID,
		&s.LangID,
		&s.Title,
	}
}

// View returns View object for that struct.
func (s *ParametrsLocalization) View() reform.View {
	return ParametrsLocalizationView
}

// check interfaces
var (
	_ reform.View   = ParametrsLocalizationView
	_ reform.Struct = (*ParametrsLocalization)(nil)
	_ fmt.Stringer  = (*ParametrsLocalization)(nil)
)

func init() {
	parse.AssertUpToDate(&ParametrsLocalizationView.s, new(ParametrsLocalization))
}
