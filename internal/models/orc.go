package models

import "math/rand"

type Orc struct {
	Position
	Health int
	Damage int
}

// TakeDamage завдає орку шкоди.
func (o *Orc) TakeDamage(amount int) {
	o.Health -= amount
	if o.Health < 0 {
		o.Health = 0
	}
}

// IsAlive перевіряє статус орка (ресивер-значення, безпечне читання).
func (o Orc) IsAlive() bool {
	return o.Health > 0
}

func (o *Orc) Move(dx, dy int) {
	o.X += dx
	o.Y += dy
}

func (o *Orc) RandomStep() (int, int) {
	dx := rand.Intn(3) - 1
	dy := rand.Intn(3) - 1
	return dx, dy
}

func (m Orc) GetPosition() (int, int) {
	return m.X, m.Y
}

func (m Orc) GetType() rune {
	return 'O'
}
