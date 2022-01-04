package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey string = os.Getenv("TOKEN")

type Claims struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	jwt.StandardClaims
}

func GenerateJWT(player Player, writer http.ResponseWriter) error {
	expiration := time.Now().Add(time.Minute * 30)
	claims := &Claims{
		ID:    player.ID,
		Email: player.Email,
		Name:  player.PlayerName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(secretKey))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return err
	}

	http.SetCookie(writer, &http.Cookie{
		Name:    "token",
		Value:   tokenStr,
		Expires: expiration,
	})

	return nil
}

func AuthenticateJWT(writer http.ResponseWriter, request *http.Request) (int, error) {
	cookie, err := request.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(writer, err.Error(), http.StatusUnauthorized)
			return -1, err
		}
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return -1, err
	}

	tokenString := cookie.Value
	var claims Claims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			http.Error(writer, err.Error(), http.StatusUnauthorized)
			return -1, err
		}
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return -1, err
	}

	if !token.Valid {
		http.Error(writer, err.Error(), http.StatusUnauthorized)
		return -1, fmt.Errorf("InvalidTokenError")
	}

	return claims.ID, nil
}

func RefreshJWT(writer http.ResponseWriter, request *http.Request) error {
	cookie, err := request.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(writer, err.Error(), http.StatusUnauthorized)
			return err
		}
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return err
	}

	tokenString := cookie.Value
	claims := Claims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			http.Error(writer, err.Error(), http.StatusUnauthorized)
			return err
		}
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return err
	}

	if !token.Valid {
		http.Error(writer, err.Error(), http.StatusUnauthorized)
		return fmt.Errorf("InvalidTokenError")
	}

	if time.Until(time.Unix(claims.ExpiresAt, 0)) > 30*time.Second {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return fmt.Errorf("RefreshTokenError")
	}

	expiration := time.Now().Add(time.Minute * 30)
	claims.ExpiresAt = expiration.Unix()
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := newToken.SignedString([]byte(secretKey))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return err
	}

	http.SetCookie(writer, &http.Cookie{
		Name:    "token",
		Value:   tokenStr,
		Expires: expiration,
	})

	return nil
}
