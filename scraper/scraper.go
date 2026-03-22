package scraper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/gocolly/colly/v2"
)

import (
	"github.com/miha-staric/goka-scraper/config"
	"github.com/miha-staric/goka-scraper/types"
	"github.com/miha-staric/goka-scraper/utils"
)

// FetchDumpings logs in and returns cleaned dumpings
func FetchDumpings(cfg config.Config) ([]types.Dumping, error) {
	c := colly.NewCollector(colly.AllowURLRevisit())
	jar, _ := cookiejar.New(nil)
	c.SetCookieJar(jar)

	var xsrfToken, inertiaVersion string
	var dumpings []types.Dumping

	// Parse Inertia version from login page
	c.OnHTML("#app", func(e *colly.HTMLElement) {
		dataPage := e.Attr("data-page")
		var parsed map[string]interface{}
		json.Unmarshal([]byte(dataPage), &parsed)
		if v, ok := parsed["version"].(string); ok {
			inertiaVersion = v
		}
	})

	// Extract XSRF token from cookies
	c.OnResponse(func(r *colly.Response) {
		u, _ := url.Parse(cfg.BaseUrl)
		for _, cookie := range jar.Cookies(u) {
			if cookie.Name == "XSRF-TOKEN" {
				decoded, _ := url.QueryUnescape(cookie.Value)
				xsrfToken = decoded
			}
		}
	})

	// Visit login page
	if err := c.Visit(cfg.LoginUrl); err != nil {
		return nil, fmt.Errorf("failed to visit login page: %w", err)
	}
	c.Wait()

	if xsrfToken == "" || inertiaVersion == "" {
		return nil, fmt.Errorf("failed to extract XSRF token or Inertia version")
	}

	// POST login
	payload := map[string]string{"chipCardNumber": cfg.Card, "password": cfg.Pass}
	jsonData, _ := json.Marshal(payload)

	if err := c.Request(
		"POST",
		cfg.LoginUrl,
		strings.NewReader(string(jsonData)),
		nil,
		http.Header{
			"Content-Type":      []string{"application/json"},
			"X-Inertia":         []string{"true"},
			"X-Inertia-Version": []string{inertiaVersion},
			"X-XSRF-TOKEN":      []string{xsrfToken},
			"Accept":            []string{"application/json"},
		},
	); err != nil {
		return nil, fmt.Errorf("login request failed: %w", err)
	}
	c.Wait()

	// GET dashboard
	c.OnResponse(func(r *colly.Response) {
		if strings.Contains(r.Request.URL.String(), "dashboard") {
			var resp types.InertiaResponse
			if err := json.Unmarshal(r.Body, &resp); err != nil {
				fmt.Println("Failed to parse dashboard JSON:", err)
				return
			}

			dumpings = resp.Props.Dumpings.Dumpings
			for i := range dumpings {
				if dumpings[i].Fraction == "common.REST 2" {
					dumpings[i].Fraction = "MKO"
				}
				dumpings[i].Location = utils.NormalizeSpaces(dumpings[i].Location)
			}
		}
	})

	if err := c.Request("GET", cfg.DashboardUrl, nil, nil, http.Header{
		"X-Inertia":         []string{"true"},
		"X-Inertia-Version": []string{inertiaVersion},
		"X-XSRF-TOKEN":      []string{xsrfToken},
		"Accept":            []string{"application/json"},
	}); err != nil {
		return nil, fmt.Errorf("failed to GET dashboard: %w", err)
	}

	c.Wait()
	return dumpings, nil
}
