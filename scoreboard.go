package main

type ScoreboardStruct struct {
	Season Season      `json:"season"`
	Day    Day         `json:"day"`
	Games  []GameScore `json:"events"`
}

type Season struct {
	Type int `json:"type"`
	Year int `json:"year"`
}

type Day struct {
	Date string `json:"date"`
}

type GameScore struct {
	ID           string        `json:"id"`
	Date         string        `json:"date"`
	Name         string        `json:"shortName"`
	Competitions []Competition `json:"competitions"`
	Links        []Link        `json:"links"`
	Status       GameStatus    `json:"status"`
}

type Competition struct {
	Competitors []Competitor `json:"competitors"`
	Broadcasts  []Broadcast  `json:"broadcasts"`
	Headlines   []Headline   `json:"headlines"`
}

type Competitor struct {
	Team       TeamStruct  `json:"team"`
	Winner     bool        `json:"winner,omitempty"`
	HomeAway   string      `json:"homeAway"`
	Points     string      `json:"score,omitempty"`
	LineScores []LineScore `json:"linescores,omitempty"`
	Statistics []Statistic `json:"statistics"`
	Records    []Record    `json:"records"`
}

type LineScore struct {
	Value float64 `json:"value"`
}

type Statistic struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	DisplayValue string `json:"displayValue"`
}

type Record struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	Type         string `json:"type"`
	Summary      string `json:"summary"`
}

type Broadcast struct {
	Market string   `json:"market"`
	Names  []string `json:"names"`
}

type Headline struct {
	Description string  `json:"description,omitempty"`
	Type        string  `json:"type,omitempty"`
	Title       string  `json:"shortLinkText,omitempty"`
	Video       []Video `json:"video,omitempty"`
}

type Video struct {
	Title string     `json:"headline,omitempty"`
	Links VideoLinks `json:"links,omitempty"`
}

type VideoLinks struct {
	Source VideoLinksSource `json:"source,omitempty"`
}

type VideoLinksSource struct {
	HD Link `json:"HD"`
}

type GameStatus struct {
	Clock  string         `json:"displayClock"`
	Period int            `json:"period"`
	Type   GameStatusType `json:"type"`
}

type GameStatusType struct {
	Detail string `json:"detail"`
}
