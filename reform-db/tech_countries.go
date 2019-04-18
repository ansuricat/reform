package main

//go:generate reform

// TechCountries represents a row in tech_countries table.
//reform:tech_countries
type TechCountries struct {
	Name      string  `reform:"name" json:"name"`
	Idcode    string  `reform:"idcode,pk" json:"idcode"`
	ZipRegexp *string `reform:"zip_regexp" json:"zip_regexp"`
}
