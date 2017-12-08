package main

import (
	"github.com/xwjdsh/promptui"
)

func newSelect(games []*Game) (int, error) {

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Help:     "使用 [↑/k],[↓/j] 选择比赛",
		Active:   "\U0001F336  {{ .HomeTeam | cyan }} {{ `VS` | cyan }} {{ .VisitTeam | cyan }}",
		Inactive: "    {{ .HomeTeam }} VS {{ .VisitTeam }}",
		Details: `
--------- {{ .PeriodCn }} ----------
    当前比分 -- {{ .HomeScore | red }} : {{ .VisitScore | red }}`,
	}

	prompt := promptui.Select{
		Label:     "",
		Items:     games,
		Templates: templates,
		Size:      4,
	}

	i, _, err := prompt.Run()
	return i, err
}
