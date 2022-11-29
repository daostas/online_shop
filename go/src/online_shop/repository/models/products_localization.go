package models

//go:generate reform

// ProductsLocalization represents a row in products_localization table.
//
//reform:products_localization
type ProductsLocalization struct {
	ProductID   *int32  `reform:"product_id"`
	LangID      *int32  `reform:"lang_id"`
	Title       string  `reform:"title"`
	Description *string `reform:"description"`
}
