package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

// for user auth endpoints
func checkTokenHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		bearerToken := strings.Split(header, " ")
		if len(bearerToken) != 2 {
			http.Error(w, "Cannot read token", http.StatusBadRequest)
			return
		}
		if bearerToken[0] != "Bearer" {
			http.Error(w, "Error in authorization token. it needs to be in form of 'Bearer <token>'", http.StatusBadRequest)
			return
		}
		token, ok := checkToken(bearerToken[1])
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			username, ok := claims["username"].(string)
			if !ok {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			if _, ok := users[username]; !ok {
				http.Error(w, "Unauthorized, user not exists", http.StatusUnauthorized)
				return
			}

			context.Set(r, "username", username)
		}
		next(w, r)
	}
}

func (app *App) HealthHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("received /health request")
	json.NewEncoder(writer).Encode("literally just vibin")
}

func (app *App) LoginHandler(writer http.ResponseWriter, request *http.Request) {
	var player Player

	if err := json.NewDecoder(request.Body).Decode(&player); err != nil {
		http.Error(writer, "cannot decode username and password", http.StatusBadRequest)
		return
	}

	token, err := getTokenUserPassword(player)
	if err != nil {
		http.Error(writer, "Failed to create token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(struct {
		Token string `json:"token"`
	}{token})
}

func (app *App) NewPlayerHandler(writer http.ResponseWriter, request *http.Request) {
	var player Player

	if err := json.NewDecoder(request.Body).Decode(&player); err != nil {
		http.Error(writer, "cannot decode username and password", http.StatusBadRequest)
		return
	}

	token, err := createPlayerGetToken(player)
	if err != nil {
		http.Error(writer, "Failed to create new user and get token", http.StatusInternalServerError)
	}

	json.NewEncoder(writer).Encode(struct {
		Token string `json:"token"`
	}{token})
}

func (app *App) RefreshTokenHandler(writer http.ResponseWriter, request *http.Request) {
	username, ok := context.Get(request, "username").(string)
	if !ok {
		http.Error(writer, "Failed checking username", http.StatusInternalServerError)
		return
	}

	token, err := createToken(username)
	if err != nil {
		http.Error(writer, "Error creating token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(struct {
		Token string `json:"token"`
	}{token})
}

func GetUserHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("receiving request to /user")

	vars := mux.Vars(request)
	user, ok := vars["username"]
	if !ok {
		http.Error(writer, "Unable to find player for username", http.StatusBadRequest)
		return
	}

	if _, ok := users[user]; !ok {
		http.Error(writer, "Player not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(writer).Encode(struct {
		Username    string `json:"username"`
		Description string `json:"description"`
	}{user, ""})
}
