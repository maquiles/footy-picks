package main

import (
	"footypicks/repo"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type App struct {
	Router *mux.Router
	DBConn repo.Repo
}

// ROUTES
func (app *App) initRoutes() {
	// health
	app.Router.HandleFunc("/health", HealthHandler).Methods("GET")
	app.Router.HandleFunc("/round/current", CurrentRoundHandler).Methods("GET").Queries(
		"league", "{league}",
		"tab", "{tab}",
		"timezone", "{timezone}")
	app.Router.HandleFunc("/games", app.GamesForUserHandler).Methods("GET").Queries("user", "{user}")

	// player
	app.Router.HandleFunc("/player", app.CreateNewPlayerHandler).Methods("POST")
	app.Router.HandleFunc("/player/game", app.AddPlayerGameHandler).Methods("POST")

	// login
	app.Router.HandleFunc("/login", app.LoginHandler).Methods("POST")
}

func Init() *App {
	app := &App{
		Router: mux.NewRouter(),
		DBConn: repo.InitDBConn(),
	}
	app.initRoutes()
	return app
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
