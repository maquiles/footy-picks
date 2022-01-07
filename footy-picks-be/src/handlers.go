package main

import (
	"encoding/json"
	"footypicks/fotmob"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// HEALTH
func HealthHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("received GET - /HEALTH request")
	json.NewEncoder(writer).Encode("literally just vibin")
}

// MATCH DATA
func CurrentRoundHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("received GET - /ROUND/CURRENT request")

	league := request.FormValue("league")
	tab := request.FormValue("tab")
	timezone := request.FormValue("timezone")

	response := fotmob.GetCurrentRoundMatches(league, tab, timezone)

	json.NewEncoder(writer).Encode(response)
}

// GAME
func (app App) NewGameHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("received POST - /GAME request")

	_, err := AuthenticateJWT(writer, request)
	if err != nil {
		return
	}

	var newGameBody NewGameBody
	err = json.NewDecoder(request.Body).Decode(&newGameBody)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	newGame, err := app.DBConn.CreateNewSurvivorGame(
		newGameBody.CreatorID,
		newGameBody.Name,
		newGameBody.Passcode,
		newGameBody.League)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(newGame)
}

func (app App) GetGameHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("received GET - /GAME/{game_id} request")

	playerID, err := AuthenticateJWT(writer, request)
	if err != nil {
		return
	}

	params := mux.Vars(request)
	gameIDParam := params["game_id"]
	gameID, err := strconv.Atoi(gameIDParam)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	game, err := app.GetSurvivorGameDetails(gameID, playerID)
	if err != nil {
		if err.Error() == "NotInvitedError" {
			http.Error(writer, err.Error(), http.StatusForbidden)
		} else {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(writer).Encode(game)
}

func (app *App) GamesForPlayerHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("received GET - /GAMES request")

	playerIDFromCookie, err := AuthenticateJWT(writer, request)
	if err != nil {
		return
	}

	params := mux.Vars(request)
	playerParam := params["player_id"]
	playerID, err := strconv.Atoi(playerParam)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if playerIDFromCookie != playerID {
		http.Error(writer, "PlayerImpersonationError", http.StatusForbidden)
		return
	}

	activeGames, err := app.GetAllGamesForPlayer(playerID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(activeGames)
}

func (app App) AddPlayerToGameHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("received POST - /GAME/PLAYER request")

	playerIDFromCookie, err := AuthenticateJWT(writer, request)
	if err != nil {
		return
	}

	params := mux.Vars(request)
	playerID, err := strconv.Atoi(params["player_id"])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if playerIDFromCookie != playerID {
		http.Error(writer, "PlayerImpersonationError", http.StatusForbidden)
		return
	}

	gameID, err := strconv.Atoi(params["game_id"])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	var body struct {
		Passcode string `json:"passcode"`
	}

	err = json.NewDecoder(request.Body).Decode(&body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	gameDeets := AddPlayerToGameBody{
		PlayerID: playerID,
		GameID:   gameID,
		Passcode: body.Passcode,
	}

	err = app.JoinGame(gameDeets)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(writer).Encode("success")
}

// PLAYER
func (app *App) CreateNewPlayerHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("received POST - /PLAYER request")

	var newPlayer NewPlayer
	err := json.NewDecoder(request.Body).Decode(&newPlayer)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	player, err := app.DBConn.CreateNewPlayer(newPlayer.Email, newPlayer.Name, newPlayer.Login)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("successfully created new player with id %d", player.ID)
	json.NewEncoder(writer).Encode(player)
}

func (app App) MakeGamePickHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("received POST - /PLAYER/{player_id}/GAME/{game_id}/PICK request")

	playerID, err := AuthenticateJWT(writer, request)
	if err != nil {
		return
	}

	var body SurvivorGamePickBody
	err = json.NewDecoder(request.Body).Decode(&body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = app.MakeSurvivorPick(body, playerID)
	if err != nil {
		if err.Error() == "CompletedGamePickError" ||
			err.Error() == "NoPicksForOngoingGameError" ||
			err.Error() == "KnockoutPickError" ||
			err.Error() == "PickAlreadyMadeError" {
			http.Error(writer, err.Error(), http.StatusForbidden)
			return
		}
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode("success")
}

// AUTH
func (app *App) LoginHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("received POST - /LOGIN request")

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

func (app App) RefreshLoginHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("received POST - /login/refresh")

	if err := RefreshJWT(writer, request); err != nil {
		return
	}

	json.NewEncoder(writer).Encode("success")
}
