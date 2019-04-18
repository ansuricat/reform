package main

//go:generate reform

// UserKycSign represents a row in user_kyc_sign table.
//reform:user_kyc_sign
type UserKycSign struct {
	Iduserkycsign int32   `reform:"iduserkycsign,pk" json:"iduserkycsign"`
	SignName      string  `reform:"sign_name" json:"sign_name"`
	SignEmail     string  `reform:"sign_email" json:"sign_email"`
	SignBday      *int32  `reform:"sign_bday" json:"sign_bday"`
	SignTelephone *string `reform:"sign_telephone" json:"sign_telephone"`
	SignSex       *int8   `reform:"sign_sex" json:"sign_sex"`
	SignAddress   *string `reform:"sign_address" json:"sign_address"`
	SignZip       *string `reform:"sign_zip" json:"sign_zip"`
	SignImage     *string `reform:"sign_image" json:"sign_image"`
	UserKyc       int32   `reform:"user_kyc" json:"user_kyc"`
}
