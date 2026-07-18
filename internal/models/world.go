package models

import (
	"fmt"
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

func (w World) isExsitingXPosition(x int) bool {
	return x >= 0 && x < w.Width
}

func (w World) isExsitingYPosition(y int) bool {
	return y >= 0 && y < w.Height
}

func (w World) isFreePosition(x, y int) bool {
	for i := 0; i < len(w.Units); i++ {
		posX, posY := w.Units[i].GetPosition()
		if posX == x && posY == y {
			return false
		}
	}
	return true
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
		var x, y int
		for {
			x = rand.Intn(w.Width)
			y = rand.Intn(w.Height)

			if w.isFreePosition(x, y) {
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
	w.InitUnits("mage", numMages)
	w.InitUnits("orc", numOrcs)
}

func (w *World) MoveUnits() {
	for i := 0; i < len(w.Units); i++ {
		if w.Units[i].IsAlive() {
			unitPosX, unitPosY := w.Units[i].GetPosition()

			dx := -1 + rand.Intn(3)
			if !w.isExsitingXPosition(unitPosX + dx) {
				fmt.Printf("%s не може зсунутись гориз. з %d на %d кроків !\n", w.Units[i].GetName(), unitPosX, dx)
				dx = 0
			}

			dy := -1 + rand.Intn(3)
			if !w.isExsitingYPosition(unitPosY + dy) {
				fmt.Printf("%s не може зсунутись верт. з %d на %d кроків !\n", w.Units[i].GetName(), unitPosY, dy)
				dy = 0
			}

			if !(dx == 0 && dy == 0) && !w.isFreePosition(unitPosX+dx, unitPosY+dy) {
				fmt.Printf("%s не може перейти з (%d, %d) на (%d, %d) !\n", w.Units[i].GetName(), unitPosX, unitPosY, unitPosX+dx, unitPosY+dy)
				dx = 0
				dy = 0
			}

			w.Units[i].Move(dx, dy)
		}
	}
}

func (w *World) GetTargetsForUnit(srcUnit Unit) []*Unit {
	var res []*Unit
	srcPosX, srcPosY := srcUnit.GetPosition()

	rectX1 := srcPosX - 1
	rectY1 := srcPosY - 1
	rectX2 := srcPosX + 1
	rectY2 := srcPosY + 1

	for i := 0; i < len(w.Units); i++ {
		if &w.Units[i] == &srcUnit {
			continue
		}

		currUnit := w.Units[i]
		posX, posY := currUnit.GetPosition()
		if posX >= rectX1 && posY >= rectY1 && posX <= rectX2 && posY <= rectY2 {
			res = append(res, &currUnit)
		}
	}

	return res
}

func getUnitType(unit Unit) string {
	switch unit.(type) {
	case *Mage:
		return "mage"
	case *Orc:
		return "orc"
	}
	return ""
}

func Battle(attacker, victim Unit) {
	damage := attacker.GetDamage()
	victim.TakeDamage(damage)
	healthInfo := ""
	if !victim.IsAlive() {
		healthInfo = fmt.Sprintf("[%s] мертвий", victim.GetName())
	}
	fmt.Printf("⚔️ [%s] атакує [%s] і завдає %d шкоди%s!\n", attacker.GetName(), victim.GetName(), damage, healthInfo)
}

func (w *World) Combats() {
	for i := 0; i < len(w.Units); i++ {
		srcType := getUnitType(w.Units[i])

		srcPosX, srcPosY := w.Units[i].GetPosition()
		rectX1 := srcPosX - 1
		rectY1 := srcPosY - 1
		rectX2 := srcPosX + 1
		rectY2 := srcPosY + 1

		for k := 0; k < len(w.Units); k++ {
			targetType := getUnitType(w.Units[k])
			if targetType == srcType {
				continue
			}

			posX, posY := w.Units[k].GetPosition()
			if posX >= rectX1 && posY >= rectY1 && posX <= rectX2 && posY <= rectY2 {
				Battle(w.Units[i], w.Units[k])
			}
		}
	}
}

func (w *World) SacrificeAllCorpses() {
	//
}

func (w *World) Tick() {
	// рухаємо живих персонажів на довільні клітини
	// зайняту клітину зайняти не можна
	w.MoveUnits()

	// світ аналізує нові координати.
	// Якщо маг і орк опиняються на сусідніх клітинках - бійка
	w.Combats()

	// Clean up Dead Bodies - мертві персонажі не будуть виводитись
	w.SacrificeAllCorpses()
}
