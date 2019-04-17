package main

//go:generate reform

// TechCountries represents a row in tech_countries table.
//reform:tech_countries
type TechCountries struct {
	Name      string  `reform:"name,pk" json:"name"`
	Code      string  `reform:"code" json:"code"`
	ZipRegexp *string `reform:"zip_regexp" json:"zip_regexp"`
}
