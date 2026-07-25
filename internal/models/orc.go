package models

import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

type Orc struct {
	Position
	Name   string
	Health int
	Damage int
	mu     sync.Mutex // 🔥 ЛОКАЛЬНИЙ ЗАМОК ОРКА
}

func (o *Orc) Move(dx, dy, worldWidth, worldHeight int) error {
	o.mu.Lock()
	defer o.mu.Unlock()

	nextX := o.X + dx
	nextY := o.Y + dy

	if nextX <= 0 || nextX >= worldWidth-1 || nextY <= 0 || nextY >= worldHeight-1 {
		return errors.New("орк намагався пробити стіну головою, хід заблоковано")
	}

	o.X = nextX
	o.Y = nextY
	return nil
}

func (o *Orc) IsAlive() bool {
	o.mu.Lock()
	defer o.mu.Unlock()
	return o.Health > 0
}

func (o *Orc) TakeDamage(amount int) {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.Health -= amount
	if o.Health < 0 {
		o.Health = 0
	}
}

func (o *Orc) GetHealth() int {
	o.mu.Lock()
	defer o.mu.Unlock()
	return o.Health
}

func (o *Orc) GetPosition() (int, int) {
	o.mu.Lock()
	defer o.mu.Unlock()
	return o.X, o.Y
}

func (o *Orc) GetType() rune          { return 'O' }
func (o *Orc) GetName() string        { return o.Name }
func (o *Orc) GetDamage() int         { return rand.Intn(10) + o.Damage }
func (o *Orc) RandomStep() (int, int) { return rand.Intn(3) - 1, rand.Intn(3) - 1 }

func (o *Orc) Brain(unitIndex int, moveChan chan<- MoveEvent) {
	for {
		if !o.IsAlive() {
			return
		}

		delay := rand.Intn(120) + 80
		time.Sleep(time.Duration(delay) * time.Millisecond)

		if !o.IsAlive() {
			return
		}

		dx, dy := o.RandomStep()

		// 🔥 Надсилаємо подію ходу орка
		moveChan <- MoveEvent{
			UnitIndex: unitIndex,
			DX:        dx,
			DY:        dy,
		}
	}
}
