package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

// HANDLERS
func (app *App) healthHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("received /health request")
	json.NewEncoder(writer).Encode("literally just vibin")
}

func (app *App) initRoutes() {
	app.Router.HandleFunc("/health", app.healthHandler).Methods("GET")
}

func (a *App) Run(addr string) {
	log.Println("Starting Server at", addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func main() {
	app := Init()
	app.Run("0.0.0.0:8080")
}
