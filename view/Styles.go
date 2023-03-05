package view

import "github.com/gdamore/tcell"

func GetDefaultStyle() tcell.Style {
	return tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
}

func GetSnakeStyle() tcell.Style {
	return tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorWhite)
}
