package models

//go:generate reform

// FavouritesProducts represents a row in favourites_products table.
//
//reform:favourites_products
type FavouritesProducts struct {
	FavouriteID *int32 `reform:"favourite_id"`
	ProductID   *int32 `reform:"product_id"`
}
