package main

import (
	"encoding/json"
	"fmt"
	"go-epic/internal/engine"
	"go-epic/internal/models"
	"log"
	"net/http" // Вбудований HTTP-пакет Go
)

func main() {
	// Створюємо великий світ (код ініціалізації з 18-го дня — автогенератор армії)
	world := models.NewWorld(20, 20)

	// Згенеруємо для тесту 5 магів та 5 орків
	for i := 1; i <= 5; i++ {
		world.Units = append(world.Units, &models.Mage{Position: models.Position{X: i + 1, Y: 2}, Name: fmt.Sprintf("Mage-%d", i), Health: 100, Mana: 50})
		world.Units = append(world.Units, &models.Orc{Position: models.Position{X: i + 1, Y: 5}, Name: "Orc-", Health: 120, Damage: 15})
	}

	// Запускаємо фонові горутини-мізки юнітів (вони живуть паралельно у фоні RAM)
	for i := 0; i < len(world.Units); i++ {
		go world.Units[i].Brain(i, world.MoveChannel)
	}

	// Фонова горутина для прорахунку боїв та читання черги ходів з каналу
	// Фонова горутина для прорахунку боїв та читання черги ходів з каналу
	go func() {
		for event := range world.MoveChannel {
			// 🔥 БЛОКУЄМО СВІТ ДЛЯ ПРОРАХУНКУ КАДРУ ПОДІЇ
			world.Lock()

			activeUnit := world.Units[event.UnitIndex]
			if activeUnit.IsAlive() {
				_ = activeUnit.Move(event.DX, event.DY, world.Width, world.Height)
			}

			// Прораховуємо сутички
			for i := 0; i < len(world.Units); i++ {
				for j := i + 1; j < len(world.Units); j++ {
					u1, u2 := world.Units[i], world.Units[j]
					if !u1.IsAlive() || !u2.IsAlive() || u1.GetType() == u2.GetType() {
						continue
					}
					x1, y1 := u1.GetPosition()
					x2, y2 := u2.GetPosition()

					if x1 == x2 && y1 == y2 {
						combatLogs := engine.Battle(u1, u2)
						world.BattleLog = append(world.BattleLog, combatLogs...)

						if !u1.IsAlive() {
							world.BattleLog = append(world.BattleLog, fmt.Sprintf("💀 %s загинув!", u1.GetName()))
						}
						if !u2.IsAlive() {
							world.BattleLog = append(world.BattleLog, fmt.Sprintf("💀 %s загинув!", u2.GetName()))
						}
					}
				}
			}

			// 🔥 ВІДПУСКАЄМО ЗАМОК: тепер HTTP-сервер може безпечно зняти Snapshot
			world.Unlock()
		}
	}()

	// 🛠️ СТВОРЮЄМО HTTP API ЕНДПОІНТ
	// http.HandleFunc реєструє маршрут та функцію-хендлер для його обробки
	http.HandleFunc("/state", func(w http.ResponseWriter, r *http.Request) {
		// 1. Встановлюємо HTTP-заголовок (Content-Type: application/json)
		w.Header().Set("Content-Type", "application/json")

		// Статус відповіді 200 OK
		w.WriteHeader(http.StatusOK)

		// 2. Серіалізуємо об'єкт world у JSON та випліскуємо його прямо в HTTP-потік відповіді клієнту
		err := json.NewEncoder(w).Encode(world)
		if err != nil {
			http.Error(w, "Критична помилка серіалізації", http.StatusInternalServerError)
		}
	})

	fmt.Println("🌐 HTTP API Сервер успішно запущено на http://localhost:8080")
	fmt.Println("Зробіть GET-запит на http://localhost:8080/state для перевірки стану гри...")

	// ListenAndServe запускає вебсервер. Цей рядок блокує головний потік main — сервер працюватиме безкінечно
	log.Fatal(http.ListenAndServe(":8080", nil))
}
