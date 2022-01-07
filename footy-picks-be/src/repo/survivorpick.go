package repo

import (
	"log"
)

type SurvivorPickEntity struct {
	Round   int    `json:"game_round"`
	Pick    string `json:"pick"`
	Correct int    `json:"correct"`
	Game    int    `json:"survivor_game"`
	Player  int    `json:"player"`
}

func (repo Repo) GetSurvivorGamePicksForPlayer(game int, player int) ([]SurvivorPickEntity, error) {
	query := `SELECT * FROM survivor_pick WHERE survivor_game = $1 AND player = $2 ORDER BY game_round ASC`
	rows, err := repo.DBConn.Query(query, game, player)
	if err != nil {
		log.Printf("error getting picks for player = %d and game = %d", player, game)
		return []SurvivorPickEntity{}, err
	}
	defer rows.Close()

	var picks []SurvivorPickEntity
	for rows.Next() {
		var pick SurvivorPickEntity
		err = rows.Scan(
			&pick.Round,
			&pick.Pick,
			&pick.Correct,
			&pick.Game,
			&pick.Player,
		)
		if err != nil {
			log.Printf("error scanning picks for player = %d and game = %d", player, game)
			return []SurvivorPickEntity{}, err
		}

		picks = append(picks, pick)
	}

	if err := rows.Err(); err != nil {
		log.Printf("error while iterating through picks for player = %d and game = %d", player, game)
		return []SurvivorPickEntity{}, err
	}

	return picks, nil
}

func (repo Repo) AddSurvivorGamePick(round int, pick string, game int, player int) error {
	query := `
	INSERT INTO survivor_pick
		(game_round, pick, correct, game, player)
		VALUES ($1, $2, 0, $3, $4)`

	_, err := repo.DBConn.Exec(query, round, pick, game, player)

	return err
}
