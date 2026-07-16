package main

import (
	"fmt"
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
	// Ініціалізуємо світ з одним магом та одним орком
	world := models.World{
		Width:  10,
		Height: 10,
		Mages: []models.Mage{
			{Position: models.Position{X: 2, Y: 2}, Health: 100, Mana: 50},
		},
		Orcs: []models.Orc{
			{Position: models.Position{X: 5, Y: 5}, Health: 150, Damage: 30},
		},
	}

	fmt.Println("=== ДЕНЬ 3: МЕТОДИ СТРУКТУР ТА ЕМУЛЯЦІЯ БОЮ ===")

	// Зручно дістаємо посилання (вказівники) на наших перших юнітів зі зрізів світу.
	// Знак & перед елементом масиву бере його точну адресу в пам'яті.
	mage := &world.Mages[0]
	orc := &world.Orcs[0]

	// 1. Юніти роблять кроки навколо або назустріч один одному
	fmt.Printf("Початкові позиції -> Маг: [%d,%d], Орк: [%d,%d]\n", mage.X, mage.Y, orc.X, orc.Y)

	mage.Move(1, 1) // Виклик методу структури через крапку
	orc.Move(-1, -1)

	fmt.Printf("Позиції після ходу -> Маг: [%d,%d], Орк: [%d,%d]\n", mage.X, mage.Y, orc.X, orc.Y)

	// 2. Симуляція атаки: Орк б'є Мага
	fmt.Printf("\n🪓 Орк замахнувся і б'є Мага на %dダメージ (урон)!\n", orc.Damage)
	mage.TakeDamage(orc.Damage)

	// 3. Перевірка стану за допомогою методів-геттерів
	fmt.Printf("Поточне здоров'я Мага: %d HP\n", mage.Health)
	fmt.Printf("Чи живий Маг? -> %t\n", mage.IsAlive()) // %t виводить булеве значення (true/false)

	renderWorld(world)
}
