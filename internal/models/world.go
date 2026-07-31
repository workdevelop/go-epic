package models

import "math"

type World struct {
	Width       int
	Height      int
	Units       []Unit
	BattleLog   []string
	MoveChannel chan MoveEvent // 🔥 ТЕПЕР КАНАЛ НАЛЕЖИТЬ СВІТУ
}

// NewWorld — конструктор для безпечного створення світу з ініціалізованим каналом
func NewWorld(width, height int) *World {
	return &World{
		Width:       width,
		Height:      height,
		Units:       []Unit{},
		BattleLog:   []string{},
		MoveChannel: make(chan MoveEvent, 1000), // Ініціалізуємо буфер тут!
	}
}

func (w *World) Tick() {}

// FindClosestEnemy — оптимізована версія (Zero Allocation)
func (w *World) FindClosestEnemy(currentUnit Unit) (int, int, bool) {
	cx, cy := currentUnit.GetPosition()

	closestX, closestY := 0, 0
	var minDistanceSq int64 = math.MaxInt64 // Використовуємо цілі числа int64
	found := false

	// Прямий перебір без виклику інтерфейсних методів усередині формули
	for i := 0; i < len(w.Units); i++ {
		enemy := w.Units[i]
		if !enemy.IsAlive() || enemy.GetType() == currentUnit.GetType() {
			continue
		}

		ex, ey := enemy.GetPosition()

		// Оптимізована математика: рахуємо тільки квадрати відстаней (без math.Sqrt!)
		dx := int64(ex - cx)
		dy := int64(ey - cy)
		distanceSq := dx*dx + dy*dy

		if distanceSq < minDistanceSq {
			minDistanceSq = distanceSq
			closestX = ex
			closestY = ey
			found = true
		}
	}

	return closestX, closestY, found
}
