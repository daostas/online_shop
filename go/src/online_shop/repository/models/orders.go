package models

//go:generate reform

// Orders represents a row in orders table.
//
//reform:orders
type Orders struct {
	OrderID        int32    `reform:"order_id,pk"`
	UserID         *int32   `reform:"user_id"`
	Address        *string  `reform:"address"`
	Status         *string  `reform:"status"`
	PaymentDetails []byte   `reform:"payment_details"` // FIXME unhandled database type "ARRAY"
	Sum            *float64 `reform:"sum"`
	CreatedAt      []byte   `reform:"created_at"` // FIXME unhandled database type "timestamp without time zone"
	UpdatedAt      []byte   `reform:"updated_at"` // FIXME unhandled database type "timestamp without time zone"
}
