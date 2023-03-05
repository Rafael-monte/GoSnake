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
	tcell.KeyEnter: func(game *Game) {
		if game.GameOver {
			go game.Run()
		}
	},
	tcell.KeyBackspace: func(game *Game) {
		if game.GameOver {
			game.Screen.Fini()
			os.Exit(0)
		}
	},
}

var optionsController = map[rune]func(game *Game){
	'y': func(game *Game) {
		go game.Run()
	},
	'n': func(game *Game) {
		game.Screen.Fini()
		os.Exit(0)
	},
}
