package models

type Unit interface {
	Move(dx, dy int)
	IsAlive() bool
}
