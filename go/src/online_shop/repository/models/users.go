package models

import (
	"time"
)

//go:generate reform

// Users represents a row in users table.
//
//reform:users
type Users struct {
	UserID       int32      `reform:"user_id,pk"`
	UserName     *string    `reform:"user_name"`
	Number       *string    `reform:"number"`
	Email        *string    `reform:"email"`
	Dob          *time.Time `reform:"dob"`
	Address      *string    `reform:"address"`
	UserPassword string     `reform:"user_password"`
}
