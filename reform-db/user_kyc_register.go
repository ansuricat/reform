package main

//go:generate reform

// UserKycRegister represents a row in user_kyc_register table.
//reform:user_kyc_register
type UserKycRegister struct {
	Idkycregister    int32   `reform:"idkycregister,pk" json:"idkycregister"`
	ComName          *string `reform:"com_name" json:"com_name"`
	ComEmail         string  `reform:"com_email" json:"com_email"`
	ComCountry       *string `reform:"com_country" json:"com_country"`
	ComState         *int32  `reform:"com_state" json:"com_state"`
	ComCity          *int32  `reform:"com_city" json:"com_city"`
	ComInterestareas *string `reform:"com_interestareas" json:"com_interestareas"`
	ComTelephone     *string `reform:"com_telephone" json:"com_telephone"`
	Pan              *string `reform:"pan" json:"pan"`
	Iec              *string `reform:"iec" json:"iec"`
	Crn              *string `reform:"crn" json:"crn"`
	User             int32   `reform:"user" json:"user"`
}
