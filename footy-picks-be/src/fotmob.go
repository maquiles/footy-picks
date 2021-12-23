package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// FOTMOB URLS
var FOTMOB_BASE_URL string = "https://www.fotmob.com"
var MATCHES_URL string = FOTMOB_BASE_URL + "/matches?"
var LEAGUES_URL string = FOTMOB_BASE_URL + "/leagues?"
var TEAMS_URL string = FOTMOB_BASE_URL + "/teams?"
var PLAYERS_URL string = FOTMOB_BASE_URL + "/platerData?"
var MATCH_DETAILS_URL string = FOTMOB_BASE_URL + "/matchDetails?"
var SEARCH_URL = FOTMOB_BASE_URL + "/searchapi/"

//FOTMOB LEAGUE IDS
var LEAGUE_IDS map[string]string = map[string]string{
	"UCL":  "42",
	"EPL":  "47",
	"EURO": "50",
}

//FOTMOB MODELS
type Team struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
}

type MatchFinishedReason struct {
	Long  string `json:"long"`
	Short string `json:"short"`
}

type MatchStatus struct {
	Cancelled         bool                `json:"cancelled"`
	Finished          bool                `json:"finished"`
	Reason            MatchFinishedReason `json:"reason"`
	ScoreStr          string              `json:"scoreStr"`
	StartDateStr      string              `json:"startDateStr"`
	StartDateStrShort string              `json:"startDateStrShort"`
	Started           bool                `json:"started"`
}

type Match struct {
	Away      Team        `json:"away"`
	Home      Team        `json:"home"`
	ID        string      `json:"id"`
	MonthKey  string      `json:"monthKey"`
	PageURL   string      `json:"pageUrl"`
	Round     int         `json:"round"`
	RoundName int         `json:"roundName"`
	Status    MatchStatus `json:"status"`
}

type LeagueDetails struct {
	Country   string      `json:"country"`
	FaqJSONLD interface{} `json:"faqJSONLD"`
	ID        int         `json:"id"`
	Name      string      `json:"name"`
	ShortName string      `json:"shortName"`
	Type      string      `json:"typw"`
}

type FirstUnplayedMatch struct {
	FirstRoundWithUnplayedMatch int `json:"firstRoundWithUnplayedMatch"`
	FirstUnplayedMatchIndex     int `json:"firstUnplayedMatchIndex"`
}

type MatchesTabData struct {
	AllMatches             []Match     `json:"allMatches"`
	MatchesCombinedByRound interface{} `json:"matchesCombinedByRound"`
}

type MatchesTab struct {
	Data               MatchesTabData     `json:"data"`
	FirstUnplayedMatch FirstUnplayedMatch `json:"firstUnplayedMatch"`
	Seostr             string             `json:"seostr"`
	Tab                string             `json:"tab"`
}

type LeagueResponse struct {
	QAData     interface{}   `json:"QAData"`
	Action     interface{}   `json:"action"`
	Datasets   interface{}   `json:"datasets"`
	Details    LeagueDetails `json:"details"`
	MatchesTab MatchesTab    `json:"matchesTab"`
	Seostr     string        `json:"seostr"`
	Tab        string        `json:"tab"`
	Tabs       []string      `json:"tabs"`
}

func GetLeagueMatches(league string, tab string, timezone string) []Match {
	leaugeID := LEAGUE_IDS[league]
	url := fmt.Sprintf("%stimezone=%s&id=%s&tab=%s", LEAGUES_URL, timezone, leaugeID, tab)
	log.Printf("sending fotmob request %s", url)

	response, err := http.Get(url)
	if err != nil {
		log.Println("ERROR in getLeague >>", err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("error parsing response body from getLeague >>", err)
	}

	var leagueResponse LeagueResponse
	json.Unmarshal(body, &leagueResponse)

	matches := leagueResponse.MatchesTab.Data.AllMatches

	return matches
}
