package models

//go:generate reform

// Baskets represents a row in baskets table.
//
//reform:baskets
type Baskets struct {
	BasketID int32 `reform:"basket_id,pk"`
	UserID   int32 `reform:"user_id"`
}
