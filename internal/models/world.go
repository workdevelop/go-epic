package models

type World struct {
	Width  int
	Height int
	Units  []Unit
}

// Tick симулює один крок ігрового часу.
// Світ перебирає всіх юнітів і змушує їх зробити хід.
func (w *World) Tick() {
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
}
