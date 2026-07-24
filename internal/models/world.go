package models

import (
	"sync"
)

type World struct {
	Width     int
	Height    int
	Units     []Unit
	BattleLog []string
	mu        sync.Mutex
}

func (w *World) AddBattleLog(msg string) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.BattleLog = append(w.BattleLog, msg)
}

func (w *World) GetBattleLog() []string {
	w.mu.Lock()
	defer w.mu.Unlock()

	logCopy := make([]string, len(w.BattleLog))
	copy(logCopy, w.BattleLog)
	return logCopy
}

// Tick симулює один крок ігрового часу.
// Світ перебирає всіх юнітів і змушує їх зробити хід.
func (w *World) Tick() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.BattleLog = []string{}
}

func (w *World) Lock() {
	w.mu.Lock()
}

func (w *World) Unlock() {
	w.mu.Unlock()
}
