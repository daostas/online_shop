package models

//go:generate reform

// ParametrsGroups represents a row in parametrs_groups table.
//
//reform:parametrs_groups
type ParametrsGroups struct {
	GroupID    *int32 `reform:"group_id"`
	ParametrID *int32 `reform:"parametr_id"`
}
