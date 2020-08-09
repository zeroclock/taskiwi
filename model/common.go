package model

type Tags []string

type ClockData struct {
	Task     string   `json:"task"`
	Parents  string   `json:"parents"`
	Category string   `json:"category"`
	Start    string   `json:"start"`
	End      string   `json:"end"`
	Effort   string   `json:"effort"`
	Ishabit  string   `json:"ishabit"`
	Tags     []string `json:"tags"`
}
