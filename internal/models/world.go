package models

import (
	"encoding/json"
	"math"
	"sync"
)

type World struct {
	Width       int            `json:"width"`
	Height      int            `json:"height"`
	Units       []Unit         `json:"-"` // Ігноруємо сирі інтерфейси при дефолтному маршалінгу
	BattleLog   []string       `json:"battle_log"`
	MoveChannel chan MoveEvent `json:"-"` // Канали в JSON передавати не можна!
	mu          sync.Mutex
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

// 🔥 ДОДАЄМО МЕТОДИ БЛОКУВАННЯ, ЯКІ ЗАГУБИЛИСЯ:
// Lock примусово закриває замок світу для зовнішніх пакетів (наприклад, для main.go)
func (w *World) Lock() {
	w.mu.Lock()
}

// Unlock відкриває замок світу
func (w *World) Unlock() {
	w.mu.Unlock()
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

// 🏗️ СТВОРЮЄМО DTO СТРУКТУРИ ДЛЯ WEB API
// Вони містять лише чисті дані без мутексів, готові для безпечного читання рефлексією JSON
type UnitDTO struct {
	Type   string `json:"type"` // "Mage" або "Orc"
	Name   string `json:"name"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Health int    `json:"health"`
}

// MarshalJSON — тепер робить безпечний Snapshot під захистом мутексу світу
func (w *World) MarshalJSON() ([]byte, error) {
	// 1. 🔥 ГАРАНТОВАНО БЛОКУЄМО СВІТ:
	// Поки HTTP-горутина зчитує дані для JSON, ігровий цикл не зможе порушити пам'ять
	w.mu.Lock()
	defer w.mu.Unlock()

	// 2. Знімаємо безпечну копію (Snapshot) всіх юнітів у формат DTO
	dtoUnits := make([]UnitDTO, 0, len(w.Units))
	for _, u := range w.Units {
		x, y := u.GetPosition()

		unitType := "Orc"
		if u.GetType() == 'M' {
			unitType = "Mage"
		}

		dtoUnits = append(dtoUnits, UnitDTO{
			Type:   unitType,
			Name:   u.GetName(),
			X:      x,
			Y:      y,
			Health: u.GetHealth(),
		})
	}

	// 3. Також копіюємо лог боїв
	logCopy := make([]string, len(w.BattleLog))
	copy(logCopy, w.BattleLog)

	// 4. Серіалізуємо безпечний Snapshot, до якого ігровий цикл вже не має доступу
	return json.Marshal(&struct {
		Width     int       `json:"width"`
		Height    int       `json:"height"`
		BattleLog []string  `json:"battle_log"`
		Units     []UnitDTO `json:"units"`
	}{
		Width:     w.Width,
		Height:    w.Height,
		BattleLog: logCopy,
		Units:     dtoUnits,
	})
}
