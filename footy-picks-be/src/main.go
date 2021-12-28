package main

import (
	"encoding/json"
	"footypicks/fotmob"
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
func HealthHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("received GET - /HEALTH request")
	json.NewEncoder(writer).Encode("literally just vibin")
}

func CurrentRoundHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("received GET - /ROUND/CURRENT request")

	league := request.FormValue("league")
	tab := request.FormValue("tab")
	timezone := request.FormValue("timezone")

	response := fotmob.GetCurrentRoundMatches(league, tab, timezone)

	json.NewEncoder(writer).Encode(response)
}

func GamesForUserHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("received GET - /GAMES request")
	user := request.FormValue("user")

	activeGames := GetActiveGamesForUser(user)

	json.NewEncoder(writer).Encode(activeGames)
}

// ROUTES
func (app *App) initRoutes() {
	// health
	app.Router.HandleFunc("/health", HealthHandler).Methods("GET")
	app.Router.HandleFunc("/round/current", CurrentRoundHandler).Methods("GET").Queries(
		"league", "{league}",
		"tab", "{tab}",
		"timezone", "{timezone}")
	app.Router.HandleFunc("/games", GamesForUserHandler).Methods("GET").Queries("user", "{user}")
}

func (a *App) Run(addr string) {
	log.Println("Starting Server at", addr)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"X-Requested-With", "Access-Control-Allow-Origin"},
		AllowedMethods:   []string{"GET", "HEAD", "POST", "PUT", "OPTIONS"},
	})
	handler := c.Handler(a.Router)

	log.Fatal(http.ListenAndServe(addr, handler))
}

// MAIN
func main() {
	app := Init()
	app.Run("0.0.0.0:8000")
}
