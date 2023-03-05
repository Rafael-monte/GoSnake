package model

import (
	"GoSnake/view"
)

const inMargin int = 0

type Snake struct {
	XAxis  int
	YAxis  int
	XSpeed int
	YSpeed int
}

func (s *Snake) ChangeDirection(coordinate view.Coordinate) {
	s.XSpeed = coordinate.X
	s.YSpeed = coordinate.Y
}

func (s *Snake) UpdateState(screenMeasurements view.ScreenMeasurements) {
	s.calculateHorizontalSpeed(screenMeasurements.Width)
	s.calculateVerticalSpeed(screenMeasurements.Height)
}

func (s *Snake) calculateHorizontalSpeed(width int) {
	s.XAxis = (s.XAxis + s.XSpeed) % width
	if s.XAxis < inMargin {
		s.XAxis += width
	}
}

func (s *Snake) calculateVerticalSpeed(height int) {
	s.YAxis = (s.YAxis + s.YSpeed) % height
	if s.YAxis < inMargin {
		s.YAxis += height
	}
}
