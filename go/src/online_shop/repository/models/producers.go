package models

import (
	"time"

	"github.com/lib/pq"
)

//go:generate reform

// Producers represents a row in producers table.
//
//reform:producers
type Producers struct {
	ProducerID int32          `reform:"producer_id,pk"`
	Photos     pq.StringArray `reform:"photos"` // FIXME unhandled database type "ARRAY"
	Status     bool           `reform:"status"`
	CreatedAt  time.Time      `reform:"created_at"` // FIXME unhandled database type "timestamp without time zone"
	UpdatedAt  time.Time      `reform:"updated_at"` // FIXME unhandled database type "timestamp without time zone"
}
