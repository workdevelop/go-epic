package main

import (
	"fmt"
	"go-epic/internal/engine"
	"go-epic/internal/models" // Корінь 'go-epic' + шлях до папки з моделями
)

func renderWorld(w models.World) {
	// Цикл по висоті (Y) та ширині (X)
	for y := 0; y < w.Height; y++ {
		for x := 0; x < w.Width; x++ {

			// Малюємо стіни по периметру карти
			if x == 0 || x == w.Width-1 || y == 0 || y == w.Height-1 {
				fmt.Print("#")
				continue
			}

			// Перевіряємо, чи стоїть Маг на цих координатах
			isMage := false
			for _, m := range w.Mages {
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
			for _, o := range w.Orcs {
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
}

func main() {
	// Ініціалізуємо світ 10х10 з одним магом та одним орком
	world := models.World{
		Width:  10,
		Height: 10,
		Mages: []models.Mage{
			{Position: models.Position{X: 2, Y: 2}, Health: 100, Mana: 50},
		},
		Orcs: []models.Orc{
			{Position: models.Position{X: 7, Y: 7}, Health: 150, Damage: 25},
		},
	}

	fmt.Println("=== ДЕНЬ 2: ЕКСПЕРИМЕНТИ З ВКАЗІВНИКАМИ ===")
	renderWorld(world)

	// 1. Спроба порухати орка НЕПРАВИЛЬНО (передаємо значення зрізу, тобто копію)
	engine.MoveOrcWrong(world.Orcs[0], -1, -1)
	fmt.Printf("\n"+"Після MoveOrcWrong: Орк стоїть на [%d, %d] (Зрушення не відбулося через копіювання)\n",
		world.Orcs[0].X, world.Orcs[0].Y)
	renderWorld(world)

	// 2. Порухаємо мага ПРАВИЛЬНО.
	// Знак & бере адресу пам'яті конкретного мага зі зрізу world.Mages
	engine.MoveMage(&world.Mages[0], 1, 1)

	// 3. Порухаємо орка ПРАВИЛЬНО.
	// Беремо адресу першого орка у масиві
	engine.MoveOrc(&world.Orcs[0], -1, -1)

	// 4. Рендеринг карти після правильного руху
	fmt.Println("\nКАРТА ПІСЛЯ РУХУ (Маг змістився на +1, +1; Орк на -1, -1):")
	renderWorld(world)
}
