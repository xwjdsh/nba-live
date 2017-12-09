package main

import (
	"github.com/fatih/color"
)

const (
	SUCCESS = "\u2714"
	ERROR   = "\u2716"
)

func main() {
	games, err := getGames()
	if err != nil {
		color.Red("%s 获取比赛数据错误:%s", ERROR, err.Error())
		return
	}

	if len(games) == 0 {
		color.Green("%s 暂无比赛数据", SUCCESS)
		return
	}

	//games := []*Game{
	//&Game{
	//HomeTeam:   "骑士",
	//VisitTeam:  "勇士",
	//HomeScore:  "100",
	//VisitScore: "90",
	//PeriodCn:   "第4节 01:30",
	//},
	//&Game{
	//HomeTeam:   "凯尔特人",
	//VisitTeam:  "公牛",
	//HomeScore:  "80",
	//VisitScore: "88",
	//PeriodCn:   "第3节 03:30",
	//},
	//}
	i, err := newSelect(games)
	if err != nil {
		color.Red("%s 选择比赛错误:%s", ERROR, err.Error())
		return
	}
	newUI(games[i])
}
