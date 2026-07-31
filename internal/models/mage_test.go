package models

import (
	"testing"
)

// TestMage_TakeDamage перевіряє механіку зменшення здоров'я
func TestMage_TakeDamage(t *testing.T) {
	// Описуємо тестову таблицю
	tests := []struct {
		name           string
		initialHealth  int
		damageAmount   int
		expectedHealth int
	}{
		{"Звичайне отримання урону", 100, 30, 70},
		{"Урон дорівнює поточному HP", 50, 50, 0},
		{"Захист від від'ємного HP (оверкіл)", 20, 100, 0},
	}

	// Запускаємо цикл по таблиці
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mage := &Mage{Health: tt.initialHealth}
			mage.TakeDamage(tt.damageAmount)

			if mage.GetHealth() != tt.expectedHealth {
				t.Errorf("TakeDamage() для '%s' зламався: отримали %d HP, а очікували %d HP",
					tt.name, mage.GetHealth(), tt.expectedHealth)
			}
		})
	}
}

// TestMage_Move перевіряє валідацію ходів та повернення помилок
func TestMage_Move(t *testing.T) {
	tests := []struct {
		name           string
		startX, startY int
		dx, dy         int
		expectError    bool
	}{
		{"Легальний крок вперед", 5, 5, 1, 0, false},
		{"Врізання в ліву стіну (X=0)", 1, 5, -1, 0, true},
		{"Врізання в нижню стіну (Y=11)", 5, 10, 0, 1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mage := &Mage{Position: Position{X: tt.startX, Y: tt.startY}, Health: 100}

			// Розмір карти фіксований 12х12
			err := mage.Move(tt.dx, tt.dy, 12, 12)

			// Перевіряємо наявність помилки за допомогою звичайного if
			if (err != nil) != tt.expectError {
				t.Errorf("Move() для '%s' повернув помилку: %v, але ми очікували статус помилки: %t",
					tt.name, err, tt.expectError)
			}
		})
	}
}
