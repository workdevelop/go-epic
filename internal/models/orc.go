package models

type Orc struct {
	Character
	Damage int
}

func NewOrc(name string, health, x, y, damage int) Orc {
	return Orc{
		Character: Character{
			Name:   name,
			Health: health,
			Position: Position{
				X: x,
				Y: y,
			},
		},
		Damage: damage,
	}
}
