package main

type Game struct {
	ID         string `json:"id"`
	Sdate      string `json:"sdate"`
	Time       string `json:"time"`
	URL        string `json:"url"`
	Type       string `json:"type"`
	Start      string `json:"start"`
	HomeTeam   string `json:"home_team"`
	VisitTeam  string `json:"visit_team"`
	HomeScore  string `json:"home_score"`
	VisitScore string `json:"visit_score"`
	PeriodCn   string `json:"period_cn"`
	From       string `json:"from"`
	Code       string `json:"code"`
	Update     string `json:"update"`
	BigScore1  string `json:"big_score_1"`
	BigScore2  string `json:"big_score_2"`
}
