package models

type Orc struct {
	Position
	Health int
	Damage int
}

// Move переміщує орка по карті (ресивер-вказівник для зміни стану).
func (o *Orc) Move(dx, dy int) {
	o.X += dx
	o.Y += dy
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
