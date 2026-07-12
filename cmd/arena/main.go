package main

import (
	"fmt"
	"go-epic/internal/models" // Корінь 'go-epic' + шлях до папки з моделями
)

func main() {
	// Створюємо мага за вашою структурою
	hero := models.Mage{
		Position: models.Position{X: 3, Y: 4}, // Назва поля відповідає назві типу
		Health:   100,
		Mana:     50,
	}

	// А ось читаємо координати вже НАПРЯМУ, без слова Position!
	fmt.Printf("Маг з'явився на клітинці [%d, %d]\n", hero.X, hero.Y)
}
