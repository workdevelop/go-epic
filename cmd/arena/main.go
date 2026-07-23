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

	fmt.Print("\x1b[2J")   // Очищення термінала при старті
	fmt.Print("\x1b[?25l") // Ховаємо курсор

	// 1. 🔒 ПЕРШИЙ DEFER: Гарантуємо увімкнення курсора назад при виході з програми.
	// Що б не сталося далі (успішний фінал, Ctrl+C чи паніка), Go виконає цей рядок останнім.
	defer fmt.Print("\x1b[?25h")

	ticker := time.NewTicker(200 * time.Millisecond) // Створюємо таймер на 200 мілісекунд (5 FPS)

	// 2. 🔒 ДРУГИЙ DEFER: Очищаємо ресурси процесора та пам'яті від таймера.
	// Працює за принципом LIFO: цей дефер зареєстровано другим, тому при виході
	// він виконається ПЕРШИМ (спочатку зупиниться таймер, потім увімкнеться курсор).
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
			// fmt.Print("\x1b[?25h")
			fmt.Print("\x1b[2J")
			fmt.Println("\n🔴 ОРКИ ПЕРЕМОГЛИ! ГРА ЗАВЕРШЕНА.")
			break
		}
		if liveOrcs == 0 {
			// fmt.Print("\x1b[?25h")
			fmt.Print("\x1b[2J")
			fmt.Println("\n🔵 МАГИ ПЕРЕМОГЛИ! ГРА ЗАВЕРШЕНА.")
			break
		}

		render.RenderWorld(world, turn, liveMages, liveOrcs, HealsSumMages, HealsSumOrcs)

		world.Tick()

		// 2. Етап Бою (Пошук Колізій через універсальний Battle)
		for i := 0; i < len(world.Units); i++ {
			for j := i + 1; j < len(world.Units); j++ {
				u1, u2 := world.Units[i], world.Units[j]

				// Перевірки: мертві не воюють, свої своїх не б'ють
				if !u1.IsAlive() || !u2.IsAlive() || u1.GetType() == u2.GetType() {
					continue
				}

				x1, y1 := u1.GetPosition()
				x2, y2 := u2.GetPosition()

				// Якщо юніти зустрілися на одній клітинці
				if x1 == x2 && y1 == y2 {
					// Запускаємо двосторонній бій через нашу універсальну функцію
					log1 := engine.Battle(u1, u2) // u1 б'є u2
					fmt.Println(log1)
					log2 := engine.Battle(u2, u1) // u2 б'є u1
					fmt.Println(log2)

					world.BattleLog = append(world.BattleLog, log1, log2)

					// Перевіряємо, чи хтось загинув після цієї сутички
					if !u1.IsAlive() {
						world.BattleLog = append(world.BattleLog, fmt.Sprintf("💀 %s загинув у епічній битві!", u1.GetName()))
					}
					if !u2.IsAlive() {
						world.BattleLog = append(world.BattleLog, fmt.Sprintf("💀 %s загинув у епічній битві!", u2.GetName()))
					}
				}
			}
		}

		turn++
	}
}
