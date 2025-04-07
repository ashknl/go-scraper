package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	jsonRes := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	// js, err := json.Marshal(jsonRes)

	// w.Header().Set("Content-Type", "application/json")

	// w.Write([]byte(js))
	err := app.writeJSON(w, http.StatusOK, jsonRes, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

}
