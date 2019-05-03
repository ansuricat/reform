package main

//go:generate reform

// ProductsRegister represents a row in products_register table.
//reform:products_register
type ProductsRegister struct {
	Idproducts  int32   `reform:"idproducts,pk" json:"idproducts"`
	Title       string  `reform:"title" json:"title"`
	Description *string `reform:"description" json:"description"`
	Price       float64 `reform:"price" json:"price"`
}
