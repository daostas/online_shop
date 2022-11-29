package models

//go:generate reform

// Favourites represents a row in favourites table.
//
//reform:favourites
type Favourites struct {
	FavouriteID int32 `reform:"favourite_id,pk"`
	UserID      int32 `reform:"user_id"`
}
