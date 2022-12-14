// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type parametersGroupsViewType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *parametersGroupsViewType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("parameters_groups").
func (v *parametersGroupsViewType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *parametersGroupsViewType) Columns() []string {
	return []string{
		"group_id",
		"parameter",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *parametersGroupsViewType) NewStruct() reform.Struct {
	return new(ParametersGroups)
}

// ParametersGroupsView represents parameters_groups view or table in SQL database.
var ParametersGroupsView = &parametersGroupsViewType{
	s: parse.StructInfo{
		Type:    "ParametersGroups",
		SQLName: "parameters_groups",
		Fields: []parse.FieldInfo{
			{Name: "GroupID", Type: "*int32", Column: "group_id"},
			{Name: "Parameter", Type: "*string", Column: "parameter"},
		},
		PKFieldIndex: -1,
	},
	z: new(ParametersGroups).Values(),
}

// String returns a string representation of this struct or record.
func (s ParametersGroups) String() string {
	res := make([]string, 2)
	res[0] = "GroupID: " + reform.Inspect(s.GroupID, true)
	res[1] = "Parameter: " + reform.Inspect(s.Parameter, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *ParametersGroups) Values() []interface{} {
	return []interface{}{
		s.GroupID,
		s.Parameter,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *ParametersGroups) Pointers() []interface{} {
	return []interface{}{
		&s.GroupID,
		&s.Parameter,
	}
}

// View returns View object for that struct.
func (s *ParametersGroups) View() reform.View {
	return ParametersGroupsView
}

// check interfaces
var (
	_ reform.View   = ParametersGroupsView
	_ reform.Struct = (*ParametersGroups)(nil)
	_ fmt.Stringer  = (*ParametersGroups)(nil)
)

func init() {
	parse.AssertUpToDate(&ParametersGroupsView.s, new(ParametersGroups))
}
