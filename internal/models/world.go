package models

import "fmt"

type World struct {
	Width     int
	Height    int
	Units     []Unit
	BattleLog []string
}

func Battle(u1, u2 Unit) string {
	dmg := u1.GetDamage()
	hpBefore := u2.GetHealth()
	u2.TakeDamage(dmg)

	return fmt.Sprintf("⚔️ [%s] атакує [%s] і завдає %d шкоди! (Було %d HP -> Стало %d HP)", u1.GetName(), u2.GetName(), dmg, hpBefore, u2.GetHealth())
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
		x, y := w.Units[i].GetPosition()

		nextX := x + dx
		nextY := y + dy

		// Перевірка меж карти (стін)
		if nextX > 0 && nextX < w.Width-1 && nextY > 0 && nextY < w.Height-1 {
			w.Units[i].Move(dx, dy)
		}
	}

	// 2. Етап Бою (Пошук Колізій через універсальний Battle)
	for i := 0; i < len(w.Units); i++ {
		for j := i + 1; j < len(w.Units); j++ {
			u1 := w.Units[i]
			u2 := w.Units[j]

			// Перевірки: мертві не воюють, свої своїх не б'ють
			if !u1.IsAlive() || !u2.IsAlive() || u1.GetType() == u2.GetType() {
				continue
			}

			x1, y1 := u1.GetPosition()
			x2, y2 := u2.GetPosition()

			// Якщо юніти зустрілися на одній клітинці
			if x1 == x2 && y1 == y2 {
				// Запускаємо двосторонній бій через нашу універсальну функцію
				log1 := Battle(u1, u2) // u1 б'є u2
				fmt.Println(log1)
				log2 := Battle(u2, u1) // u2 б'є u1
				fmt.Println(log2)

				w.BattleLog = append(w.BattleLog, log1, log2)

				// Перевіряємо, чи хтось загинув після цієї сутички
				if !u1.IsAlive() {
					w.BattleLog = append(w.BattleLog, fmt.Sprintf("💀 %s загинув у епічній битві!", u1.GetName()))
				}
				if !u2.IsAlive() {
					w.BattleLog = append(w.BattleLog, fmt.Sprintf("💀 %s загинув у епічній битві!", u2.GetName()))
				}
			}
		}
	}
}
