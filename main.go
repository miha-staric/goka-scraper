package main

import (
	"fmt"
)

import (
	"github.com/miha-staric/goka-scraper/config"
	"github.com/miha-staric/goka-scraper/scraper"
)

func main() {
	cfg := config.LoadConfig()

	dumpings, err := scraper.FetchDumpings(cfg)
	if err != nil {
		panic(fmt.Sprintf("Failed to fetch dumpings: %v", err))
	}

	fmt.Printf("Dumpings fetched: %+v\n", dumpings)
}
