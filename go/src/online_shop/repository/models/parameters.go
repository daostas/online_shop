package models

//go:generate reform

// Parameters represents a row in parameters table.
//
//reform:parameters
type Parameters struct {
	Parameter string `reform:"parameter,pk"`
}
