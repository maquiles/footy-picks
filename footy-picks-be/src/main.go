package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type App struct {
	Router *mux.Router
}

func Init() *App {
	app := &App{
		Router: mux.NewRouter(),
	}
	app.initRoutes()
	return app
}

// ROUTE HANDLERS
func (app *App) HealthHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("received /health request")
	json.NewEncoder(writer).Encode("literally just vibin")
}

func (app *App) MatchDataRefreshHandler(writer http.ResponseWriter, request *http.Request) {
	league := request.FormValue("league")
	tab := request.FormValue("tab")
	timeZone := request.FormValue("timezone")

	response := GetLeagueMatches(league, tab, timeZone)
	json.NewEncoder(writer).Encode(response)
}

// ROUTES
func (app *App) initRoutes() {
	// health
	app.Router.HandleFunc("/health", app.HealthHandler).Methods("GET")

	// get fotmob data
	app.Router.HandleFunc("/match-data/refresh", app.MatchDataRefreshHandler).Methods("GET").Queries(
		"league", "{league}",
		"tab", "{tab}",
		"timezone", "{timezone}")
}

func (a *App) Run(addr string) {
	log.Println("Starting Server at", addr)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
	})
	handler := c.Handler(a.Router)

	log.Fatal(http.ListenAndServe(addr, handler))
}

// MAIN
func main() {
	app := Init()
	app.Run("0.0.0.0:8000")
}
