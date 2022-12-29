package models

//go:generate reform

// Admins represents a row in admins table.
//
//reform:admins
type Admins struct {
	AdminID  int32  `reform:"admin_id,pk"`
	Login    string `reform:"login"`
	Password string `reform:"password"`
	Role     string `reform:"role"`
}
