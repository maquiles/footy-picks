package main

import "time"

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

type NewPlayerGame struct {
	PlayerID int `json:"player_id"`
	GameID   int `json:"game_id"`
}

// for login requests
type Login struct {
	Email string `json:"email"`
	Login string `json:"login"`
}

type Token struct {
	TokenString string    `json:"token_string"`
	ExpireTime  time.Time `json:"expire_time"`
}
