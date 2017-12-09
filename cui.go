package main

import (
	"bytes"
	"fmt"

	figure "github.com/common-nighthawk/go-figure"
	tui "github.com/marcusolsson/tui-go"
)

func newUI(game *Game) {

	grid := tui.NewGrid(0, 0)

	homeBuf := new(bytes.Buffer)
	figure.Write(homeBuf, figure.NewFigure(game.HomeScore.(string), "slant", true))
	visitBuf := new(bytes.Buffer)
	figure.Write(visitBuf, figure.NewFigure(game.VisitScore.(string), "slant", true))

	homeLabel := tui.NewLabel(homeBuf.String())
	visitLabel := tui.NewLabel(visitBuf.String())
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

	if err := ui.Run(); err != nil {
		panic(err)
	}
}
