package models

type Mage struct {
	Character
	Mana int
}

func NewMage(name string, health, x, y, mana int) Mage {
	return Mage{
		Character: Character{
			Name:   name,
			Health: health,
			Position: Position{
				X: x,
				Y: y,
			},
		},
		Mana: mana,
	}
}

// func (o *Mage) Move(dx, dy int) {
// 	o.X += dx
// 	o.Y += dy
// }

// func (o *Mage) TakeDamage(amount int) {
// 	o.Health -= amount
// }

// func (o Mage) IsAlive() bool {
// 	return o.Health > 0
// }
