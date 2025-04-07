package main

import (
	"fmt"
	"net/http"

	"github.com/ashknl/go-scraper/internal/scraper"
)

func (app *application) scrapeRequestHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Term        string `json:"term"`
		CountryCode string `json:"cc_code"`
		LangCode    string `json:"lang_code"`
		Pages       int    `json:"pages"`
		Count       int    `json:"count"`
		Backoff     int    `json:"backoff"`
	}

	err := app.readJSON(w, r, &input)

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	resultsMap := make(map[int]scraper.SearchResult)

	results, err := scraper.GoogleScrape(input.Term, input.CountryCode, input.LangCode, nil, input.Pages, input.Count, input.Backoff)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for i, result := range results {
		resultsMap[i] = result
	}

	err = app.writeJSON(w, http.StatusOK, resultsMap, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
