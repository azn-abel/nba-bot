package main

type TeamResponse struct {
	Team TeamStruct `json:"team"`
}

type TeamStruct struct {
	ID        string     `json:"id"`
	Name      string     `json:"displayName"`
	Color     string     `json:"color"`
	Logos     []Link     `json:"logos,omitempty"`
	Record    TeamRecord `json:"record,omitempty"`
	Links     []Link     `json:"links"`
	Logo      string     `json:"logo,omitempty"`
	NextEvent []Game     `json:"nextEvent,omitempty"`
}

type Link struct {
	URL string `json:"href"`
}

type TeamRecord struct {
	Items []TeamRecordItem `json:"items"`
}

type TeamRecordItem struct {
	Description string
	Summary     string
	Stats       []TeamStat `json:"stats"`
}

type TeamStat struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type Game struct {
	Name  string `json:"name"`
	Date  string `json:"date"`
	Links []Link `json:"links"`
}

var TeamIDByName map[string]int = map[string]int{
	"hawks":        1,
	"celtics":      2,
	"pelicans":     3,
	"bulls":        4,
	"cavaliers":    5,
	"mavericks":    6,
	"nuggets":      7,
	"pistons":      8,
	"warriors":     9,
	"rockets":      10,
	"pacers":       11,
	"clippers":     12,
	"lakers":       13,
	"heat":         14,
	"bucks":        15,
	"timberwolves": 16,
	"nets":         17,
	"knicks":       18,
	"magic":        19,
	"sixers":       20,
	"suns":         21,
	"blazers":      22,
	"kings":        23,
	"spurs":        24,
	"thunder":      25,
	"jazz":         26,
	"wizards":      27,
	"raptors":      28,
	"grizzlies":    29,
	"hornets":      30,
}
