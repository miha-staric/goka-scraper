// Package config holds configuration using environment variables or .env file
package config

import (
	"fmt"
	"os"
)

import (
	"github.com/joho/godotenv"
)

type Config struct {
	Card string
	Pass string
	LoginUrl string
	DashboardUrl string
	BaseUrl string
}

func LoadConfig() Config {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("No .env file found")
	}

	card := os.Getenv("GOKA_CARD")
	pass := os.Getenv("GOKA_PASS")
	loginUrl := os.Getenv("GOKA_LOGINURL")
	dashboardUrl := os.Getenv("GOKA_DASHBOARDURL")
	baseUrl := os.Getenv("GOKA_BASEURL")

	if card == "" || pass == "" || loginUrl == "" || dashboardUrl == "" || baseUrl == "" {
		panic("Missing environment variables (GOKA_CARD, GOKA_PASS, GOKA_LOGINURL, GOKA_DASHBOARDURL, GOKA_BASEURL")
	}

	return Config{
		Card: card,
		Pass: pass,
		LoginUrl: loginUrl,
		DashboardUrl: dashboardUrl,
		BaseUrl: baseUrl,
	}
}
