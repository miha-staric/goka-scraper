// Package types defines data structures used by the VoKa scraper.
package types

type Dumping struct {
	Quantity int    `json:"quantity"`
	Fraction string `json:"fraction"`
	Location string `json:"location"`
	DumpedAtDate string `json:"dumpedAtDate"`
}

type DumpingsProps struct {
	Dumpings struct {
		Dumpings []Dumping `json:"dumpings"`
	} `json:"dumpings"`
}

type InertiaResponse struct {
	Props DumpingsProps `json:"props"`
}

type MonthData struct {
	BIO int
	MKO int
}

type YearData struct {
	BIO int
	MKO int
}
