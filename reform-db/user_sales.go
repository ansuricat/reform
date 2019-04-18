package main

//go:generate reform

// UserSales represents a row in user_sales table.
//reform:user_sales
type UserSales struct {
	Idsales int32 `reform:"idsales,pk" json:"idsales"`
}
