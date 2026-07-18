package models

type Mage struct {
	Name   string
	Health int
	Mana   int
	Damage int
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

func (o Mage) GetPosition() (int, int) {
	return o.X, o.Y
}

func (o Mage) GetName() string {
	return o.Name
}

func (o Mage) GetDamage() int {
	return o.Damage
}
