package main

import (
	"fmt"
	"go-epic/internal/models"
)

func main() {
	// Ініціалізуємо світ, де всі юніти лежать в одному зрізі
	world := models.World{
		Width:  10,
		Height: 10,
		Units: []models.Unit{
			&models.Mage{Position: models.Position{X: 2, Y: 2}, Name: "Гендальф", Health: 100, Mana: 50},
			&models.Mage{Position: models.Position{X: 3, Y: 3}, Name: "Саруман", Health: 100, Mana: 50},
			&models.Mage{Position: models.Position{X: 2, Y: 8}, Name: "Рейстлін", Health: 90, Mana: 70},
			&models.Mage{Position: models.Position{X: 4, Y: 5}, Name: "Медів", Health: 110, Mana: 40},
			&models.Mage{Position: models.Position{X: 3, Y: 9}, Name: "Джайна", Health: 95, Mana: 60},

			&models.Orc{Position: models.Position{X: 9, Y: 9}, Name: "Тралл", Health: 150, Damage: 20},
			&models.Orc{Position: models.Position{X: 8, Y: 8}, Name: "Громмаш", Health: 160, Damage: 25},
			&models.Orc{Position: models.Position{X: 9, Y: 2}, Name: "Дуротан", Health: 140, Damage: 18},
			&models.Orc{Position: models.Position{X: 7, Y: 5}, Name: "Оргрім", Health: 150, Damage: 22},
			&models.Orc{Position: models.Position{X: 8, Y: 3}, Name: "Гулдан", Health: 120, Damage: 30},
		},
	}

	fmt.Println("=== ДЕНЬ 5: ПОЛІМОРФІЗМ ТА ІНТЕРФЕЙСИ ===")

	fmt.Println("\n🎬 Початковий стан мапи:")
	renderWorld(world)

	// Симулюємо 3 покрокові ходи
	for turn := 1; turn <= 3; turn++ {
		fmt.Printf("\n➡️ Натисніть Enter для переходу до кроку %d...", turn)
		var input string
		fmt.Scanln(&input)

		world.Tick()

		fmt.Printf("\n🎨 Стан мапи після кроку №%d:\n", turn)
		renderWorld(world)
	}
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
