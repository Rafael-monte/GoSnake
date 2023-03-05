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

func (s *Snake) UpdateState(screenMeasurements view.ScreenMeasurements, longerSnake bool) {
	s.SnakeParts = append(s.SnakeParts, s.SnakeParts[len(s.SnakeParts)-1].GetUpdatedPart(s, screenMeasurements))
	if !longerSnake {
		s.SnakeParts = s.SnakeParts[1:]
	}

}

func (sb *Snake) ResetPos(coordinate view.Coordinate) {
	snakeParts := []SnakePart{
		{
			XAxis: int(coordinate.X / 2),
			YAxis: int(coordinate.Y / 2),
		},
		{
			XAxis: int(coordinate.X/2) + 1,
			YAxis: int(coordinate.Y / 2),
		},
		{
			XAxis: int(coordinate.X/2) + 2,
			YAxis: int(coordinate.Y / 2),
		},
	}
	sb.SnakeParts = snakeParts
	sb.XSpeed = 1
	sb.YSpeed = 0
}

func NewSnake() Snake {
	return Snake{
		XSpeed:     1,
		YSpeed:     0,
		SnakeParts: NewSnakeParts(),
	}
}
