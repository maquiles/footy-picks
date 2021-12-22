package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (app *App) HealthHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("received /health request")
	json.NewEncoder(writer).Encode("literally just vibin")
}

func (app *App) MatchDataRefreshHandler(writer http.ResponseWriter, request *http.Request) {
	response := getLeague("42", "overview", "league", "America/New_York")
	json.NewEncoder(writer).Encode(response)
}
