package model

import "GoSnake/view"

const inMargin int = 0

type SnakePart struct {
	XAxis, YAxis int
}

func (sp *SnakePart) GetUpdatedPart(s *Snake, measurements view.ScreenMeasurements) SnakePart {
	newPart := *sp
	newPart.XAxis = (newPart.XAxis + s.XSpeed) % measurements.Width
	if newPart.XAxis < inMargin {
		newPart.XAxis += measurements.Width
	}
	newPart.YAxis = (newPart.YAxis + s.YSpeed) % measurements.Height
	if newPart.YAxis < inMargin {
		newPart.YAxis += measurements.Height
	}
	return newPart
}

func NewSnakeParts() []SnakePart {
	return []SnakePart{
		{
			XAxis: 5,
			YAxis: 10,
		},
		{
			XAxis: 6,
			YAxis: 10,
		},
		{
			XAxis: 7,
			YAxis: 10,
		},
	}
}
