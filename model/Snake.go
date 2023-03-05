package model

import (
	"GoSnake/view"
)

type Snake struct {
	XSpeed     int
	YSpeed     int
	SnakeParts []SnakePart
}

func (s *Snake) ChangeDirection(coordinate view.Coordinate) {
	s.XSpeed = coordinate.X
	s.YSpeed = coordinate.Y
}

func (s *Snake) UpdateState(screenMeasurements view.ScreenMeasurements) {
	s.SnakeParts = append(s.SnakeParts, s.SnakeParts[len(s.SnakeParts)-1].GetUpdatedPart(s, screenMeasurements))
	s.SnakeParts = s.SnakeParts[1:]
}

func NewSnake() Snake {
	return Snake{
		XSpeed:     1,
		YSpeed:     0,
		SnakeParts: NewSnakeParts(),
	}
}
