package main

import (
	"github.com/go-gl/mathgl/mgl32"

	g "github.com/markov/gojira2d/pkg/graphics"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type window struct {
	w, h int
}

var (
	win            = window{w: 800, h: 600}
	buttonPressed  *g.Primitive2D
	buttonReleased *g.Primitive2D
)

func createHud() {
	buttonPressed = g.NewQuadPrimitive(
		mgl32.Vec3{float32(win.w - 220), 40, -1},
		mgl32.Vec2{48, 40},
	)
	buttonPressed.SetTexture(g.NewTextureFromFile("bojack/sprites/button/button_pressed.png"))

	buttonReleased = g.NewQuadPrimitive(
		mgl32.Vec3{float32(win.w - 220), 40, -1},
		mgl32.Vec2{48, 40},
	)
	buttonReleased.SetTexture(g.NewTextureFromFile("bojack/sprites/button/button_unpressed.png"))
}

func updateHud()  {
	
}

func drawHud(ctx *g.Context) {
	if int32(glfw.GetTime()) % 2 == 0 {
		buttonPressed.EnqueueForDrawing(ctx)
	} else {
		buttonReleased.EnqueueForDrawing(ctx)
	}
}
