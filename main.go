package main

import (
	"fmt"
)

const (
	IconGood = "✔"
	IconBad  = "✗"
)

func main() {
	games, err := getGames()
	if err != nil {
		fmt.Printf("%s %s\n", Red(IconBad), "获取比赛列表失败")
		return
	}

	if len(games) == 0 {
		fmt.Printf("%s %s\n", Green(IconGood), "暂无比赛数据")
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
	//&Game{
	//HomeTeam:   "火箭",
	//VisitTeam:  "活塞",
	//HomeScore:  "120",
	//VisitScore: "99",
	//PeriodCn:   "第4节 05:30",
	//},
	//}
	i, err := newSelect(games)
	if err != nil {
		fmt.Printf("%s %s\n", Red(IconBad), "选择比赛错误")
		return
	}
	newUI(games[i])
}

func Red(str string) string {
	return fmt.Sprintf("\x1b[0;31m%s\x1b[0m", str)
}

func Green(str string) string {
	return fmt.Sprintf("\x1b[0;32m%s\x1b[0m", str)
}
