package models

import "github.com/lib/pq"

//go:generate reform

// Groups represents a row in groups table.
//
//reform:groups
type Groups struct {
	GroupID  int32          `reform:"group_id,pk"`
	ParentID *int32         `reform:"parent_id"`
	Photos   pq.StringArray `reform:"photos"` // FIXME unhandled database type "ARRAY"
	Status   bool           `reform:"status"`
}
