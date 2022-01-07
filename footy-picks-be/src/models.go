package main

import "footypicks/repo"

// for returning survivor game data to FE
type SurvivorGamePick struct {
	Round   int    `json:"round"`
	Pick    string `json:"pick"`
	Correct bool   `json:"correct"`
}

type SurvivorGameTableRow struct {
	Player string             `json:"player"`
	Rounds []SurvivorGamePick `json:"rounds"`
}

type SurvivorGameTable struct {
	ID     string                 `json:"id"`
	Name   string                 `json:"name"`
	League string                 `json:"league"`
	Rows   []SurvivorGameTableRow `json:"rows"`
}

// for returning player data to FE
type Player struct {
	ID         int    `json:"player_id"`
	Email      string `json:"email"`
	PlayerName string `json:"player_name"`
	Created    string `json:"created"`
	Games      []int  `json:"games"`
}

// for player requests
type NewPlayer struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Login string `json:"login"`
}

// for login requests
type Login struct {
	Email string `json:"email"`
	Login string `json:"login"`
}

// for game requests
type PlayerSurvivorGames struct {
	ActiveGames []repo.SurvivorGameEntity `json:"active_games"`
	PastGames   []repo.SurvivorGameEntity `json:"past_games"`
}

type NewGameBody struct {
	CreatorID int    `json:"creator"`
	Name      string `json:"name"`
	Passcode  string `json:"passcode"`
	League    string `json:"league"`
}

type AddPlayerToGameBody struct {
	PlayerID int    `json:"player_id"`
	GameID   int    `json:"game_id"`
	Passcode string `json:"passcode"`
}

type SurvivorGamePickBody struct {
	Round int    `json:"game_round"`
	Pick  string `json:"pick"`
	Game  int    `json:"game_id"`
}
