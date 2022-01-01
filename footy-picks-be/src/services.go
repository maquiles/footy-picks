package main

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (app App) SignIn(login Login) (Player, Token, error) {
	playerEntity, err := app.DBConn.GetPlayer(login.Email)
	if err != nil {
		return Player{ID: 0}, Token{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(playerEntity.PlayerLogin), []byte(login.Login))
	if err != nil {
		log.Printf("error incorrect creds for player with email %s >> %s", login.Email, err)
		return Player{ID: -1}, Token{}, err
	}

	player := Player{
		ID:         playerEntity.ID,
		Email:      playerEntity.Email,
		PlayerName: playerEntity.PlayerName,
		Created:    playerEntity.Created,
		Games:      playerEntity.Games,
	}

	token, err := GenerateJWT(player)
	if err != nil {
		return Player{ID: -2}, token, err
	}

	return player, token, nil
}
