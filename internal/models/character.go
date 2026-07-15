package models

type Character struct {
	Name   string
	Health int
	Position
}

func (o *Character) Move(dx, dy int) {
	o.X += dx
	o.Y += dy
}

func (o *Character) TakeDamage(amount int) int {
	o.Health = max(o.Health - amount, 0)
	return o.Health
}

func (o Character) IsAlive() bool {
	return o.Health > 0
}
