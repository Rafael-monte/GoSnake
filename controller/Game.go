package controller

import (
	"GoSnake/model"
	"GoSnake/view"
	"os"
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
		g.Screen.SetContent(g.Snake.XAxis, g.Snake.YAxis, ' ', nil, snakeStyle)
		time.Sleep(40 * time.Millisecond)
		g.Screen.Show()
	}
}

func (game *Game) ListenKeyEvents() {
	switch event := game.Screen.PollEvent().(type) {
	case *tcell.EventResize:
		game.Screen.Sync()
	case *tcell.EventKey:
		if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
			game.Screen.Fini()
			os.Exit(0)
		} else if event.Key() == tcell.KeyUp {
			game.Snake.ChangeDirection(view.Coordinate{Y: -1, X: 0})
		} else if event.Key() == tcell.KeyDown {
			game.Snake.ChangeDirection(view.Coordinate{Y: 1, X: 0})
		} else if event.Key() == tcell.KeyLeft {
			game.Snake.ChangeDirection(view.Coordinate{Y: 0, X: -1})
		} else if event.Key() == tcell.KeyRight {
			game.Snake.ChangeDirection(view.Coordinate{Y: 0, X: 1})
		}
	}
}
