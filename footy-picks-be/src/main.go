package main

import (
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

func (app *App) initRoutes() {
	// health
	app.Router.HandleFunc("/health", app.HealthHandler).Methods("GET")

	// get fotmob data
	app.Router.HandleFunc("/match-data/refresh", app.MatchDataRefreshHandler).Methods("GET")
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

func main() {
	app := Init()
	app.Run("0.0.0.0:8000")
}
