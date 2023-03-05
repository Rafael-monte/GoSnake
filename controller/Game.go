package controller

import (
	"GoSnake/model"
	"GoSnake/view"
	"time"

	"github.com/gdamore/tcell"
)

type Game struct {
	Screen tcell.Screen
	Snake  model.Snake
}

func (g *Game) Run() {
	defaultStyle := view.GetDefaultStyle()
	g.Screen.SetStyle(defaultStyle)
	height, width := g.Screen.Size()
	screenMeasurements := view.NewScreenMeasurements(height, width)
	snakeStyle := view.GetSnakeStyle()

	for {
		g.Screen.Clear()
		g.Snake.UpdateState(screenMeasurements)
		drawParts(g.Screen, g.Snake.SnakeParts, snakeStyle)
		time.Sleep(40 * time.Millisecond)
		g.Screen.Show()
	}
}

func (game *Game) ListenKeyEvents() {
	switch event := game.Screen.PollEvent().(type) {
	case *tcell.EventResize:
		game.Screen.Sync()
	case *tcell.EventKey:
		inputController[event.Key()](game)
	}
}

func drawParts(s tcell.Screen, parts []model.SnakePart, style tcell.Style) {
	for _, part := range parts {
		s.SetContent(part.XAxis, part.YAxis, ' ', nil, style)
	}
}
