package models

import (
	"math/rand"
	s "strings"
)

const MageHealth = 100
const MageDamage = 10
const MageMana = 100

const OrcHealth = 100
const OrcDamage = 10

const mageNames = "Alaric,Ambrose,Dorian,Eldon,Gideon,Ignatius,Zephyr,Volcanthus,Boreas,Morcant"
const orcNames = "Grutok,Urth,Thak,Rogar,Vrakul,Azog,Grommash,Thrall,Uglúk,Karg"

type World struct {
	Width  int
	Height int
	Units  []Unit
}

func (w *World) InitUnits(unitType string, numUnits int) {
	var names []string
	if unitType == "mage" {
		names = s.Split(mageNames, ",")
	} else {
		names = s.Split(orcNames, ",")
	}

	nameIndex := 0
	for range numUnits {
		x := rand.Intn(w.Width)
		y := rand.Intn(w.Height)

		if unitType == "mage" {
			mage := Mage{
				Name:   "Mage " + names[nameIndex],
				Health: MageHealth,
				Position: Position{
					X: x,
					Y: y,
				},
				Damage: MageDamage,
				Mana:   MageMana,
			}
			w.Units = append(w.Units, &mage)
		} else {
			orc := Orc{
				Name:   "Orc " + names[nameIndex],
				Health: OrcHealth,
				Position: Position{
					X: x,
					Y: y,
				},
				Damage: OrcDamage,
			}
			w.Units = append(w.Units, &orc)
		}

		nameIndex++
		if nameIndex >= len(names) {
			nameIndex = 0
		}
	}

}

func (w *World) Init(numMages, numOrcs int) {
	w.InitUnits("mage", numMages)
	w.InitUnits("orc", numOrcs)
}

func (w *World) RandomMoveUnits() {
	for i := 0; i < len(w.Units); i++ {
		if w.Units[i].IsAlive() {
			unitPosX, unitPosY := w.Units[i].GetPosition()
			dx := -1 + rand.Intn(3)
			if unitPosX+dx < 0 || unitPosX+dx >= w.Width {
				dx = 0
			}

			dy := -1 + rand.Intn(3)
			if unitPosY+dy < 0 || unitPosY+dy >= w.Height {
				dy = 0
			}

			w.Units[i].Move(dx, dy)
		}
	}
}

func (w *World) Tick() {
	w.RandomMoveUnits()
}
