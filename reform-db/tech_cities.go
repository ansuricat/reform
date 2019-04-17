package main

//go:generate reform

// TechCities represents a row in tech_cities table.
//reform:tech_cities
type TechCities struct {
	Idcities int32  `reform:"idcities" json:"idcities"`
	Name     string `reform:"name" json:"name"`
	Region   int32  `reform:"region,pk" json:"region"`
	Country  string `reform:"country" json:"country"`
}
