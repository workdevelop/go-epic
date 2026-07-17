package models

import "math/rand"

type Mage struct {
	Position
	Health int
	Mana   int
}

// TakeDamage зменшує здоров'я мага на вказану величину шкоди (мутатор).
func (m *Mage) TakeDamage(amount int) {
	m.Health -= amount
	if m.Health < 0 {
		m.Health = 0
	}
}

// IsAlive перевіряє, чи живий маг.
// Використовуємо ресивер-значення (m Mage) без зірочки, бо метод лише читає дані.
// Go зробить безпечну копію мага для виконання цієї функції.
func (m Mage) IsAlive() bool {
	return m.Health > 0
}

func (m *Mage) Move(dx, dy int) {
	m.X += dx
	m.Y += dy
}

// RandomStep вираховує випадкове зміщення для мага
func (m *Mage) RandomStep() (int, int) {
	dx := rand.Intn(3) - 1 // Повертає -1, 0 або 1
	dy := rand.Intn(3) - 1
	return dx, dy
}

func (m Mage) GetPosition() (int, int) {
	return m.X, m.Y
}

func (m Mage) GetType() rune {
	return 'M'
}
