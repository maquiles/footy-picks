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

	// fotmob
	app.Router.HandleFunc("/round/current", CurrentRoundHandler).Methods("GET").Queries(
		"league", "{league}",
		"tab", "{tab}",
		"timezone", "{timezone}")

	// game
	app.Router.HandleFunc("/games/{player_id}", app.GamesForUserHandler).Methods("GET")
	app.Router.HandleFunc("/game", app.NewGameHandler).Methods("POST")
	// TODO app.Router.HandleFunc("/game/player", app.AddPlayerToGameHandler).Methods("POST")
	// TODO app.Router.HandleFunc("/game/pick", app.MakeGamePickHandler).Methods("POST")
	// TODO app.Router.HandleFunc("/game/{game_id}", app.GetGameHandler).Methods("GET")
	// TODO app.Router.HandleFunc("/game/{game_id}/players", app.GetPlayersForGameHandler).Methods("GET")
	// TODO app.Router.HandleFunc("/game/{game_id}/table", app.GetTableForGame).Methods("GET")
	// TODO app.Router.HandleFunc("/game/{game_id}/picks", app.GetGamePicksHandler).Methods("GET")

	// player
	app.Router.HandleFunc("/player", app.CreateNewPlayerHandler).Methods("POST")
	app.Router.HandleFunc("/player/game", app.AddPlayerGameHandler).Methods("POST")

	// login
	app.Router.HandleFunc("/login", app.LoginHandler).Methods("POST")
	// TODO app.Router.HandleFunc("/login/refresh", app.RefreshLoginHandler).Methods("POST")
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
