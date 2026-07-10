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
