package models

type World struct {
	Width  int
	Height int
	Mages  []Mage
	Orcs   []Orc
}

func NewWorld(width, height int, mages []Mage, orcs []Orc) World {
	world := World{
		Width:  width,
		Height: width,
		Mages:  mages,
		Orcs:   orcs,
	}
	return world
}
