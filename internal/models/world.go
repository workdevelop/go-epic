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

	// Етап Бою (Пошук Колізій) залишається з 6-го дня (виправлений варіант)
	for i := 0; i < len(w.Units); i++ {
		for j := i + 1; j < len(w.Units); j++ {
			u1, u2 := w.Units[i], w.Units[j]
			if !u1.IsAlive() || !u2.IsAlive() || u1.GetType() == u2.GetType() {
				continue
			}

			x1, y1 := u1.GetPosition()
			x2, y2 := u2.GetPosition()

			if x1 == x2 && y1 == y2 {
				// Пакет engine викликається на рівні main.go,
				// тут ми просто фіксуємо факт перетину в лог для main
				w.BattleLog = append(w.BattleLog, fmt.Sprintf("TRIGGER_BATTLE:%d:%d", i, j))
			}
		}
	}
}
