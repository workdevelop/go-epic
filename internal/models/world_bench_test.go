package models

import (
	"fmt"
	"math/rand"
	"testing"
)

// BenchmarkFindClosestEnemy заміряє швидкість та споживання пам'яті алгоритму
func BenchmarkFindClosestEnemy(b *testing.B) {
	// 1. Готуємо велику арену для тесту (100 магів та 100 орків)
	world := NewWorld(40, 40)
	for i := 0; i < 100; i++ {
		world.Units = append(world.Units, &Mage{Position: Position{X: rand.Intn(38) + 1, Y: rand.Intn(38) + 1}, Health: 100, Name: fmt.Sprintf("M-%d", i)})
		world.Units = append(world.Units, &Orc{Position: Position{X: rand.Intn(38) + 1, Y: rand.Intn(38) + 1}, Health: 100, Name: fmt.Sprintf("O-%d", i)})
	}

	testUnit := world.Units[0] // Беремо першого мага як піддослідного

	b.ResetTimer() // Скидаємо таймер, щоб не враховувати час генерації світу

	// 2. Рантайм Go сам вирішить, скільки мільйонів разів (b.N) прогнати цей цикл для точного заміру
	for i := 0; i < b.N; i++ {
		_, _, _ = world.FindClosestEnemy(testUnit)
	}
}
