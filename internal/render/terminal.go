package render

import (
	"fmt"
	"go-epic/internal/models"
	s "strings"
)

func resetCursor() {
	fmt.Print("\x1b[H")
}

func HideCursor() {
	fmt.Print("\x1b[?25l")
}

func ShowCursor() {
	fmt.Print("\x1b[?25h")
}

func ClearScreen() {
	fmt.Print("\x1b[2J")
}

func renderMap(w *models.World) {
	var positionLabel string
	var mapLine string

	fmt.Println(s.Repeat("-", w.Width+2))

	for y := 0; y < w.Height; y++ {
		mapLine = "|"
		for x := 0; x < w.Width; x++ {
			positionLabel = " "

			for i := 0; i < len(w.Units); i++ {
				unit := w.Units[i]
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

	fmt.Println(s.Repeat("-", w.Width+2))
}

func RenderLogs(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
	}
}

func renderAliveUnits(w *models.World) {
	fmt.Println()
	fmt.Println("=========================")
	fmt.Println("==== Still in battle ====")
	fmt.Println("=========================")
	fmt.Printf("%15s| %4s   |\n", "Name", "Pos")
	fmt.Printf("---------------|--------|\n")

	for i := 0; i < len(w.Units); i++ {
		if !w.Units[i].IsAlive() {
			continue
		}
		posX, posY := w.Units[i].GetPosition()
		fmt.Printf("%15s| %2d %2d  |\n", w.Units[i].GetName(), posX, posY)
	}
}

func RenderWorld(w *models.World) {
	resetCursor()
	renderMap(w)
	renderAliveUnits(w)
}
