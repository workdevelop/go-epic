package main

import (
	"database/sql" // Стандартний пакет для роботи з SQL БД
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go-epic/internal/engine"
	"go-epic/internal/models"

	_ "modernc.org/sqlite" // Драйвер SQLite (імпорт чисто для реєстрації)
)

// MatchHistoryDTO описує структуру для виведення історії в JSON
type MatchHistoryDTO struct {
	ID         int       `json:"id"`
	PlayedAt   time.Time `json:"played_at"`
	Winner     string    `json:"winner"`
	TurnsTotal int       `json:"turns_total"`
}

func main() {
	// 1. 🔥 ІНІЦІАЛІЗАЦІЯ БАЗИ ДАНИХ (Пул з'єднань створюється автоматично)
	db, err := sql.Open("sqlite", "arena.db")
	if err != nil {
		log.Fatalf("Неможливо відкрити базу даних: %v", err)
	}
	// Гарантуємо закриття конектів до бази при завершенні роботи сервера
	defer db.Close()

	// 2. СТВОРЮЄМО ТАБЛИЦЮ (Авто-міграція при старті)
	query := `
	CREATE TABLE IF NOT EXISTS match_history (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		played_at DATETIME,
		winner TEXT,
		turns_total INTEGER
	);`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatalf("Помилка створення таблиці: %v", err)
	}

	world := models.NewWorld(20, 20)

	// Заселяємо 5 магів та 5 орків
	for i := 1; i <= 5; i++ {
		world.Units = append(world.Units, &models.Mage{Position: models.Position{X: i + 1, Y: 2}, Name: fmt.Sprintf("Mage-%d", i), Health: 110, Mana: 50})
		world.Units = append(world.Units, &models.Orc{Position: models.Position{X: i + 1, Y: 5}, Name: fmt.Sprintf("Orc-%d", i), Health: 120, Damage: 15})
	}

	for i := 0; i < len(world.Units); i++ {
		go world.Units[i].Brain(i, world.MoveChannel)
	}

	// Змінна для фіксації фіналу, щоб не записувати гру в БД по 100 разів
	var gameSaved = false

	// Фонова горутина для прорахунку боїв
	go func() {
		turn := 1
		for event := range world.MoveChannel {
			world.Lock()

			activeUnit := world.Units[event.UnitIndex]
			if activeUnit.IsAlive() {
				_ = activeUnit.Move(event.DX, event.DY, world.Width, world.Height)
			}

			// Підрахунок живих
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

			// 🛢️ ЛОГІКА ФІНАЛУ: Якщо одна з армій вимерла і матч ще не збережено
			if (liveMages == 0 || liveOrcs == 0) && !gameSaved {
				gameSaved = true
				winner := "Орки"
				if liveOrcs == 0 {
					winner = "Маги"
				}

				// Записуємо фінал у базу через параметризований SQL-запит (захист від SQL-ін'єкцій)
				insertQuery := `INSERT INTO match_history (played_at, winner, turns_total) VALUES (?, ?, ?)`
				_, dbErr := db.Exec(insertQuery, time.Now(), winner, turn)
				if dbErr != nil {
					log.Printf("Помилка збереження матчу в БД: %v", dbErr)
				} else {
					log.Printf("💾 Матч завершено! Переможець: %s. Результат успішно збережено в БД SQLite.", winner)
				}
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
					}
				}
			}

			world.Unlock()
			turn++
		}
	}()

	// Ендпоінт стану гри (з Дня 22)
	http.HandleFunc("/state", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(world)
	})

	// 🛠️ НОВИЙ ЕНДПОІНТ: /history (Зчитування з бази даних)
	http.HandleFunc("/history", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Виконуємо SELECT-запит
		rows, qErr := db.Query("SELECT id, played_at, winner, turns_total FROM match_history ORDER BY id DESC")
		if qErr != nil {
			http.Error(w, "Помилка читання бази даних", http.StatusInternalServerError)
			return
		}
		defer rows.Close() // Обов'язково закриваємо курсор rows, щоб не було витоку дескрипторів

		var history []MatchHistoryDTO

		// Перебираємо рядки з бази
		for rows.Next() {
			var h MatchHistoryDTO
			// Явно скануємо колонки у підготовлені вказівники структури
			scanErr := rows.Scan(&h.ID, &h.PlayedAt, &h.Winner, &h.TurnsTotal)
			if scanErr != nil {
				log.Printf("Помилка сканування рядка: %v", scanErr)
				continue
			}
			history = append(history, h)
		}

		// Віддаємо масив історії в JSON
		_ = json.NewEncoder(w).Encode(history)
	})

	fmt.Println("🌐 HTTP API Сервер запущено на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
