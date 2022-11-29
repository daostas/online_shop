package models

//go:generate reform

// Groups represents a row in groups table.
//
//reform:groups
type Groups struct {
	GroupID  int32  `reform:"group_id,pk"`
	ParentID *int32 `reform:"parent_id"`
	Photos   []byte `reform:"photos"` // FIXME unhandled database type "ARRAY"
	Status   bool   `reform:"status"`
}
