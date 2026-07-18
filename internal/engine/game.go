package engine

import (
	"fmt"
	"go-epic/internal/models"
	"math/rand"
)

type Engine struct {
	World *models.World
}

func (e *Engine) moveUnits() {
	w := e.World
	for i := 0; i < len(w.Units); i++ {
		if !e.World.Units[i].IsAlive() {
			continue
		}
		unitPosX, unitPosY := w.Units[i].GetPosition()

		dx := -1 + rand.Intn(3)
		if !w.IsExistingXPosition(unitPosX + dx) {
			fmt.Printf("%s не може зсунутись гориз. з %d на %d кроків !\n", w.Units[i].GetName(), unitPosX, dx)
			dx = 0
		}

		dy := -1 + rand.Intn(3)
		if !w.IsExistingYPosition(unitPosY + dy) {
			fmt.Printf("%s не може зсунутись верт. з %d на %d кроків !\n", w.Units[i].GetName(), unitPosY, dy)
			dy = 0
		}

		if !(dx == 0 && dy == 0) && !w.IsFreePosition(unitPosX+dx, unitPosY+dy) {
			fmt.Printf("%s не може перейти з (%d, %d) на (%d, %d) !\n", w.Units[i].GetName(), unitPosX, unitPosY, unitPosX+dx, unitPosY+dy)
			dx = 0
			dy = 0
		}

		w.Units[i].Move(dx, dy)
	}
}

func (e *Engine) battle(attacker, victim models.Unit) {
	damage := attacker.GetDamage()
	victim.TakeDamage(damage)
	healthInfo := ""
	if !victim.IsAlive() {
		healthInfo = fmt.Sprintf("[%s] мертвий", victim.GetName())
	}
	fmt.Printf("⚔️ [%s] атакує [%s] і завдає %d шкоди%s!\n", attacker.GetName(), victim.GetName(), damage, healthInfo)
}

func getUnitType(unit models.Unit) string {
	switch unit.(type) {
	case *models.Mage:
		return "mage"
	case *models.Orc:
		return "orc"
	}
	return ""
}

func (e *Engine) combats() {
	w := e.World
	for i := 0; i < len(w.Units); i++ {
		if !w.Units[i].IsAlive() {
			continue
		}

		srcType := getUnitType(w.Units[i])

		srcPosX, srcPosY := w.Units[i].GetPosition()
		rectX1 := srcPosX - 1
		rectY1 := srcPosY - 1
		rectX2 := srcPosX + 1
		rectY2 := srcPosY + 1

		for k := 0; k < len(w.Units); k++ {
			if !w.Units[k].IsAlive() {
				continue
			}
			targetType := getUnitType(w.Units[k])
			if targetType == srcType {
				continue
			}

			posX, posY := w.Units[k].GetPosition()
			if posX >= rectX1 && posY >= rectY1 && posX <= rectX2 && posY <= rectY2 {
				e.battle(w.Units[i], w.Units[k])
			}
		}
	}
}

func (w *Engine) sacrificeAllCorpses() {
	//
}

func (e *Engine) Tick() {
	// рухаємо живих персонажів на довільні клітини
	// зайняту клітину зайняти не можна
	e.moveUnits()

	// світ аналізує нові координати.
	// Якщо маг і орк опиняються на сусідніх клітинках - бійка
	e.combats()

	// Clean up Dead Bodies - мертві персонажі не будуть виводитись
	e.sacrificeAllCorpses()
}
