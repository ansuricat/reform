package main

//go:generate reform

// TechRegions represents a row in tech_regions table.
//reform:tech_regions
type TechRegions struct {
	Idregions int32  `reform:"idregions,pk" json:"idregions"`
	Name      string `reform:"name" json:"name"`
	Country   string `reform:"country" json:"country"`
}
