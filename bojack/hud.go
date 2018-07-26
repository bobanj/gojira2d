package main

import (
	"github.com/go-gl/mathgl/mgl32"

	g "github.com/markov/gojira2d/pkg/graphics"
	"github.com/markov/gojira2d/pkg/ui"
)

type window struct {
	w, h int
}

var (
	win    = window{w: 800, h: 600}
	rhythm *ui.Text
	line   *g.Primitive2D
)

func createHud() {
	font := ui.NewFontFromFiles(
		"regular",
		"examples/assets/fonts/roboto-regular.fnt",
		"examples/assets/fonts/roboto-regular.png",
	)

	rhythm = ui.NewText(
		"rhythm",
		font,
		mgl32.Vec3{float32(win.w - 120), 10, -1},
		mgl32.Vec2{25, 25},
		g.Color{0.8, 0.7, 0.6, 1},
		mgl32.Vec4{0, 0, 0, -.17},
	)

	line = g.NewQuadPrimitive(
		mgl32.Vec3{float32(win.w - 220), 10, -1},
		mgl32.Vec2{100, 25},
	)
}

func drawHud(ctx *g.Context) {
	rhythm.EnqueueForDrawing(ctx)
	line.EnqueueForDrawing(ctx)
}
