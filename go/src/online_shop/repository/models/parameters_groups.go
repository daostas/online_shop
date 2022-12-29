package models

//go:generate reform

// ParametersGroups represents a row in parameters_groups table.
//
//reform:parameters_groups
type ParametersGroups struct {
	GroupID   *int32  `reform:"group_id"`
	Parameter *string `reform:"parameter"`
}
