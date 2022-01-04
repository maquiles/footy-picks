package main

import (
	"fmt"
	"footypicks/repo"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// AUTH
func (app App) SignIn(login Login, writer http.ResponseWriter) (Player, error) {
	playerEntity, err := app.DBConn.GetPlayer(login.Email)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return Player{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(playerEntity.PlayerLogin), []byte(login.Login))
	if err != nil {
		log.Printf("error incorrect creds for player with email %s >> %s", login.Email, err)
		http.Error(writer, err.Error(), http.StatusUnauthorized)
		return Player{}, err
	}

	player := Player{
		ID:         playerEntity.ID,
		Email:      playerEntity.Email,
		PlayerName: playerEntity.PlayerName,
		Created:    playerEntity.Created,
		Games:      playerEntity.Games,
	}

	err = GenerateJWT(player, writer)
	if err != nil {
		return Player{}, err
	}

	return player, nil
}

// GAME
func (app App) GetAllGamesForPlayer(id int) (PlayerSurvivorGames, error) {
	games, err := app.DBConn.GetGamesForPlayer(id)
	if err != nil {
		return PlayerSurvivorGames{}, err
	}

	var gamesSorted PlayerSurvivorGames
	for _, game := range games {
		if game.Ongoing {
			gamesSorted.ActiveGames = append(gamesSorted.ActiveGames, game)
		} else {
			gamesSorted.PastGames = append(gamesSorted.PastGames, game)
		}
	}

	return gamesSorted, nil
}

func (app App) GetSurvivorGameDetails(gameID int, playerID int) (repo.SurvivorGameEntity, error) {
	game, err := app.DBConn.GetSurvivorGameByID(gameID)
	if err != nil {
		return repo.SurvivorGameEntity{}, err
	}

	if !contains(game.Players, playerID) {
		return repo.SurvivorGameEntity{}, fmt.Errorf("NotInvitedError")
	}

	return game, nil
}

func (app App) JoinGame(body AddPlayerToGameBody) error {
	var gameID int
	var err error

	if body.Passcode == "" {
		gameID = body.GameID
	} else {
		gameID, err = app.DBConn.GetGameIDByPasscode(body.Passcode)
		if err != nil {
			log.Printf("error >> no game found with passcode = %s", body.Passcode)
			return err
		}
	}

	err = app.DBConn.AddPlayerToSurvivorGame(body.PlayerID, gameID)
	if err != nil {
		log.Printf("error adding player to survivor game >> player_id = %d; game_id = %d", body.PlayerID, body.GameID)
		return err
	}

	err = app.DBConn.UpdatePlayerGames(body.PlayerID, gameID)
	if err != nil {
		log.Printf("error udpating player games >> player_id = %d; game_id = %d", body.PlayerID, body.GameID)
		return err
	}

	return nil
}
