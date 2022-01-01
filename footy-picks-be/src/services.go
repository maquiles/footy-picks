package main

import (
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

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
