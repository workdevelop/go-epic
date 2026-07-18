package engine

import (
	"fmt"
	"go-epic/internal/models"
)

func Battle(attacker, victim models.Unit) {
	damage := attacker.GetDamage()
	victim.TakeDamage(damage)
	healthInfo := ""
	if !victim.IsAlive() {
		healthInfo = fmt.Sprintf("[%s] мертвий", victim.GetName())
	}
	fmt.Printf("⚔️ [%s] атакує [%s] і завдає %d шкоди%s!\n", attacker.GetName(), victim.GetName(), damage, healthInfo)
}
