package main

import (
	"bytes"
	"fmt"
	"log"
	"time"

	figure "github.com/common-nighthawk/go-figure"
	tui "github.com/marcusolsson/tui-go"
)

var (
	scoreChan = make(chan *GameInfo)
	overChan  = make(chan bool)
)

func newUI(game *Game) {
	grid := tui.NewGrid(0, 0)

	homeScore, visitScore := drawScore(game.HomeScore.(string), game.VisitScore.(string))
	homeLabel := tui.NewLabel(homeScore)
	visitLabel := tui.NewLabel(visitScore)
	homeLabel.SetStyleName("score")
	visitLabel.SetStyleName("score")

	grid.AppendRow(tui.NewHBox(homeLabel, visitLabel))

	scoreBox := tui.NewVBox(grid)
	scoreBox.SetTitle(fmt.Sprintf("%s VS %s", game.HomeTeam, game.VisitTeam))
	scoreBox.SetBorder(true)

	liveBox := tui.NewVBox()
	liveBox.SetTitle("直播")
	liveBox.SetBorder(true)
	liveBox.SetSizePolicy(tui.Maximum, tui.Expanding)

	theme := tui.NewTheme()
	theme.SetStyle("label.score", tui.Style{Fg: tui.ColorCyan, Bold: true})

	root := tui.NewVBox(scoreBox, liveBox)
	ui := tui.New(root)
	ui.SetTheme(theme)
	ui.SetKeybinding("Esc", func() { ui.Quit() })
	ui.SetKeybinding("q", func() { ui.Quit() })
	go update(ui, homeLabel, visitLabel)
	go fetch(game.ID)
	if err := ui.Run(); err != nil {
		panic(err)
	}
}

func update(ui tui.UI, homeLabel, visitLabel *tui.Label) {
	for {
		select {
		case info := <-scoreChan:
			if info != nil {
				ui.Update(func() {
					homeScore, visitScore := drawScore(info.HomeScore, info.VisitScore)
					homeLabel.SetText(homeScore)
					visitLabel.SetText(visitScore)
				})
			}
		case <-overChan:
			close(scoreChan)
			return
		}
	}
}

func fetch(gameId string) {
	lastMaxSid := 0
	for {
		maxSid, err := getMaxsid(gameId)
		if err != nil {
			log.Println("get max sid error:", err.Error())
			time.Sleep(3 * time.Second)
			continue
		}
		if maxSid > lastMaxSid {
			gameInfo, err := getGameInfo(gameId)
			if err != nil {
				log.Println("get game info error:", err.Error())
				time.Sleep(3 * time.Second)
				continue
			}
			lastMaxSid = maxSid
			time.Sleep(2 * time.Second)
			scoreChan <- gameInfo
		}
	}
}

func drawScore(homeScore, visitScore string) (string, string) {
	homeBuf := new(bytes.Buffer)
	figure.Write(homeBuf, figure.NewFigure(homeScore, "slant", true))
	visitBuf := new(bytes.Buffer)
	figure.Write(visitBuf, figure.NewFigure(visitScore, "slant", true))

	return homeBuf.String(), visitBuf.String()
}
