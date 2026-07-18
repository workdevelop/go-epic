package models

type Mage struct {
	Name   string
	Health int
	Mana   int
	Position
}

func (o *Mage) Move(dx, dy int) {
	o.X += dx
	o.Y += dy
}

func (o *Mage) TakeDamage(amount int) {
	o.Health = max(o.Health-amount, 0)
}

func (o Mage) IsAlive() bool {
	return o.Health > 0
}
