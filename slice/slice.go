package slice

import (
	"fmt"
	"time"

	"github.com/go-gota/gota/dataframe"
	"github.com/miha-staric/goka-scraper/utils"
	"github.com/miha-staric/goka-scraper/types"
)

func SummarizeDumpings(dumpings []types.Dumping) {
	df := dataframe.LoadStructs(dumpings)

	fmt.Println("Full data:")
	fmt.Println(df)

	// Sum quantities by fraction
	bioSum := df.Filter(dataframe.F{Colname: "Fraction", Comparator: "==", Comparando: "BIO"}).Col("Quantity").Sum()
	mkoSum := df.Filter(dataframe.F{Colname: "Fraction", Comparator: "==", Comparando: "MKO"}).Col("Quantity").Sum()

	fmt.Println("Summarized data:")
	fmt.Printf("BIO: %.0f, MKO: %.0f\n", bioSum, mkoSum)
}

func MonthlyAggregation(dumpings []types.Dumping, costBio float64, costMko float64, minBio float64, minMko float64) {
	costMap := map[string]float64{
		"BIO": costBio,
		"MKO": costMko,
	}
	minCosts := map[string]float64{
		"BIO": minBio,
		"MKO": minMko,
	}

	// Group by month
	months := make(map[string]*types.MonthData)

	for _, d := range dumpings {
		t, err := time.Parse("2006-01-02", d.DumpedAtDate)
		if err != nil {
			continue
		}
		month := t.Format("2006-01")

		if _, exists := months[month]; !exists {
			months[month] = &types.MonthData{}
		}

		switch d.Fraction {
		case "BIO":
			months[month].BIO += d.Quantity
		case "MKO":
			months[month].MKO += d.Quantity
		}
	}

	fmt.Println("\nMonths data:")
	fmt.Printf("%-10s %-5s %-5s %-12s %-12s\n", "month", "BIO", "MKO", "real_cost", "total_cost")

	for month, data := range months {
		// Real costs
		realBIO := float64(data.BIO) * costMap["BIO"]
		realMKO := float64(data.MKO) * costMap["MKO"]
		realCost := realBIO + realMKO

		// Apply minimums
		totalBIO := utils.Max(realBIO, minCosts["BIO"])
		totalMKO := utils.Max(realMKO, minCosts["MKO"])
		totalCost := totalBIO + totalMKO

		fmt.Printf("%-10s %-5d %-5d %-12.4f %-12.4f\n",
			month,
			data.BIO,
			data.MKO,
			realCost,
			totalCost,
		)
	}
}

func YearlyAggregation(dumpings []types.Dumping) {
	years := make(map[string]*types.YearData)

	// Group + sum
	for _, d := range dumpings {
		t, err := time.Parse("2006-01-02", d.DumpedAtDate)
		if err != nil {
			continue
		}

		year := t.Format("2006")

		if _, exists := years[year]; !exists {
			years[year] = &types.YearData{}
		}

		switch d.Fraction {
		case "BIO":
			years[year].BIO += d.Quantity
		case "MKO":
			years[year].MKO += d.Quantity
		}
	}

	// Print like pandas
	fmt.Println("\nYears data:")
	fmt.Printf("%-6s %-5s %-5s\n", "year", "BIO", "MKO")

	for year, data := range years {
		fmt.Printf("%-6s %-5d %-5d\n",
			year,
			data.BIO,
			data.MKO,
		)
	}
}
