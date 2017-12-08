package main

import (
	"log"
	"os"

	"github.com/fatih/color"
)

func main() {
	games, err := getGames()
	if err != nil {
		color.Red("\u2716 获取比赛数据错误:", err.Error())
		return
	}

	if len(games) == 0 {
		color.Green("\u2714 暂无比赛数据")
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
		os.Exit(1)
	}
	log.Println(i)
	newUI()
}
