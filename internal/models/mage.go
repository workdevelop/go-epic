package models

type Mage struct {
	Position
	Health int
	Mana   int
}

// Move змінює координати мага у світі.
// Використовуємо ресивер-вказівник (*Mage), щоб мутувати оригінальні дані.
func (m *Mage) Move(dx, dy int) {
	m.X += dx // Поля X та Y доступні напряму завдяки вбудовуванню Position
	m.Y += dy
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

func (m Mage) IsOnMap(w World) bool {
	return m.Y > 0 && m.Y <= w.Height && m.X > 0 && m.X <= w.Width
}
