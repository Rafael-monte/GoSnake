package main

import (
	"GoSnake/controller"
	"GoSnake/model"
	"log"

	"github.com/gdamore/tcell"
)

func main() {
	screen, err := tcell.NewScreen()

	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	screen.SetStyle(defStyle)

	snake := model.Snake{
		XAxis:  5,
		YAxis:  10,
		XSpeed: 1,
		YSpeed: 0,
	}

	game := controller.Game{
		Screen: screen,
		Snake:  snake,
	}

	go game.Run()
	for {
		game.ListenKeyEvents()
	}
}
