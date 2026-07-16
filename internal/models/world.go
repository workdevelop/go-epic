package models

import (
	"math/rand"
)

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

func (w World) CharacterCanMoveByX(character Character, dx int) bool {
	return character.X+dx >= 0 && character.X+dx < w.Width
}

func (w World) CharacterCanMoveByY(character Character, dy int) bool {
	return character.Y+dy >= 0 && character.Y+dy < w.Height
}

func (w *World) RandomMoveMages() {
	for i := 0; i < len(w.Mages); i++ {
		dx := -1 + rand.Intn(3)
		if !w.CharacterCanMoveByX(w.Mages[i].Character, dx) {
			dx = 0
		}

		dy := -1 + rand.Intn(3)
		if !w.CharacterCanMoveByY(w.Mages[i].Character, dy) {
			dy = 0
		}

		w.Mages[i].Move(dx, dy)
	}
}

func (w *World) RandomMoveOrcs() {
	for i := 0; i < len(w.Orcs); i++ {
		dx := -1 + rand.Intn(3)
		if !w.CharacterCanMoveByX(w.Orcs[i].Character, dx) {
			dx = 0
		}

		dy := -1 + rand.Intn(3)
		if !w.CharacterCanMoveByY(w.Orcs[i].Character, dy) {
			dy = 0
		}

		w.Orcs[i].Move(dx, dy)
	}
}

func (w *World) Tick() {
	w.RandomMoveMages()
	w.RandomMoveOrcs()
}
