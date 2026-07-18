package models

type Unit interface {
	Move(dx, dy int)
	IsAlive() bool
	GetPosition() (int, int)
	GetName() string
	GetDamage() int
	TakeDamage(amount int)
}
