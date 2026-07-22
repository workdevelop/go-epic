package main

import (
	"fmt"
	"go-epic/internal/models"
	"time"
)

func main() {
	// Ініціалізуємо світ 12х12 з героями
	world := models.World{
		Width:  12,
		Height: 12,
		Units: []models.Unit{
			&models.Mage{Position: models.Position{X: 2, Y: 2}, Name: "Гендальф", Health: 210, Mana: 60},
			&models.Mage{Position: models.Position{X: 3, Y: 3}, Name: "Саруман", Health: 130, Mana: 60},
			&models.Mage{Position: models.Position{X: 2, Y: 8}, Name: "Рейстлін", Health: 190, Mana: 80},
			&models.Mage{Position: models.Position{X: 4, Y: 5}, Name: "Медів", Health: 210, Mana: 50},
			&models.Mage{Position: models.Position{X: 3, Y: 9}, Name: "Джайна", Health: 195, Mana: 70},

			&models.Orc{Position: models.Position{X: 9, Y: 9}, Name: "Грязний", Health: 150, Damage: 20},
			&models.Orc{Position: models.Position{X: 8, Y: 8}, Name: "Червивий", Health: 160, Damage: 25},
			&models.Orc{Position: models.Position{X: 9, Y: 2}, Name: "Блювотний", Health: 140, Damage: 18},
			&models.Orc{Position: models.Position{X: 7, Y: 5}, Name: "Жабо-гадюк", Health: 150, Damage: 22},
			&models.Orc{Position: models.Position{X: 8, Y: 3}, Name: "Чурбано-поп", Health: 120, Damage: 30},
		},
	}

	// 1. Повністю очищаємо екран один раз при старті гри
	fmt.Print("\x1b[2J")

	// 2. Ховаємо системний курсор термінала
	fmt.Print("\x1b[?25l")

	// 3. Створюємо таймер на 200 мілісекунд (5 FPS)
	ticker := time.NewTicker(200 * time.Millisecond)

	// Переконуємось, що коли гра завершиться, ресурси таймера звільняться
	defer ticker.Stop()

	turn := 1

	for range ticker.C {
		// Перевірка умов завершення гри
		liveMages, liveOrcs := 0, 0
		HealsSumMages, HealsSumOrcs := 0, 0
		for _, u := range world.Units {
			if u.IsAlive() {
				if u.GetType() == 'M' {
					HealsSumMages += u.GetHealth()
					liveMages++
				} else {
					HealsSumOrcs += u.GetHealth()
					liveOrcs++
				}
			}
		}

		// 4. 🚀 МАГІЯ АНІМАЦІЇ: Повертаємо курсор термінала в самий верхній лівий кут (0,0)
		// Наступний вивід тексту почнеться з самого верху, плавно затираючи старий кадр
		fmt.Print("\x1b[H")

		if liveMages == 0 {
			// Перед виходом повертаємо видимість курсору термінала
			fmt.Print("\x1b[?25h")
			fmt.Println("\n🔴 ОРКИ ПЕРЕМОГЛИ! ГРА ЗАВЕРШЕНА.")
			break
		}
		if liveOrcs == 0 {
			fmt.Print("\x1b[?25h")
			fmt.Println("\n🔵 МАГИ ПЕРЕМОГЛИ! ГРА ЗАВЕРШЕНА.")
			break
		}

		// Рендеринг інтерфейсу
		fmt.Printf("=== MAGE ARENA v0.2 | КАДР №%d ===\n", turn)
		fmt.Printf("📊 Живих Магів: %d (%d) | Живих Орків: %d (%d) Хід %d     \n\n", liveMages, HealsSumMages, liveOrcs, HealsSumOrcs, turn) // Пробіли в кінці затирають старі символи

		renderWorld(world)

		// Виводимо лог останньої сутички (якщо вона була)
		fmt.Println("\n📝 ОСТАННЯ ПОДІЯ БОЮ:                                                                    ")
		if len(world.BattleLog) > 0 {
			// Беремо останній рядок логу для компактності
			lastLog := world.BattleLog[len(world.BattleLog)-1]
			fmt.Printf("- %-80s\n", lastLog)
		} else {
			fmt.Println("- На арені поки що тихо...                                                              ")
		}

		// Прораховуємо наступний крок світу
		world.Tick()
		turn++
	}
}

func renderWorld(w models.World) {
	for y := 0; y < w.Height; y++ {
		for x := 0; x < w.Width; x++ {
			if x == 0 || x == w.Width-1 || y == 0 || y == w.Height-1 {
				fmt.Print("#")
				continue
			}

			var foundUnit models.Unit
			for _, u := range w.Units {
				ux, uy := u.GetPosition()
				if ux == x && uy == y && u.IsAlive() {
					foundUnit = u
					break
				}
			}

			if foundUnit != nil {
				fmt.Printf("%c", foundUnit.GetType())
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
