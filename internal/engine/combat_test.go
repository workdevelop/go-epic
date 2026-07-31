package engine

import (
	"go-epic/internal/models"
	"testing"
)

// MockUnit — тестова структура для детермінованих перевірок бою
type MockUnit struct {
	models.Position
	Name   string
	Health int
	Damage int
}

func (m *MockUnit) Move(dx, dy, w, h int) error { return nil }
func (m *MockUnit) IsAlive() bool               { return m.Health > 0 }
func (m *MockUnit) RandomStep() (int, int)      { return 0, 0 }
func (m *MockUnit) GetPosition() (int, int)     { return m.X, m.Y }
func (m *MockUnit) GetType() rune               { return 'U' }
func (m *MockUnit) GetName() string             { return m.Name }
func (m *MockUnit) GetDamage() int              { return m.Damage } // Завжди фіксований урон!
func (m *MockUnit) GetHealth() int              { return m.Health }
func (m *MockUnit) TakeDamage(amount int) {
	m.Health -= amount
	if m.Health < 0 {
		m.Health = 0
	}
}
func (m *MockUnit) Brain(idx int, c chan<- models.MoveEvent) {}

// TestBattle перевіряє логіку двосторонньої сутички
func TestBattle(t *testing.T) {
	u1 := &MockUnit{Name: "Атакуючий", Health: 100, Damage: 30}
	u2 := &MockUnit{Name: "Захисник", Health: 100, Damage: 20}

	// Викликаємо наш бойовий движок
	logs := Battle(u1, u2)

	// Очікуваний результат:
	// u1 б'є u2 на 30 -> у u2 залишається 70 HP (він живий)
	// u2 у відповідь б'є u1 на 20 -> у u1 залишається 80 HP
	if u2.GetHealth() != 70 {
		t.Errorf("Захисник отримав неправильний урон: залишилось %d HP, очікували 70", u2.GetHealth())
	}
	if u1.GetHealth() != 80 {
		t.Errorf("Атакуючий отримав неправильний урон у відповідь: залишилось %d HP, очікували 80", u1.GetHealth())
	}

	// Перевіряємо, що функція згенерувала рівно 2 рядки логів
	if len(logs) != 2 {
		t.Errorf("Очікували 2 рядки бойового логу, отримали: %d", len(logs))
	}
}
