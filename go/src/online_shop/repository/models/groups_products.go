package models

//go:generate reform

// GroupsProducts represents a row in groups_products table.
//
//reform:groups_products
type GroupsProducts struct {
	GroupID   *int32 `reform:"group_id"`
	ProductID *int32 `reform:"product_id"`
}
