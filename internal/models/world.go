package models

import (
	"fmt"
)

type World struct {
	Width     int
	Height    int
	Units     []Unit
	BattleLog []string
}

// Tick симулює один крок ігрового часу.
// Світ перебирає всіх юнітів і змушує їх зробити хід.
func (w *World) Tick() {
	w.BattleLog = []string{}

	for i := 0; i < len(w.Units); i++ {
		// Якщо юніт мертвий — пропускаємо його хід
		if !w.Units[i].IsAlive() {
			continue
		}

		dx, dy := w.Units[i].RandomStep()
		err := w.Units[i].Move(dx, dy, w.Width, w.Height)
		if err != nil {
			logMsg := fmt.Sprintf("⚠️ Помилка ходу [%s]: %s", w.Units[i].GetName(), err.Error())
			w.BattleLog = append(w.BattleLog, logMsg)
			continue
		}

	}
}
