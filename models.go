package main

type Game struct {
	ID         string      `json:"id"`
	Sdate      string      `json:"sdate"`
	Time       string      `json:"time"`
	URL        string      `json:"url"`
	Type       string      `json:"type"`
	Start      string      `json:"start"`
	HomeTeam   string      `json:"home_team"`
	VisitTeam  string      `json:"visit_team"`
	HomeScore  interface{} `json:"home_score"`
	VisitScore interface{} `json:"visit_score"`
	PeriodCn   string      `json:"period_cn"`
	From       string      `json:"from"`
	Code       string      `json:"code"`
	Update     string      `json:"update"`
	//BigScore1     string `json:"big_score_1"`
	//BigScore2     string `json:"big_score_2"`
}

type GameInfo struct {
	ID             string   `json:"id"`
	Sdate          string   `json:"sdate"`
	Time           string   `json:"time"`
	URL            string   `json:"url"`
	Type           string   `json:"type"`
	Start          string   `json:"start"`
	HomeTeam       string   `json:"home_team"`
	VisitTeam      string   `json:"visit_team"`
	HomeScore      string   `json:"home_score"`
	VisitScore     string   `json:"visit_score"`
	Team1Scores    []string `json:"team1_scores"`
	Team2Scores    []string `json:"team2_scores"`
	PeriodCn       string   `json:"period_cn"`
	From           string   `json:"from"`
	Code           string   `json:"code"`
	Update         string   `json:"update"`
	FullTimeouts1  string   `json:"full_timeouts_1"`
	FullTimeouts2  string   `json:"full_timeouts_2"`
	ShortTimeouts1 string   `json:"short_timeouts_1"`
	ShortTimeouts2 string   `json:"short_timeouts_2"`
	TeamFouls1     string   `json:"team_fouls_1"`
	TeamFouls2     string   `json:"team_fouls_2"`
	BigScore1      string   `json:"big_score_1"`
	BigScore2      string   `json:"big_score_2"`
}

type GameLiveRecord struct {
	LiveID      string `json:"live_id"`
	RoomID      string `json:"room_id"`
	LiveSid     string `json:"live_sid"`
	MatchID     string `json:"match_id"`
	SaishiID    string `json:"saishi_id"`
	ScoreStatus string `json:"score_status"`
	UserID      string `json:"user_id"`
	UserChn     string `json:"user_chn"`
	GuessID     string `json:"guess_id"`
	GuessText   string `json:"guess_text"`
	GuessData   string `json:"guess_data"`
	LiveText    string `json:"live_text"`
	LivePid     string `json:"live_pid"`
	TextColor   string `json:"text_color"`
	TextURL     string `json:"text_url"`
	ImgURL      string `json:"img_url"`
	LiveTime    string `json:"live_time"`
	HomeScore   string `json:"home_score"`
	VisitScore  string `json:"visit_score"`
	SnTeamName  string `json:"sn_team_name"`
	LivePtime   string `json:"live_ptime"`
	PeriodScore string `json:"period_score"`
	PidText     string `json:"pid_text"`
	ImgMode     string `json:"img_mode"`
}

type GameInfoAndLive struct {
	Info        *GameInfo
	LiveRecords []*GameLiveRecord
}
