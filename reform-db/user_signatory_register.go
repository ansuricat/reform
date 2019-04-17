package main

//go:generate reform

// UserSignatoryRegister represents a row in user_signatory_register table.
//reform:user_signatory_register
type UserSignatoryRegister struct {
	Idsignatoryusers int32   `reform:"idsignatoryusers" json:"idsignatoryusers"`
	ComName          *string `reform:"com_name" json:"com_name"`
	ComEmail         string  `reform:"com_email,pk" json:"com_email"`
	ComCountry       *string `reform:"com_country" json:"com_country"`
	ComState         *int32  `reform:"com_state" json:"com_state"`
	ComCity          *int32  `reform:"com_city" json:"com_city"`
	ComInterestareas *string `reform:"com_interestareas" json:"com_interestareas"`
	ComTelephone     *string `reform:"com_telephone" json:"com_telephone"`
	Pan              *string `reform:"pan" json:"pan"`
	Iec              *string `reform:"iec" json:"iec"`
	Crn              *string `reform:"crn" json:"crn"`
	SignName         *string `reform:"sign_name" json:"sign_name"`
	SignEmail        *string `reform:"sign_email" json:"sign_email"`
	SignBday         *int32  `reform:"sign_bday" json:"sign_bday"`
	SignTelephone    *string `reform:"sign_telephone" json:"sign_telephone"`
	SignSex          *int8   `reform:"sign_sex" json:"sign_sex"`
	SignAddress      *string `reform:"sign_address" json:"sign_address"`
	SignZip          *string `reform:"sign_zip" json:"sign_zip"`
	SignImage        *string `reform:"sign_image" json:"sign_image"`
	User             int32   `reform:"user" json:"user"`
}
