package repo

import (
	"log"
)

type PlayerEntity struct {
	ID          int    `json:"player_id"`
	Email       string `json:"email"`
	PlayerName  string `json:"player_name"`
	PlayerLogin string `json:"player_login"`
	Created     string `json:"created"`
	Games       []int  `json:"games"`
}

func (repo Repo) CreateNewPlayer(email string, name string, login string) (PlayerEntity, error) {
	created := getCurrentTimestamp()
	query := `
		INSERT INTO player (email, player_name, player_login, created, games)
			VALUES ($1, $2, $3, $4, '{}')
			RETURNING player_id`

	var playerID int
	err := repo.DBConn.QueryRow(query, email, name, login, created).Scan(&playerID)
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
		Games:       []int{},
	}, nil
}

func (repo Repo) UpdatePlayerGames(playerID int, gameID int) error {
	query := `
		UPDATE player SET games = array_append(games, $1) 
			WHERE player_id = $2`

	_, err := repo.DBConn.Exec(query, gameID, playerID)

	return err
}

func (repo Repo) GetPlayerByEmail(email string) (PlayerEntity, error) {
	query := `
		SELECT * FROM player WHERE email = $1`

	var player PlayerEntity
	err := repo.DBConn.QueryRow(query, email).Scan(
		&player.ID,
		&player.Email,
		&player.PlayerName,
		&player.PlayerLogin,
		&player.Created,
		&player.Games)

	if err != nil {
		log.Print("error getting player with email %s >> %s", email, err)
		return PlayerEntity{ID: -1}, err
	}

	return player, nil
}
