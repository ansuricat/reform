package main

//go:generate reform

// Products represents a row in products table.
//reform:products
type Products struct {
	Idproducts int32 `reform:"idproducts,pk" json:"idproducts"`
}
