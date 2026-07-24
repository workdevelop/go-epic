package models

import (
	"errors"
	"math/rand"
	"time"
)

type Orc struct {
	Position
	Name   string
	Health int
	Damage int
}

// TakeDamage завдає орку шкоди.
func (o *Orc) TakeDamage(amount int) {
	o.Health -= amount
	if o.Health < 0 {
		o.Health = 0
	}
}

// IsAlive перевіряє статус орка (ресивер-значення, безпечне читання).
func (o Orc) IsAlive() bool {
	return o.Health > 0
}

func (o *Orc) Move(dx, dy, worldWidth, worldHeight int) error {
	if o.X+dx <= 0 || o.X+dx >= worldWidth-1 || o.Y+dy <= 0 || o.Y+dy >= worldHeight-1 {
		return errors.New("Орк впилячився головою об стіни, намагаючись втекти")
	}

	o.X += dx
	o.Y += dy
	return nil
}

func (o *Orc) RandomStep() (int, int) {
	dx := rand.Intn(3) - 1
	dy := rand.Intn(3) - 1
	return dx, dy
}

func (o Orc) GetPosition() (int, int) {
	return o.X, o.Y
}

func (o Orc) GetType() rune {
	return 'O'
}

func (o Orc) GetName() string {
	return o.Name
}

func (o Orc) GetDamage() int {
	return rand.Intn(10) + o.Damage
}

func (o Orc) GetHealth() int {
	return o.Health
}

func (o *Orc) Brain(worldWidth, worldHeight int) {
	for {

		if !o.IsAlive() {
			return
		}
		ms := rand.Intn(200) + 100
		time.Sleep(time.Duration(ms) * time.Millisecond)
		if !o.IsAlive() {
			return
		}
		dx, dy := o.RandomStep()

		_ = o.Move(dx, dy, worldWidth, worldHeight)
	}
}
