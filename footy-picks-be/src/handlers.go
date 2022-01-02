package main

import (
	"encoding/json"
	"footypicks/fotmob"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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

func (app *App) GamesForUserHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("received GET - /GAMES request")

	if err := AuthenticateJWT(writer, request); err != nil {
		return
	}

	params := mux.Vars(request)
	playerParam := params["player_id"]
	playerID, err := strconv.Atoi(playerParam)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	activeGames, err := app.GetAllGamesForPlayer(playerID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(activeGames)
}

// PLAYER CRUD
func (app *App) CreateNewPlayerHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("received POST - /PLAYER request")

	var newPlayer NewPlayer
	err := json.NewDecoder(request.Body).Decode(&newPlayer)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	player, err := app.DBConn.CreateNewPlayer(newPlayer.Email, newPlayer.Name, newPlayer.Login)
	// TODO error handling
	log.Printf("successfully created new player with id %d", player.ID)

	json.NewEncoder(writer).Encode(player)
}

func (app *App) AddPlayerGameHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("received POST - /PLAYER/GAMES request")

	var newPlayerGame NewPlayerGame
	err := json.NewDecoder(request.Body).Decode(&newPlayerGame)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = app.DBConn.UpdatePlayerGames(newPlayerGame.PlayerID, newPlayerGame.GameID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("successfully add game_id %d to player_id %d", newPlayerGame.GameID, newPlayerGame.PlayerID)

	json.NewEncoder(writer).Encode(1)
}

// AUTH
func (app *App) LoginHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("received POST = /LOGIN request")

	var login Login
	err := json.NewDecoder(request.Body).Decode(&login)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	player, err := app.SignIn(login, writer)
	if err != nil {
		return
	}

	json.NewEncoder(writer).Encode(player)
}
