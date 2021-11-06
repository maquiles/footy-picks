package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"golang.org/x/crypto/bcrypt"
)

// TODO
// not a real worl auth scenario, just POC
// need 2 token authentication / authorization
// 1 token with long expire timeout for authen
// 2nd token with short expire time for author

var users map[string][]byte = make(map[string][]byte)
var idxUsers int = 0

func getSecret() string {
	secret := os.Getenv("ACCESS_SCRET")
	if secret == "" {
		// TODO in reality should throw an error
		secret = "aSecretString"
	}

	return secret
}

func createToken(id string) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["username"] = id
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodES256, atClaims)
	secret := getSecret()

	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func checkToken(tokenString string) (*jwt.Token, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(getSecret()), nil
	})
	if err != nil || !token.Valid {
		return nil, false
	}
	return token, true
}

func isUsernameContextOk(username string, request *http.Request) bool {
	usernameCtx, ok := context.Get(request, "username").(string)
	if !ok {
		return false
	}
	if usernameCtx != username {
		return false
	}
	return true

	// USAGE
	// if !isUsernameContextOk(post.Username, r){
	// 	http.Error(w, "Cannot manage post for another user", http.StatusUnauthorized)
	// 	return
	// }
}

func createPlayerGetToken(player Player) (string, error) {
	_, found := users[player.ID]
	if found {
		return "", fmt.Errorf("Player already exists for ID: %s", player.ID)
	}

	value, _ := bcrypt.GenerateFromPassword([]byte(player.Password), bcrypt.DefaultCost)
	users[player.ID] = value
	return createToken(player.ID)
}

func getTokenUserPassword(player Player) (string, error) {
	pwHash, found := users[player.ID]
	if !found {
		return "", fmt.Errorf("Player not found for ID: %s", player.ID)
	}

	err := bcrypt.CompareHashAndPassword(pwHash, []byte(player.Password))
	if err != nil {
		return "", err
	}

	return createToken(player.ID)
}
