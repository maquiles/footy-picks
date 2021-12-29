package repo

import (
	"log"
	"time"
)

type PlayerEntity struct {
	ID          int    `json:"player_id"`
	Email       string `json:"email"`
	PlayerName  string `json:"player_name"`
	PlayerLogin string `json:"player_login"`
	Created     string `json:"created"`
	Games       []int  `json:"games"`
}

func (repo Repo) CreateNewPlayer(email string, name string, login string) PlayerEntity {
	created := time.Now().Format("2006-01-02")
	query := `
		INSERT INTO player (email, player_name, player_login, created, games)
			VALUES ($1, $2, $3, $4, '{}')
			RETURNING player_id`

	var playerID int
	err := repo.DBConn.QueryRow(query, email, name, login, created).Scan(&playerID)
	if err != nil {
		log.Println("error creating new player >>", err)
		return PlayerEntity{ID: -1}
	}

	return PlayerEntity{
		ID:          playerID,
		Email:       email,
		PlayerName:  name,
		PlayerLogin: login,
		Created:     created,
		Games:       []int{},
	}
}

func (repo Repo) UpdatePlayerGames(playerID int, gameID int) error {
	query := `
		UPDATE player SET games = array_append(games, $1) 
			WHERE player_id = $2`

	_, err := repo.DBConn.Exec(query, gameID, playerID)

	return err
}
