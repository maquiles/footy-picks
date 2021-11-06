package main

type Player struct {
	ID       string   `json:"id"`
	Password string   `json:"password"`
	Name     string   `json:"name"`
	Games    []string `json:"games"`
}

type Game struct {
	ID        string
	Name      string
	Players   []string
	StartDate string
	EndDate   string
	Status    string
	Type      string
	League    string
}

type Round struct {
	ID        string
	Game      string
	StartDate string
	EndDate   string
	Picks     []string
}

type Pick struct {
	ID     string
	Round  string
	Game   string
	Player string
	Pick   string
}
