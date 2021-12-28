package fotmob

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
