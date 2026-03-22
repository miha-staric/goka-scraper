// Package types defines data structures used by the VoKa scraper.
package types

// Dumping represents a single dumping entry.
type Dumping struct {
	Quantity int    `json:"quantity"`
	Fraction string `json:"fraction"`
	Location string `json:"location"`
}

// DumpingsProps mirrors the structure of the JSON payload.
type DumpingsProps struct {
	Dumpings struct {
		Dumpings []Dumping `json:"dumpings"`
	} `json:"dumpings"`
}

// InertiaResponse wraps the props object returned by the API.
type InertiaResponse struct {
	Props DumpingsProps `json:"props"`
}
