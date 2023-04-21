package models

import "time"

//go:generate reform

// Groups represents a row in groups table.
//
//reform:groups
type Groups struct {
	GroupID   int32     `reform:"group_id,pk"`
	ParentID  *int32    `reform:"parent_id"`
	Photos    *string   `reform:"photos"`
	Status    bool      `reform:"status"`
	SortOrder int32     `reform:"sort_order"`
	CreatedAt time.Time `reform:"created_at"` // FIXME unhandled database type "timestamp without time zone"
	UpdatedAt time.Time `reform:"updated_at"` // FIXME unhandled database type "timestamp without time zone"
}
