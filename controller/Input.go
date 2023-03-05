package controller

import (
	"GoSnake/view"
	"os"

	"github.com/gdamore/tcell"
)

var inputController = map[tcell.Key]func(game *Game){
	tcell.KeyEscape: func(game *Game) {
		game.Screen.Fini()
		os.Exit(0)
	},
	tcell.KeyCtrlC: func(game *Game) {
		game.Screen.Fini()
		os.Exit(0)
	},
	tcell.KeyUp: func(game *Game) {
		game.Snake.ChangeDirection(view.Up)
	},
	tcell.KeyDown: func(game *Game) {
		game.Snake.ChangeDirection(view.Down)
	},
	tcell.KeyLeft: func(game *Game) {
		game.Snake.ChangeDirection(view.Left)
	},
	tcell.KeyRight: func(game *Game) {
		game.Snake.ChangeDirection(view.Right)
	},
}
