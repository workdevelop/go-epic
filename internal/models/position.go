package models

type Position struct {
	X int
	Y int
}

func NewPosition(x, y int) Position {
	return Position{
		X: x,
		Y: y,
	}
}

func (o Position) IsAt(x, y int) bool {
	return o.X == x && o.Y == y
}
