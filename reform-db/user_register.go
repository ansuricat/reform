package main

//go:generate reform

// UserRegister represents a row in user_register table.
//reform:user_register
type UserRegister struct {
	Idusers       int32   `reform:"idusers,pk" json:"idusers"`
	Password      string  `reform:"password" json:"password"`
	Fullname      string  `reform:"fullname" json:"fullname"`
	Email         string  `reform:"email" json:"email"`
	Companyname   *string `reform:"companyname" json:"companyname"`
	Interestareas *string `reform:"interestareas" json:"interestareas"`
	Type          int8    `reform:"type" json:"type"`
	Kycstatus     int8    `reform:"kycstatus" json:"kycstatus"`
	Crn           *string `reform:"crn" json:"crn"`
	Pan           *string `reform:"pan" json:"pan"`
	Iec           *string `reform:"iec" json:"iec"`
	Address       *string `reform:"address" json:"address"`
	Country       *string `reform:"country" json:"country"`
	State         *int32  `reform:"state" json:"state"`
	City          *int32  `reform:"city" json:"city"`
	Zip           *string `reform:"zip" json:"zip"`
	Image         *string `reform:"image" json:"image"`
	Idxion        *string `reform:"idxion" json:"idxion"`
	EthPub        string  `reform:"eth_pub" json:"eth_pub"`
	Website       *string `reform:"website" json:"website"`
}
