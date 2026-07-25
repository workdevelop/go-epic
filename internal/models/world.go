package models

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
		MoveChannel: make(chan MoveEvent, 100), // Ініціалізуємо буфер тут!
	}
}

func (w *World) Tick() {}
