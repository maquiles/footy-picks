package main

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
