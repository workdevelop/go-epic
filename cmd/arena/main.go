package main

import (
	"fmt"
	"go-epic/internal/models"
	"strconv"
	s "strings"
)

const MapWidth = 10
const MapHeight = 10

const NumMage = 5
const NumOrc = 4

const MageHealth = 100
const MageMana = 100

const OrcHealth = 200
const OrcDamage = 10

var World models.World = models.World{
	Width:  MapWidth,
	Height: MapHeight,
	Mages:  []models.Mage{},
	Orcs:   []models.Orc{},
}

func initMages() {
	mages := []models.Mage{}

	xMage := 2
	y := 2

	for i := range NumMage {
		mages = append(mages, models.Mage{
			Character: models.Character{
				Name:   "mage " + strconv.Itoa(i+1),
				Health: MageHealth,
				Position: models.Position{
					X: xMage,
					Y: y,
				},
			},
			Mana: MageMana,
		})

		y = y + 1
	}

	World.Mages = mages
}

func initOrcs() {
	orcs := []models.Orc{}

	xOrc := 7
	y := 2

	for i := range NumOrc {
		orcs = append(orcs, models.Orc{
			Character: models.Character{
				Name:   "orc " + strconv.Itoa(i+1),
				Health: MageHealth,
				Position: models.Position{
					X: xOrc,
					Y: y,
				},
			},
			Damage: OrcDamage,
		})

		y = y + 1
	}

	World.Orcs = orcs
}

func renderMap() {
	var mage models.Mage
	var orc models.Orc

	var positionLabel string
	var mapLine string

	fmt.Println(s.Repeat("-", World.Width+2))

	for y := 0; y < World.Height; y++ {
		mapLine = "|"
		for x := 0; x < World.Width; x++ {
			positionLabel = " "

			for i := 0; i < len(World.Mages); i++ {
				mage = World.Mages[i]
				if mage.IsAt(x, y) {
					positionLabel = "M"
				}
			}
			for i := 0; i < len(World.Orcs); i++ {
				orc = World.Orcs[i]
				if orc.IsAt(x, y) {
					positionLabel = "O"
				}
			}

			mapLine = mapLine + positionLabel
		}
		mapLine = mapLine + "|"
		fmt.Println(mapLine)
	}

	fmt.Println(s.Repeat("-", World.Width+2))
}

func main() {
	fmt.Println("Hello, Game!")

	initMages()
	initOrcs()

	renderMap()

	fmt.Println("Moving Mage 0...")
	World.Mages[0].Move(1, 0)
	fmt.Println("Moving Orc 0...")
	World.Orcs[0].Move(-3, 0)

	fmt.Println("Mage 0 takes damage from Orc 0...")
	World.Mages[0].TakeDamage(World.Orcs[0].Damage)
	fmt.Println("Mage 0 Health: ", World.Mages[0].Health)
	if World.Mages[0].IsAlive() {
		fmt.Println("Mage 0 is alive")
	} else {
		fmt.Println("Mage 0 is dead!")
	}

	renderMap()
}
