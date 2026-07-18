package main

import (
	"fmt"
	"go-epic/internal/engine"
	"go-epic/internal/models"
	"strconv"
	s "strings"
)

const MapWidth = 10
const MapHeight = 10

const NumMage = 5
const NumOrc = 5

var World models.World = models.World{
	Width:  MapWidth,
	Height: MapHeight,
}

func renderMap() {
	var positionLabel string
	var mapLine string

	fmt.Println(s.Repeat("-", World.Width+2))

	for y := 0; y < World.Height; y++ {
		mapLine = "|"
		for x := 0; x < World.Width; x++ {
			positionLabel = " "

			for i := 0; i < len(World.Units); i++ {
				unit := World.Units[i]
				if !unit.IsAlive() {
					// не виводимо дохлятину
					continue
				}

				switch v := unit.(type) {
				case *models.Mage:
					if v.IsAt(x, y) {
						positionLabel = "M"
					}
				case *models.Orc:
					if v.IsAt(x, y) {
						positionLabel = "O"
					}
				}
			}

			mapLine = mapLine + positionLabel
		}
		mapLine = mapLine + "|"
		fmt.Println(mapLine)
	}

	fmt.Println(s.Repeat("-", World.Width+2))
}

func waitForInput() {
	fmt.Scanln()
}

func main() {
	fmt.Println("Hello, Game!")

	World.Init(NumMage, NumOrc)
	e := engine.Engine{
		World: &World,
	}

	renderMap()

	i := 0
	for {
		waitForInput()
		i++
		fmt.Println("Tick " + strconv.Itoa(i))
		e.Tick()
		renderMap()
	}
}
