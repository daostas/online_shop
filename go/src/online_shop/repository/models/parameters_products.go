package models

//go:generate reform

// ParametersProducts represents a row in parameters_products table.
//
//reform:parameters_products
type ParametersProducts struct {
	ProductID *int32  `reform:"product_id"`
	Parameter *string `reform:"parameter"`
	Value     *string `reform:"value"`
}
