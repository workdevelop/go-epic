package models

type Unit interface {
	Move(dx, dy int)
	IsAlive() bool
	RandomStep() (int, int)
	GetPosition() (int, int) // Повертає X та Y одночасно
	GetType() rune           // Повертає символ 'M' або 'O'
}
