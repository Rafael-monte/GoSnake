package main

import (
	"GoSnake/controller"
	"GoSnake/model"
	"GoSnake/view"
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
	screen.SetStyle(view.GetDefaultStyle())

	snake := model.NewSnake()

	game := controller.Game{
		Screen: screen,
		Snake:  snake,
	}

	go game.Run()
	for {
		game.ListenKeyEvents()
	}
}
