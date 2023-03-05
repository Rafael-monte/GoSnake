package view

type Coordinate struct {
	X, Y int
}

type ScreenMeasurements struct {
	Width, Height int
}

func NewScreenMeasurements(Width, Height int) ScreenMeasurements {
	return ScreenMeasurements{
		Height: Height,
		Width:  Width,
	}
}
