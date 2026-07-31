package main

import (
	"context" // Пакет для керування контекстом виконання та таймаутами сервера
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal" // Перехоплювач системних сигналів
	"syscall"
	"time"

	"go-epic/internal/engine"
	"go-epic/internal/models"

	_ "modernc.org/sqlite"
)

type MatchHistoryDTO struct {
	ID         int       `json:"id"`
	PlayedAt   time.Time `json:"played_at"`
	Winner     string    `json:"winner"`
	TurnsTotal int       `json:"turns_total"`
}

// SpawnRequest описує JSON-структуру для створення нового юніта
type SpawnRequest struct {
	Type   string `json:"type"` // "mage" або "orc"
	Name   string `json:"name"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Health int    `json:"health"`
}

func main() {
	// Ініціалізація бази даних SQLite
	db, err := sql.Open("sqlite", "arena.db")
	if err != nil {
		log.Fatalf("Неможливо відкрити базу даних: %v", err)
	}

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

	// Початковий спавн 5 магів та 5 орків
	for i := 1; i <= 5; i++ {
		world.Units = append(world.Units, &models.Mage{Position: models.Position{X: i + 1, Y: 2}, Name: fmt.Sprintf("Mage-%d", i), Health: 100, Mana: 50})
		world.Units = append(world.Units, &models.Orc{Position: models.Position{X: i + 1, Y: 5}, Name: fmt.Sprintf("Orc-%d", i), Health: 120, Damage: 15})
	}

	// Запуск початкових горутин-мозків юнітів
	for i := 0; i < len(world.Units); i++ {
		go world.Units[i].Brain(i, world.MoveChannel)
	}

	var gameSaved = false

	// Фонова горутина для прорахунку боїв та читання каналу
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

			if (liveMages == 0 || liveOrcs == 0) && !gameSaved && len(world.Units) > 0 {
				gameSaved = true
				winner := "Орки"
				if liveOrcs == 0 {
					winner = "Маги"
				}

				insertQuery := `INSERT INTO match_history (played_at, winner, turns_total) VALUES (?, ?, ?)`
				_, dbErr := db.Exec(insertQuery, time.Now(), winner, turn)
				if dbErr != nil {
					log.Printf("Помилка збереження матчу: %v", dbErr)
				} else {
					log.Printf("💾 Матч завершено! Переможець: %s збережений в БД.", winner)
				}
			}

			// Прораховуємо колізії
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

	// --- МАРШРУТИ REST API ---

	// 1. GET /api/state — Отримання стану гри
	http.HandleFunc("/api/state", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Метод не підтримується", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(world)
	})

	// 2. GET /api/history — Отримання історії матчів
	http.HandleFunc("/api/history", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Метод не підтримується", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		rows, qErr := db.Query("SELECT id, played_at, winner, turns_total FROM match_history ORDER BY id DESC")
		if qErr != nil {
			http.Error(w, "Помилка БД", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var history []MatchHistoryDTO
		for rows.Next() {
			var h MatchHistoryDTO
			if err := rows.Scan(&h.ID, &h.PlayedAt, &h.Winner, &h.TurnsTotal); err == nil {
				history = append(history, h)
			}
		}
		_ = json.NewEncoder(w).Encode(history)
	})

	// 3. POST /api/spawn — Динамічне створення нового юніта через JSON-боді
	http.HandleFunc("/api/spawn", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Метод не підтримується", http.StatusMethodNotAllowed)
			return
		}

		var req SpawnRequest
		// Парсимо JSON, який прислав клієнт
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Некоректний JSON-боді", http.StatusBadRequest)
			return
		}

		world.Lock()
		defer world.Unlock()

		var newUnit models.Unit
		newPos := models.Position{X: req.X, Y: req.Y}

		if req.Type == "mage" {
			newUnit = &models.Mage{Position: newPos, Name: req.Name, Health: req.Health, Mana: 50}
		} else if req.Type == "orc" {
			newUnit = &models.Orc{Position: newPos, Name: req.Name, Health: req.Health, Damage: 15}
		} else {
			http.Error(w, "Невідомий тип юніта. Дозволено: mage, orc", http.StatusBadRequest)
			return
		}

		// Додаємо нового юніта на арену на льоту
		world.Units = append(world.Units, newUnit)
		newIndex := len(world.Units) - 1

		// 🔥 ОДРАЗУ ЗАПУСКАЄМО ДЛЯ НЬОГО ОКРЕМУ ГОРУТИНУ МОЗКУ!
		go world.Units[newIndex].Brain(newIndex, world.MoveChannel)

		// Скидаємо прапорець фіналу матчу, бо прийшов новий боєць
		gameSaved = false

		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(fmt.Sprintf(`{"message": "Юніт %s успішно спавнений у горутині №%d"}`, req.Name, newIndex)))
	})

	// 4. POST /api/reset — Повний перезапуск гри
	http.HandleFunc("/api/reset", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Метод не підтримується", http.StatusMethodNotAllowed)
			return
		}
		world.Lock()
		world.Units = []models.Unit{} // Очищуємо армію
		world.BattleLog = []string{}
		gameSaved = false
		world.Unlock()

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"message": "Арену повністю очищено й перезапущено"}`))
	})

	// Ініціалізуємо структуру HTTP-сервера для можливості його Graceful зупинки
	server := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	// 🔥 НАЛАШТУВАННЯ GRACEFUL SHUTDOWN
	// Створюємо буферизований канал для перехоплення сигналів ОС
	shutdownChan := make(chan os.Signal, 1)
	// Кажемо системі надсилати нам сповіщення про Ctrl+C (Interrupt) або SIGTERM (Kubernetes)
	signal.Notify(shutdownChan, os.Interrupt, syscall.SIGTERM)

	// Запускаємо вебсервер в ОКРЕМІЙ горутині, щоб він не блокував main
	go func() {
		fmt.Println("🌐 PROD-READY REST API Сервер запущено на http://localhost:8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Помилка запуску сервера: %v", err)
		}
	}()

	// 🛑 ГОЛОВНИЙ ПОТІК ЗАВИСАЄ ТУТ І ЧЕКАЄ НА СИГНАЛ ВИМКНЕННЯ ВІД ОС
	<-shutdownChan
	fmt.Println("\n⚠️ Отримано сигнал зупинки! Починається процес Graceful Shutdown...")

	// Створюємо контекст із таймаутом 5 секунд. Якщо за 5 сек сервер не закриється сам,
	// ОС вимкне його примусово
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 1. Зупиняємо вебсервер (він припиняє приймати нові HTTP-запити)
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Помилка чистого закриття сервера: %v", err)
	}

	// 2. Закриваємо базу даних SQLite, зберігаючи цілісність arena.db файлу
	log.Println("🛢️ Закриття пулу з'єднань SQLite...")
	db.Close()

	fmt.Println("🏁 Сервер успішно та чисто завершив свою роботу. Проєкт Mage Arena v1.0 готовий!")
}
