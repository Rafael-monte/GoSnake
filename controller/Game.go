package controller

import (
	"GoSnake/model"
	"GoSnake/view"
	"math/rand"
	"strconv"
	"time"

	"github.com/gdamore/tcell"
)

type Game struct {
	Screen   tcell.Screen
	Snake    model.Snake
	GameOver bool
	FoodPos  model.SnakePart
	Score    int
}

func (g *Game) Run() {
	defaultStyle := view.GetDefaultStyle()
	g.Screen.SetStyle(defaultStyle)
	height, width := g.Screen.Size()
	g.Snake.ResetPos(view.Coordinate{X: width, Y: height})
	g.UpdateFoodPos(width, height)
	g.GameOver = false
	g.Score = 0
	screenMeasurements := view.NewScreenMeasurements(height, width)
	snakeStyle := view.GetSnakeStyle()

	for {
		longerSnake := false
		g.Screen.Clear()
		if checkCollision(g.Snake.SnakeParts[len(g.Snake.SnakeParts)-1:], g.FoodPos) {
			//g.UpdateFoodPos(width, height)
			g.UpdateFoodPos(width, height)
			longerSnake = true
			g.Score++
		}
		if checkCollision(g.Snake.SnakeParts[:len(g.Snake.SnakeParts)-1], g.Snake.SnakeParts[len(g.Snake.SnakeParts)-1]) {
			break
		}
		g.Screen.Clear()
		g.Snake.UpdateState(screenMeasurements, longerSnake)
		drawParts(g.Screen, g.Snake.SnakeParts, g.FoodPos, snakeStyle, defaultStyle)
		drawText(g.Screen, 1, 1, 8+len(strconv.Itoa(g.Score)), 1, "Score: "+strconv.Itoa(g.Score))
		time.Sleep(60 * time.Millisecond)
		g.Screen.Show()
	}
	g.GameOver = true
	drawText(g.Screen, height/2-20, width/2, height/2+20, width/2, "Fim de jogo!, Pontos: "+strconv.Itoa(g.Score)+", Recomeçar? y/n")
	g.Screen.Show()
}

func (game *Game) ListenKeyEvents() {
	switch event := game.Screen.PollEvent().(type) {
	case *tcell.EventResize:
		game.Screen.Sync()
	case *tcell.EventKey:
		if !game.GameOver {
			inputController[event.Key()](game)
		}
		if _, isPresent := optionsController[event.Rune()]; isPresent {
			optionsController[event.Rune()](game)
		}
	}
}

func (g *Game) UpdateFoodPos(width, height int) {
	g.FoodPos.XAxis = rand.Intn(height)
	g.FoodPos.YAxis = rand.Intn(width)
	if g.FoodPos.YAxis == 1 && g.FoodPos.XAxis < 10 {
		g.UpdateFoodPos(width, height)
	}
}

func drawParts(s tcell.Screen, parts []model.SnakePart, foodPos model.SnakePart, snakeStyle, foodStyle tcell.Style) {
	s.SetContent(foodPos.XAxis, foodPos.YAxis, '\u25CF', nil, foodStyle)

	for _, part := range parts {
		s.SetContent(part.XAxis, part.YAxis, ' ', nil, snakeStyle)
	}
}

func drawText(s tcell.Screen, x1, y1, x2, y2 int, text string) {
	row := y1
	col := x1
	style := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	for _, r := range text {
		s.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}

func checkCollision(parts []model.SnakePart, otherPart model.SnakePart) bool {
	for _, part := range parts {
		if part.XAxis == otherPart.XAxis && part.YAxis == otherPart.YAxis {
			return true
		}
	}
	return false
}
