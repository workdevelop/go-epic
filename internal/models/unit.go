package models

type Unit interface {
	Move(dx, dy, worldWidth, worldHeight int) error
	IsAlive() bool
	RandomStep() (int, int)
	GetPosition() (int, int) // Повертає X та Y одночасно
	GetType() rune           // Повертає символ 'M' або 'O'3

	GetName() string
	GetDamage() int
	GetHealth() int
	TakeDamage(amount int)
	Brain(world *World)
}
