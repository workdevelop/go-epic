package render

import (
	"fmt"
	"go-epic/internal/models"
)

// RenderWorld малює поточний кадр гри. Назва з Великої літери робить функцію публічною.
func RenderWorld(w models.World, turn, liveMages, liveOrcs, HealsSumMages, HealsSumOrcs int) {
	// Повертаємо курсор термінала в лівий верхній кут (0,0) для анімації без блимання
	fmt.Print("\x1b[H")

	// Рендеринг верхньої панелі (додаємо фіксовану ширину %-40s для безпеки)
	fmt.Printf("%-40s\n", fmt.Sprintf("=== MAGE ARENA v0.2 | КАДР №%d ===", turn))
	fmt.Printf("%-40s\n\n", fmt.Sprintf("📊 Живих Магів: %d (%d) | Живих Орків: %d (%d) Хід %d     \n\n", liveMages, HealsSumMages, liveOrcs, HealsSumOrcs, turn))

	// Малюємо ASCII-карту
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

	// 🛠️ ВИПРАВЛЕННЯ БАГУ: Додаємо %-100s, щоб пробіли змили старі хвости тексту
	fmt.Println("\n📝 ОСТАННЯ ПОДІЯ БОЮ:                                                                    ")
	if len(w.BattleLog) > 0 {
		lastLog := w.BattleLog[len(w.BattleLog)-1]
		// %-100s гарантує, що рядок розшириться пробілами вправо на 100 символів
		fmt.Printf("- %-100s\n", lastLog)
	} else {
		fmt.Printf("- %-100s\n", "На арені поки що тихо...")
	}
}
