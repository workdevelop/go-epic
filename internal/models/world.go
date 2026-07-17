package models

type World struct {
	Width  int
	Height int
	Mages  []Mage // Зріз (список) магічних юнітів
	Orcs   []Orc  // Зріз (список) орків
}
