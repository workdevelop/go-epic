package main

import (
	"go-epic/internal/engine"
	"go-epic/internal/models"
	"go-epic/internal/render"
	"time"
)

const MapWidth = 10
const MapHeight = 10

const NumMage = 5
const NumOrc = 5

var World models.World = models.World{
	Width:  MapWidth,
	Height: MapHeight,
}

func main() {
	World.Init(NumMage, NumOrc)
	e := engine.Engine{
		World: &World,
	}

	render.ClearScreen()
	render.RenderWorld(&World)

	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	i := 0
	for range ticker.C {
		i++
		e.Tick()

		render.RenderWorld(&World)
		render.RenderLogs(e.TickLog)
	}
}
