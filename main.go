package main

import (
	"fmt"

	"github.com/miha-staric/goka-scraper/config"
	"github.com/miha-staric/goka-scraper/scraper"
	"github.com/miha-staric/goka-scraper/slice"
)

func main() {
	cfg := config.LoadConfig()

	dumpings, err := scraper.FetchDumpings(cfg)
	if err != nil {
		panic(fmt.Sprintf("Failed to fetch dumpings: %v", err))
	}

	switch cfg.Mode {
	case "default":
		slice.SummarizeDumpings(dumpings)
	case "months":
		slice.MonthlyAggregation(dumpings, cfg.CostBio, cfg.CostMko, cfg.MinBio, cfg.MinMko)
	case "years":
		slice.YearlyAggregation(dumpings)
	default:
		fmt.Println("Unknown mode.")
	}
}
