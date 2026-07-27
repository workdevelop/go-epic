package engine

import (
	"fmt"
	"go-epic/internal/models"
)

func Battle(u1, u2 models.Unit) []string {
	logs := []string{}

	if !u1.IsAlive() || !u2.IsAlive() {
		return logs
	}

	dmg1 := u1.GetDamage()
	hpBefore2 := u2.GetHealth()
	u2.TakeDamage(dmg1)

	log1 := fmt.Sprintf("⚔️ [%s] атакує [%s] і завдає %d шкоди! (Було %d HP -> Стало %d HP)",
		u1.GetName(), u2.GetName(), dmg1, hpBefore2, u2.GetHealth())
	logs = append(logs, log1)

	if !u2.IsAlive() {
		return logs
	}
	dmg2 := u2.GetDamage()
	hpBefore1 := u1.GetHealth()
	u1.TakeDamage(dmg2)
	log2 := fmt.Sprintf("⚔️ [%s] відбивається і завдає [%s] %d шкоди! (Було %d HP -> Стало %d HP)",
		u2.GetName(), u1.GetName(), dmg2, hpBefore1, u1.GetHealth())
	logs = append(logs, log2)
	return logs
}
