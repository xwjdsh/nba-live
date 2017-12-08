package main

import (
	"log"
	"os"
)

func main() {
	games := []*Game{
		&Game{
			HomeTeam:   "骑士",
			VisitTeam:  "勇士",
			HomeScore:  "100",
			VisitScore: "90",
			PeriodCn:   "第4节 01:30",
		},
		&Game{
			HomeTeam:   "凯尔特人",
			VisitTeam:  "公牛",
			HomeScore:  "80",
			VisitScore: "88",
			PeriodCn:   "第3节 03:30",
		},
	}
	i, err := newSelect(games)
	if err != nil {
		os.Exit(1)
	}
	log.Println(i)
	newUI()
}
