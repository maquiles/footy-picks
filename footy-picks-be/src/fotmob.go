package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var FOTMOB_BASE_URL string = "https://www.fotmob.com"
var MATCHES_URL string = FOTMOB_BASE_URL + "/matches?"
var LEAGUES_URL string = FOTMOB_BASE_URL + "/leagues?"
var TEAMS_URL string = FOTMOB_BASE_URL + "/teams?"
var PLAYERS_URL string = FOTMOB_BASE_URL + "/platerData?"
var MATCH_DETAILS_URL string = FOTMOB_BASE_URL + "/matchDetails?"
var SEARCH_URL = FOTMOB_BASE_URL + "/searchapi/"

func getLeague(id string, tab string, typ string, timeZone string) string {
	url := fmt.Sprintf("%sid=%s&tab=%s&type=%s&timeZone=%s", LEAGUES_URL, id, tab, typ, timeZone)
	log.Printf("sending fotmob request %s", url)

	response, err := http.Get(url)
	if err != nil {
		log.Println("ERROR in getLeague >>", err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("error parsing response body from getLeague >>", err)
	}

	return string(body)
}
