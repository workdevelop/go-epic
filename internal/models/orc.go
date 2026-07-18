package models

type Orc struct {
	Name   string
	Health int
	Damage int
	Position
}

func (o *Orc) Move(dx, dy int) {
	o.X += dx
	o.Y += dy
}

func (o *Orc) TakeDamage(amount int) {
	o.Health = max(o.Health-amount, 0)
}

func (o Orc) IsAlive() bool {
	return o.Health > 0
}

func (o Orc) GetPosition() (int, int) {
	return o.X, o.Y
}

func (o Orc) GetName() string {
	return o.Name
}

func (o Orc) GetDamage() int {
	return o.Damage
}
