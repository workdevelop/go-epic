package models

type Unit interface {
	Move(mapWidth, mapHeight, dx, dy int) error
	IsAlive() bool
	GetPosition() (int, int)
	GetName() string
	GetDamage() int
	TakeDamage(amount int)
}
