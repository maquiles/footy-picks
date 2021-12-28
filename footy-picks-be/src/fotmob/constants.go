package fotmob

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
