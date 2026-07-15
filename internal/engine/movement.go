package engine

import "go-epic/internal/models"

func MoveMage(m *models.Mage, dx int, dy int) {
	m.X += dx
	m.Y += dy
}

func MoveOrcWrong(o models.Orc, dx int, dy int) {
	o.X += dx
	o.Y += dy
}

func MoveOrc(o *models.Orc, dx int, dy int) {
	o.X += dx
	o.Y += dy
}
