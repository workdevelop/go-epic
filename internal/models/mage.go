package models

import (
	"errors"
	"math/rand"
	"sync" // Імпортуємо sync
	"time"
)

type Mage struct {
	Position `json:"position"` // Вбудована структура серіалізується як окремий об'єкт
	Name     string            `json:"name"`
	Health   int               `json:"health"`
	Mana     int               `json:"mana"`
	mu       sync.Mutex
}

func (m *Mage) Move(dx, dy, worldWidth, worldHeight int) error {
	m.mu.Lock() // Блокуємо індивідуальну пам'ять мага
	defer m.mu.Unlock()

	nextX := m.X + dx
	nextY := m.Y + dy

	if nextX <= 0 || nextX >= worldWidth-1 || nextY <= 0 || nextY >= worldHeight-1 {
		return errors.New("маг уперся в магічний бар'єр карти")
	}

	m.X = nextX
	m.Y = nextY
	return nil
}

func (m *Mage) IsAlive() bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.Health > 0
}

func (m *Mage) TakeDamage(amount int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Health -= amount
	if m.Health < 0 {
		m.Health = 0
	}
}

func (m *Mage) GetHealth() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.Health
}

// Ці два методи можна залишити без змін, бо вони повертають копії значень
// або працюють з незмінними даними (Name), але для повної безпеки GetPosition теж захистимо:
func (m *Mage) GetPosition() (int, int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.X, m.Y
}

func (m *Mage) GetType() rune          { return 'M' }
func (m *Mage) GetName() string        { return m.Name }
func (m *Mage) GetDamage() int         { return rand.Intn(15) + 15 }
func (m *Mage) RandomStep() (int, int) { return rand.Intn(3) - 1, rand.Intn(3) - 1 }

func (m *Mage) Brain(unitIndex int, moveChan chan<- MoveEvent) {
	for {
		if !m.IsAlive() {
			return
		}

		delay := rand.Intn(200) + 100
		time.Sleep(time.Duration(delay) * time.Millisecond)

		if !m.IsAlive() {
			return
		}

		dx, dy := m.RandomStep()

		// 🔥 ЗАМІСТЬ МУТЕКСА: Просто кидаємо подію в канал.
		// Цей рядок є потокобезпечним. Стрілка вказує В КАНАЛ.
		moveChan <- MoveEvent{
			UnitIndex: unitIndex,
			DX:        dx,
			DY:        dy,
		}
	}
}
