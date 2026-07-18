package models

import (
	"math/rand"
	"strconv"
)

const MageHealth = 100
const MageMana = 100

const OrcHealth = 200
const OrcDamage = 10

type World struct {
	Width  int
	Height int
	Units  []Unit
}

func (w *World) InitMages(numUnits int) {
	xMage := 2
	y := 2

	for i := range numUnits {
		mage := Mage{
			Name:   "mage " + strconv.Itoa(i+1),
			Health: MageHealth,
			Position: Position{
				X: xMage,
				Y: y,
			},
			Mana: MageMana,
		}
		w.Units = append(w.Units, &mage)

		y = y + 1
	}
}

func (w *World) InitOrcs(numUnits int) {
	xOrc := 7
	y := 2

	for i := range numUnits {
		orc := Orc{
			Name:   "orc " + strconv.Itoa(i+1),
			Health: MageHealth,
			Position: Position{
				X: xOrc,
				Y: y,
			},
			Damage: OrcDamage,
		}
		w.Units = append(w.Units, &orc)

		y = y + 1
	}
}

func (w *World) Init(numMages, numOrcs int) {
	w.InitMages(numMages)
	w.InitOrcs(numOrcs)
}

func (w *World) RandomMoveUnits() {
	for i := 0; i < len(w.Units); i++ {
		if w.Units[i].IsAlive() {
			dx := -1 + rand.Intn(3)
			dy := -1 + rand.Intn(3)
			w.Units[i].Move(dx, dy)
		}
	}
}

func (w *World) Tick() {
	w.RandomMoveUnits()
}
