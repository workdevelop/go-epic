package main

import (
	"fmt"
	"go-epic/internal/models" // Корінь 'go-epic' + шлях до папки з моделями
)

func main() {
	// 1. Створюємо ігровий світ та заселяємо його юнітами
	// Використовуємо ваше правильне анонімне вбудовування Position
	world := models.World{
		Width:  10,
		Height: 10,
		Mages: []models.Mage{
			{
				Position: models.Position{X: 2, Y: 3},
				Health:   100,
				Mana:     80,
			},
		},
		Orcs: []models.Orc{
			{
				Position: models.Position{X: 7, Y: 6},
				Health:   150,
				Damage:   25,
			},
		},
	}

	// 2. 🧪 ЕКСПЕРИМЕНТ: Зсув мислення (PHP-FPM vs Go)
	// Беремо першого мага зі світу. В PHP це було б посилання на об'єкт.
	mageCopy := world.Mages[0]
	mageCopy.Health = 999 // Змінюємо здоров'я у локальній змінній

	fmt.Println("--- Перевірка копіювання значень ---")
	// В Go структури копіюються за значенням, тому оригінал в масиві world НЕ змінився!
	fmt.Printf("Здоров'я мага у світі: %d HP\n", world.Mages[0].Health) // Виведе: 100
	fmt.Printf("Здоров'я мага у копії: %d HP\n", mageCopy.Health)       // Виведе: 999
	fmt.Println("-------------------------------------")

	// 3. 🎨 РЕНДЕРИНГ ASCII-КАРТИ (TUI)
	fmt.Println("=== MAGE ARENA v0.1 ===")

	// Цикл по висоті (Y) та ширині (X)
	for y := 0; y < world.Height; y++ {
		for x := 0; x < world.Width; x++ {

			// Малюємо стіни по периметру карти
			if x == 0 || x == world.Width-1 || y == 0 || y == world.Height-1 {
				fmt.Print("#")
				continue
			}

			// Перевіряємо, чи стоїть Маг на цих координатах
			isMage := false
			for _, m := range world.Mages {
				// Зверніть увагу: завдяки вбудовуванню пишемо m.X та m.Y напряму!
				if m.X == x && m.Y == y {
					fmt.Print("M")
					isMage = true
					break
				}
			}
			if isMage {
				continue
			}

			// Перевіряємо, чи стоїть Орк на цих координатах
			isOrc := false
			for _, o := range world.Orcs {
				if o.X == x && o.Y == y {
					fmt.Print("O")
					isOrc = true
					break
				}
			}
			if isOrc {
				continue
			}

			// Якщо клітинка порожня — малюємо крапку (землю)
			fmt.Print(".")
		}
		fmt.Println() // Перехід на новий рядок карти
	}
	fmt.Println("=======================")
}
