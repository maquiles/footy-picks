package main

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey string = os.Getenv("TOKEN")

func GenerateJWT(player Player) (Token, error) {
	sessionToken := Token{}
	signingKey := []byte(secretKey)

	token := jwt.New(jwt.SigningMethodHS256)
	expiration := time.Now().Add(time.Minute * 30)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = player.Email
	claims["id"] = player.ID
	claims["name"] = player.PlayerName
	claims["exp"] = expiration.Unix()

	tokenStr, err := token.SignedString(signingKey)
	if err != nil {
		fmt.Errorf("error generating token >> %s", err)
		return sessionToken, err
	}

	sessionToken.TokenString = tokenStr
	sessionToken.ExpireTime = expiration

	return sessionToken, nil
}
