package repo

import (
	"footypicks/fotmob"
	"log"
)

type SurvivorGameEntity struct {
	ID             int    `json:"game_id"`
	GameName       string `json:"game_name"`
	Passcode       string `json:"passcode"`
	LeagueID       int    `json:"league_id"`
	League         string `json:"league"`
	Ongoing        bool   `json:"ongoing"`
	BeginningRound int    `json:"beginning_round"`
	Created        string `json:"created"`
	Creator        int    `json:"creator"`
	Players        []int  `json:"players"`
}

func (repo Repo) CreateNewSurvivorGame(playerID int, gameName string, passcode string, league string) (SurvivorGameEntity, error) {
	leagueID := fotmob.LEAGUE_IDS[league]
	beginningRound := fotmob.GetNextRound(league, "matches", "America/New_York")
	created := getCurrentTimestamp()
	players := []int{playerID}

	query := `
		INSERT INTO survivor_game 
			(game_name, passcode, league_id, league, ongoing, beginning_round, created, creator, players)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
			RETURNING game_id`

	var gameID int
	err := repo.DBConn.QueryRow(query, gameName, passcode, leagueID, league, true, beginningRound, created, playerID, players).Scan(&gameID)
	if err != nil {
		log.Println("error creating new game >>", err)
		return SurvivorGameEntity{}, err
	}

	return SurvivorGameEntity{
		ID:             gameID,
		GameName:       gameName,
		Passcode:       passcode,
		LeagueID:       leagueID,
		League:         league,
		Ongoing:        true,
		BeginningRound: beginningRound,
		Created:        created,
		Creator:        playerID,
		Players:        players,
	}, nil
}

func (repo Repo) AddPlayerToSurvivorGame(playerID int, gameID int) error {
	query := `
		UPDATE survivor_game SET players = array_append(player, $1)
			WHERE game_id = $2`

	_, err := repo.DBConn.Exec(query, playerID, gameID)

	return err
}

func (repo Repo) GetSurvivorGameByID(gameID int) (SurvivorGameEntity, error) {
	query := `SELECT * FROM survivor_game WHERE game_id = $1`

	var game SurvivorGameEntity
	err := repo.DBConn.QueryRow(query, gameID).Scan(
		game.ID,
		game.GameName,
		game.Passcode,
		game.LeagueID,
		game.League,
		game.Ongoing,
		game.BeginningRound,
		game.Created,
		game.Creator,
		game.Players)

	if err != nil {
		log.Printf("error getting game with id = %d >> %s", gameID, err)
		return SurvivorGameEntity{ID: -1}, err
	}

	return game, nil
}

func (repo Repo) GetGamesForPlayer(playerID int) ([]SurvivorGameEntity, error) {
	query := `SELECT * FROM survivor_game WHERE game_id = ANY(unnest(SELECT games FROM player WHERE player_id = $1));`

	rows, err := repo.DBConn.Query(query, playerID)
	if err != nil {
		log.Printf("error getting games for player with id = %d >> %s", playerID, err)
		return []SurvivorGameEntity{}, err
	}
	defer rows.Close()

	var games []SurvivorGameEntity
	for rows.Next() {
		var game SurvivorGameEntity
		err = rows.Scan(
			&game.ID,
			&game.GameName,
			&game.Passcode,
			&game.LeagueID,
			&game.League,
			&game.Ongoing,
			&game.BeginningRound,
			&game.Created,
			&game.Creator,
			&game.Players,
		)
		if err != nil {
			log.Printf("error scanning games for player with id = %d >> %s", playerID, err)
		}

		games = append(games, game)
	}

	if err := rows.Err(); err != nil {
		log.Printf("error while iterating through games for player with id = %d >> %s", playerID, err)
	}

	return games, nil
}

func (repo Repo) GetGameIDByPasscode(passcode string) (int, error) {
	query := `SELECT game_id FROM survivor_game WHERE passcode = $1`

	var id int
	err := repo.DBConn.QueryRow(query, passcode).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}
