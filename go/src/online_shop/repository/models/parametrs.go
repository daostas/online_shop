package models

import "time"

//go:generate reform

// Parametrs represents a row in parametrs table.
//
//reform:parametrs
type Parametrs struct {
	ParametrID int32     `reform:"parametr_id,pk"`
	Status     bool      `reform:"status"`
	CreatedAt  time.Time `reform:"created_at"` // FIXME unhandled database type "timestamp without time zone"
	UpdatedAt  time.Time `reform:"updated_at"` // FIXME unhandled database type "timestamp without time zone"
}
