package models

import "math/rand"

type Orc struct {
	Position
	Name   string
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

func (o Orc) GetPosition() (int, int) {
	return o.X, o.Y
}

func (o Orc) GetType() rune {
	return 'O'
}

func (o Orc) GetName() string {
	return o.Name
}

func (o Orc) GetDamage() int {
	return rand.Intn(10) + o.Damage
}

func (o Orc) GetHealth() int {
	return o.Health
}
