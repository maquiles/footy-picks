package repo

import (
	"log"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type PlayerEntity struct {
	ID          int           `json:"player_id"`
	Email       string        `json:"email"`
	PlayerName  string        `json:"player_name"`
	PlayerLogin string        `json:"player_login"`
	Created     string        `json:"created"`
	Games       pq.Int64Array `json:"games"`
}

func (repo Repo) CreateNewPlayer(email string, name string, login string) (PlayerEntity, error) {
	// TODO: make sure that email is valid. currectly will work for any string

	created := getCurrentTimestamp()
	hash, _ := bcrypt.GenerateFromPassword([]byte(login), bcrypt.DefaultCost)

	query := `
		INSERT INTO player (email, player_name, player_login, created, games)
			VALUES ($1, $2, $3, $4, '{}')
			RETURNING player_id`

	var playerID int
	err := repo.DBConn.QueryRow(query, email, name, string(hash), created).Scan(&playerID)
	if err != nil {
		log.Println("error creating new player >>", err)
		return PlayerEntity{ID: -1}, err
	}

	return PlayerEntity{
		ID:          playerID,
		Email:       email,
		PlayerName:  name,
		PlayerLogin: login,
		Created:     created,
		Games:       pq.Int64Array{},
	}, nil
}

func (repo Repo) UpdatePlayerGames(playerID int, gameID int) error {
	query := `
		UPDATE player SET games = array_append(games, $1) 
			WHERE player_id = $2`

	_, err := repo.DBConn.Exec(query, gameID, playerID)

	return err
}

func (repo Repo) GetPlayer(email string) (PlayerEntity, error) {
	query := `SELECT * FROM player WHERE email = $1`

	var player PlayerEntity
	err := repo.DBConn.QueryRow(query, email).Scan(
		&player.ID,
		&player.Email,
		&player.PlayerName,
		&player.PlayerLogin,
		&player.Created,
		&player.Games)

	if err != nil {
		log.Printf("error getting player with email %s >> %s", email, err)
		return PlayerEntity{ID: 0}, err
	}

	return player, nil
}

func (repo Repo) GetPlayerNameByID(playerID int) (string, error) {
	query := `SELECT name from player where player_id = $1`
	var name string
	err := repo.DBConn.QueryRow(query, playerID).Scan(&name)
	if err != nil {
		log.Printf("error getting name for player = %d", playerID)
		return "", err
	}

	return name, nil
}
