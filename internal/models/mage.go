package models

import (
	"errors"
	"math/rand"
	"time"
)

type Mage struct {
	Position
	Name   string
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

func (m *Mage) Move(dx, dy, worldWidth, worldHeight int) error {

	if m.X+dx <= 0 || m.X+dx >= worldWidth-1 || m.Y+dy <= 0 || m.Y+dy >= worldHeight-1 {
		return errors.New("Маг намагався втекти з поля. Хід заблоковано")
	}
	m.X += dx
	m.Y += dy
	return nil
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

func (m Mage) GetName() string {
	return m.Name
}

func (m Mage) GetHealth() int {
	return m.Health
}

func (m Mage) GetDamage() int {
	return rand.Intn(15) + 15
}

func (m *Mage) Brain(worldWidth, worldHeight int) {
	for {

		if !m.IsAlive() {
			return
		}
		ms := rand.Intn(120) + 80
		time.Sleep(time.Duration(ms) * time.Millisecond)
		if !m.IsAlive() {
			return
		}
		dx, dy := m.RandomStep()

		_ = m.Move(dx, dy, worldWidth, worldHeight)
	}
}
