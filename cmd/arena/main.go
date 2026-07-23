package main

import (
	"fmt"
	"go-epic/internal/engine"
	"go-epic/internal/models"
	"go-epic/internal/render"
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

	fmt.Print("\x1b[2J")   // Очищення екрана при старті
	fmt.Print("\x1b[?25l") // Ховаємо курсор

	defer fmt.Print("\x1b[?25h") // Захист: повертаємо курсор при виході

	ticker := time.NewTicker(200 * time.Millisecond) // 5 FPS
	defer ticker.Stop()

	turn := 1

	for range ticker.C {
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
			fmt.Println("\n🔴 ОРКИ ПЕРЕМОГЛИ! ГРА ЗАВЕРШЕНА.                                                ")
			break
		}
		if liveOrcs == 0 {
			fmt.Println("\n🔵 МАГИ ПЕРЕМОГЛИ! ГРА ЗАВЕРШЕНА.                                                ")
			break
		}

		// Рендеринг нового TUI-інтерфейсу
		render.RenderWorld(world, turn, liveMages, liveOrcs)

		// Прораховуємо рух
		world.Tick()

		// Прораховуємо бої та колізії
		for i := 0; i < len(world.Units); i++ {
			for j := i + 1; j < len(world.Units); j++ {
				u1, u2 := world.Units[i], world.Units[j]
				if !u1.IsAlive() || !u2.IsAlive() || u1.GetType() == u2.GetType() {
					continue
				}

				x1, y1 := u1.GetPosition()
				x2, y2 := u2.GetPosition()

				if x1 == x2 && y1 == y2 {
					log1 := engine.Battle(u1, u2)
					log2 := engine.Battle(u2, u1)
					world.BattleLog = append(world.BattleLog, log1, log2)

					if !u1.IsAlive() {
						world.BattleLog = append(world.BattleLog, fmt.Sprintf("💀 %s загинув!", u1.GetName()))
					}
					if !u2.IsAlive() {
						world.BattleLog = append(world.BattleLog, fmt.Sprintf("💀 %s загинув!", u2.GetName()))
					}
				}
			}
		}

		turn++
	}
}
