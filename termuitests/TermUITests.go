package main

import (
	terminalui "ml253/graphictermui"

	color "golang.org/x/image/colornames"
)

func main() {
	terminalui.NewTerminalUI(mainThread, 1280, 800)
}

func mainThread(tui *terminalui.TermUI) {
	tui.View.SetFillChar('+')
	tui.View.SetColor(color.Green, color.Black)
	tui.View.FillRect(0, 0, 100, 40)
	tui.View.SetColor(color.Red, color.Yellow)
	tui.View.FillRect(50, 0, 30, 30)
	tui.View.DrawLine(5, 5, 10, 10)
	tui.View.DrawLine(10, 5, 5, 10)
	tui.View.ClearLines()
	tui.View.DrawLine(10, 10, 10, 20)
	tui.View.DrawLine(10, 20, 20, 10)
}
