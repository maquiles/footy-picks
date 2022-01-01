package main

import (
	"encoding/json"
	"footypicks/fotmob"
	"log"
	"net/http"
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
	user := request.FormValue("user")

	// TODO
	// activeGames := GetActiveGamesForUser(user)

	// json.NewEncoder(writer).Encode(activeGames)
}

// PLAYER CRUD
func (app *App) CreateNewPlayerHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("received POST - /PLAYER request")

	var newPlayer NewPlayer
	err := json.NewDecoder(request.Body).Decode(&newPlayer)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
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
	}

	err = app.DBConn.UpdatePlayerGames(newPlayerGame.PlayerID, newPlayerGame.GameID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
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
	}

	player, token, err := app.SignIn(login)
	if err != nil {
		if player.ID == 0 {
			http.Error(writer, err.Error(), http.StatusNotFound)
		} else if player.ID == -1 {
			http.Error(writer, err.Error(), http.StatusUnauthorized)
		} else {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	}

	http.SetCookie(writer, &http.Cookie{
		Name:    "token",
		Value:   token.TokenString,
		Expires: token.ExpireTime,
	})

	json.NewEncoder(writer).Encode(player)
}
