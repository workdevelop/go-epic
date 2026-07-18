package models

type Position struct {
	X int
	Y int
}

func (o Position) IsAt(x, y int) bool {
	return o.X == x && o.Y == y
}
