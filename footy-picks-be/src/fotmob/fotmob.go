package fotmob

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// PRIVATE
func getFotmobLeague(league string, tab string, timezone string) LeagueResponse {
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
	return leagueResponse
}

func getFirstRoundWithUnplatedMatch(leagueResponse LeagueResponse) int {
	return leagueResponse.MatchesTab.FirstUnplayedMatch.FirstRoundWithUnplayedMatch
}

func getAllLeagueMatches(leagueResponse LeagueResponse) []Match {
	return leagueResponse.MatchesTab.Data.AllMatches
}

func sortAllLeagueMatchesByRound(matches []Match) map[int][]Match {
	rounds := map[int][]Match{}

	for _, match := range matches {
		round := rounds[match.Round]

		if round != nil {
			round = append(round, match)
			rounds[match.Round] = round
		} else {
			rounds[match.Round] = []Match{match}
		}

	}

	return rounds
}

// PUBLIC
func GetCurrentRoundMatches(league string, tab string, timezone string) []Match {
	leagueResponse := getFotmobLeague(league, tab, timezone)
	currentRound := getFirstRoundWithUnplatedMatch(leagueResponse)
	matches := getAllLeagueMatches(leagueResponse)
	matchesByRound := sortAllLeagueMatchesByRound(matches)

	return matchesByRound[currentRound]
}
