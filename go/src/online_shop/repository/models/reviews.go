package models

//go:generate reform

// Reviews represents a row in reviews table.
//
//reform:reviews
type Reviews struct {
	ReviewID  int32   `reform:"review_id,pk"`
	UserID    *int32  `reform:"user_id"`
	ProductID *int32  `reform:"product_id"`
	Content   *string `reform:"content"`
	Photos    []byte  `reform:"photos"` // FIXME unhandled database type "ARRAY"
	Rating    float64 `reform:"rating"`
	CreatedAt []byte  `reform:"created_at"` // FIXME unhandled database type "timestamp without time zone"
	UpdatedAt []byte  `reform:"updated_at"` // FIXME unhandled database type "timestamp without time zone"
}
