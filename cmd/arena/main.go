package main

import (
	"fmt"
	"go-epic/internal/models"
	"time"
)

func main() {
	// Ініціалізуємо світ, де всі юніти лежать в одному зрізі
	world := models.World{
		Width:  10,
		Height: 10,
		Units: []models.Unit{
			&models.Mage{Position: models.Position{X: 2, Y: 2}, Name: "Гендальф", Health: 301, Mana: 150},
			&models.Mage{Position: models.Position{X: 3, Y: 3}, Name: "Саруман", Health: 220, Mana: 50},
			&models.Mage{Position: models.Position{X: 2, Y: 8}, Name: "Рейстлін", Health: 132, Mana: 70},
			&models.Mage{Position: models.Position{X: 4, Y: 5}, Name: "Медів", Health: 170, Mana: 140},
			&models.Mage{Position: models.Position{X: 3, Y: 9}, Name: "Джайна", Health: 121, Mana: 60},

			&models.Orc{Position: models.Position{X: 9, Y: 9}, Name: "Тралл", Health: 150, Damage: 20},
			&models.Orc{Position: models.Position{X: 8, Y: 8}, Name: "Громмаш", Health: 160, Damage: 25},
			&models.Orc{Position: models.Position{X: 9, Y: 2}, Name: "Дуротан", Health: 140, Damage: 18},
			&models.Orc{Position: models.Position{X: 7, Y: 5}, Name: "Оргрім", Health: 150, Damage: 22},
			&models.Orc{Position: models.Position{X: 8, Y: 3}, Name: "Гулдан", Health: 120, Damage: 30},
		},
	}

	fmt.Println("=== ДЕНЬ 7:  Анатомія довгоживучих процесів та time.Ticker ===")

	fmt.Println("\n🎬 Початковий стан мапи:")
	renderWorld(world)

	ticker := time.NewTicker(200 * time.Microsecond)
	defer ticker.Stop()
	turn := 0
	for range ticker.C {
		turn++
		liveMages, liveOrcs := 0, 0
		for _, u := range world.Units {
			if u.IsAlive() {
				if u.GetType() == 'M' {
					liveMages++
				} else {
					liveOrcs++
				}
			}
		}

		if liveMages == 0 {
			fmt.Println("\n🔴 ОРКИ ПЕРЕМОГЛИ! ГРА ЗАВЕРШЕНА.")
			break // Вихід з циклу тікера
		}
		if liveOrcs == 0 {
			fmt.Println("\n🔵 МАГИ ПЕРЕМОГЛИ! ГРА ЗАВЕРШЕНА.")
			break
		}
		// Виводимо поточний кадр гри
		fmt.Printf("\n--- КАДР №%d (FPS: 5) ---\n", turn)
		renderWorld(world)

		// Виводимо логи сутичок, якщо вони були на цьому кроці
		if len(world.BattleLog) > 0 {
			for _, log := range world.BattleLog {
				fmt.Println(log)
			}
		}

		// Прораховуємо наступний крок світу
		world.Tick()
		turn++
	}

	fmt.Println("Дякуємо за гру! Процес успішно завершено, пам'ять тікера звільнено через defer.")
}

// Нова оптимізована функція рендерингу
func renderWorld(w models.World) {
	for y := 0; y < w.Height; y++ {
		for x := 0; x < w.Width; x++ {
			// Стіни
			if x == 0 || x == w.Width-1 || y == 0 || y == w.Height-1 {
				fmt.Print("#")
				continue
			}

			// Перевіряємо, чи є БУДЬ-ЯКИЙ юніт на поточних координатах
			var foundUnit models.Unit
			for _, u := range w.Units {
				ux, uy := u.GetPosition()
				if ux == x && uy == y && u.IsAlive() {
					foundUnit = u
					break
				}
			}

			// Якщо знайшли юніта — малюємо його унікальний символ ('M' або 'O')
			if foundUnit != nil {
				fmt.Printf("%c", foundUnit.GetType()) // %c виводить rune як символ
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
