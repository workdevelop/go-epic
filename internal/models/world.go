package models

import (
	"math/rand"
	s "strings"
)

type World struct {
	Width  int
	Height int
	Units  []Unit
}

const mageNames = "Alaric,Ambrose,Dorian,Eldon,Gideon,Ignatius,Zephyr,Volcanthus,Boreas,Morcant"
const orcNames = "Grutok,Urth,Thak,Rogar,Vrakul,Azog,Grommash,Thrall,Uglúk,Karg"

const MageHealth = 100
const MageDamage = 10
const MageMana = 100

const OrcHealth = 100
const OrcDamage = 10

func (w World) IsFreePosition(x, y int) bool {
	for i := 0; i < len(w.Units); i++ {
		posX, posY := w.Units[i].GetPosition()
		if posX == x && posY == y {
			return false
		}
	}
	return true
}

func (w *World) initUnits(unitType string, numUnits int) {
	var names []string
	if unitType == "mage" {
		names = s.Split(mageNames, ",")
	} else {
		names = s.Split(orcNames, ",")
	}

	nameIndex := 0
	for range numUnits {
		var x, y int
		for {
			x = rand.Intn(w.Width)
			y = rand.Intn(w.Height)

			if w.IsFreePosition(x, y) {
				break
			}
		}

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
	w.initUnits("mage", numMages)
	w.initUnits("orc", numOrcs)
}
