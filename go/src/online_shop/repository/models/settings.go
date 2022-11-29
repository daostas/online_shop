package models

//go:generate reform

// Settings represents a row in settings table.
//
//reform:settings
type Settings struct {
	Key   string `reform:"key,pk"`
	Value string `reform:"value"`
}
