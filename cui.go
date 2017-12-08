package main

import tui "github.com/marcusolsson/tui-go"

func newUI() {
	scoreBox := tui.NewVBox()
	scoreBox.SetBorder(true)

	liveBox := tui.NewVBox()
	liveBox.SetBorder(true)
	liveBox.SetSizePolicy(tui.Maximum, tui.Expanding)

	root := tui.NewVBox(scoreBox, liveBox)
	ui := tui.New(root)
	ui.SetKeybinding("Esc", func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		panic(err)
	}
}
