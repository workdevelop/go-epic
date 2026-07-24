package render

import (
	"fmt"
	"go-epic/internal/models"
)

// RenderWorld малює поточний кадр гри. Назва з Великої літери робить функцію публічною.
func RenderWorld(w *models.World, turn, liveMages, liveOrcs int) {
	// 1. Повертаємо курсор термінала в лівий верхній кут (0,0) для плавної анімації
	fmt.Print("\x1b[H")

	// 2. Малюємо верхню панель (сервісна інформація)
	fmt.Printf("%-60s\n", fmt.Sprintf("=== MAGE ARENA v0.2 | КАДР №%d ===", turn))
	fmt.Printf("%-60s\n\n", fmt.Sprintf("📊 Живих Магів: %d | Живих Орків: %d", liveMages, liveOrcs))

	// 3. ГОРИЗОНТАЛЬНИЙ LAYOUT: Рядок за рядком малюємо і карту, і бічну панель
	for y := 0; y < w.Height; y++ {

		// --- ЧАСТИНА А: Малюємо один рядок карти ---
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
				// Приклад ручного розфарбовування символів:
				if foundUnit.GetType() == 'M' {
					fmt.Printf("\x1b[34m%c\x1b[0m", foundUnit.GetType()) // Надрукує СИНЮ літеру M
				} else {
					fmt.Printf("\x1b[32m%c\x1b[0m", foundUnit.GetType()) // Надрукує ЗЕЛЕНУ літеру O
				}
			} else {
				fmt.Print(".")
			}
		}

		// --- ЧАСТИНА Б: Малюємо відступ між картою та бічною панеллю ---
		fmt.Print("   |   ")

		// --- ЧАСТИНА В: Малюємо один рядок бічної панелі (Sidebar) ---
		// Оскільки цикл мапи йде від y=0 до y=w.Height-1, ми використовуємо індекс 'y',
		// щоб послідовно діставати юнітів із масиву у кожному рядку.
		if y < len(w.Units) {
			unit := w.Units[y]
			status := "Живий"
			if !unit.IsAlive() {
				status = "МЕРТВИЙ"
			}
			// %c — символ, %-10s — ім'я, %3d — здоров'я, %-8s — статус
			// Пробіли в кінці гарантують зачистку старих артефактів тексту
			fmt.Printf("[%c] %-10s: %3d HP (%-8s)   ",
				unit.GetType(), unit.GetName(), unit.GetHealth(), status)
		} else {
			// Якщо юніти закінчилися, а карта ще малюється — просто зачищаємо рядок справа пробілами
			fmt.Print("                                          ")
		}

		fmt.Println() // Перехід на наступний комбінований рядок
	}

	// 4. Малюємо нижню панель (Лог боїв під картою)
	fmt.Println("\n📝 ОСТАННЯ ПОДІЯ БОЮ:                                                                    ")
	if len(w.BattleLog) > 0 {
		lastLog := w.BattleLog[len(w.BattleLog)-1]
		fmt.Printf("- %-100s\n", lastLog)
	} else {
		fmt.Printf("- %-100s\n", "На арені поки що тихо...")
	}
}
