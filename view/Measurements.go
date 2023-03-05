package view

type Coordinate struct {
	X, Y int
}

var (
	Up    Coordinate = Coordinate{Y: -1, X: 0}
	Down  Coordinate = Coordinate{Y: 1, X: 0}
	Left  Coordinate = Coordinate{Y: 0, X: -1}
	Right Coordinate = Coordinate{Y: 0, X: 1}
)

type ScreenMeasurements struct {
	Width, Height int
}

func GetCenterOfScreen(width, height int) (int, int, int, int) {
	return height/2 - 20, width / 2, height/2 + 20, width / 2
}

func NewScreenMeasurements(Width, Height int) ScreenMeasurements {
	return ScreenMeasurements{
		Height: Height,
		Width:  Width,
	}
}
