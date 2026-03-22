// Package config holds configuration using environment variables or .env file
package config

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Mode string
	LogLevel string
	Card string
	Pass string
	LoginUrl string
	DashboardUrl string
	BaseUrl string
	DateFrom time.Time
	DateTo time.Time
	CostMko float64
	CostBio float64
	MinMko float64
	MinBio float64
}

func LoadConfig() Config {
	error := godotenv.Load()

	if error != nil {
		fmt.Println("No .env file found")
	}

	mode := os.Getenv("GOKA_MODE")
	logLevel := os.Getenv("GOKA_LOGLEVEL")
	card := os.Getenv("GOKA_CARD")
	pass := os.Getenv("GOKA_PASS")
	loginUrl := os.Getenv("GOKA_LOGINURL")
	dashboardUrl := os.Getenv("GOKA_DASHBOARDURL")
	baseUrl := os.Getenv("GOKA_BASEURL")
	dateFrom := os.Getenv("GOKA_DATEFROM")
	dateTo := os.Getenv("GOKA_DATETO")
	costBio := os.Getenv("GOKA_COSTBIO")
	costMko := os.Getenv("GOKA_COSTMKO")
	minBio := os.Getenv("GOKA_MINBIO")
	minMko := os.Getenv("GOKA_MINMKO")


	if mode == "" || logLevel == "" || card == "" || pass == "" || loginUrl == "" || dashboardUrl == "" || baseUrl == "" || dateFrom == "" || dateTo == "" || costBio == "" || costMko == "" || minBio == "" || minMko == "" {
		panic("Missing environment variables (GOKA_MODE, GOKA_LOGLEVEL, GOKA_CARD, GOKA_PASS, GOKA_LOGINURL, GOKA_DASHBOARDURL, GOKA_BASEURL, GOKA_DATEFROM, GOKA_DATETO, GOKA_COSTBIO, GOKA_COSTMKO, GOKA_MINBIO, GOKA_MINMKO")
	}

	mode = strings.ToLower(mode)
	logLevel = strings.ToLower(logLevel)

	dateFromParsed, error := time.Parse("2006-01-02", dateFrom)
	if error != nil {
		panic("GOKA_DATEFROM in incorrect format, expected YYYY-MM-DD.")
	}

	dateToParsed, error := time.Parse("2006-01-02", dateTo)
	if error != nil {
		panic("GOKA_DATETO in incorrect format, expected YYYY-MM-DD.")
	}

	costBioFloat, err := strconv.ParseFloat(costBio, 64)
	if err != nil {
		panic(err)
	}

	costMkoFloat, err := strconv.ParseFloat(costMko, 64)
	if err != nil {
		panic(err)
	}

	minBioFloat, err := strconv.ParseFloat(minBio, 64)
	if err != nil {
		panic(err)
	}

	minMkoFloat, err := strconv.ParseFloat(minMko, 64)
	if err != nil {
		panic(err)
	}

	return Config{
		Mode: mode,
		LogLevel: logLevel,
		Card: card,
		Pass: pass,
		LoginUrl: loginUrl,
		DashboardUrl: dashboardUrl,
		BaseUrl: baseUrl,
		DateFrom: dateFromParsed,
		DateTo: dateToParsed,
		CostMko: costMkoFloat,
		CostBio: costBioFloat,
		MinMko: minMkoFloat,
		MinBio: minBioFloat,
	}
}
