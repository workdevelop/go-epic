package engine

import (
	"fmt"
	"go-epic/internal/models"
)

func Battle(u1, u2 models.Unit) string {
	dmg := u1.GetDamage()
	hpBefore := u2.GetHealth()
	u2.TakeDamage(dmg)

	return fmt.Sprintf("⚔️ [%s] атакує [%s] і завдає %d шкоди! (Було %d HP -> Стало %d HP)", u1.GetName(), u2.GetName(), dmg, hpBefore, u2.GetHealth())
}
