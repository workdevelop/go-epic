package models

import "errors"

type Orc struct {
	Name   string
	Health int
	Damage int
	Position
}

func (o *Orc) Move(mapWidth, mapHeight, dx, dy int) error {
	if o.X+dx < 0 || o.X+dx >= mapWidth || o.Y+dy < 0 || o.Y+dy >= mapHeight {
		return errors.New("out of bounds")
	}

	o.X += dx
	o.Y += dy

	return nil
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
