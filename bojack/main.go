package main

import (
	"github.com/markov/gojira2d/pkg/app"
	"github.com/markov/gojira2d/pkg/graphics"
	"github.com/markov/gojira2d/pkg/ui"
	"github.com/go-gl/mathgl/mgl32"
)

type window struct {
	w, h int
}

var win = window{w: 800, h: 600}
var (
	rythm *ui.Text
)

func createUI() {
	font := ui.NewFontFromFiles(
		"regular",
		"examples/assets/fonts/roboto-regular.fnt",
		"examples/assets/fonts/roboto-regular.png",
	)

	rythm = ui.NewText(
		"rhythm",
		font,
		mgl32.Vec3{float32(win.w - 120), 10, -1},
		mgl32.Vec2{25, 25},
		graphics.Color{0.8, 0.7, 0.6, 1},
		mgl32.Vec4{0, 0, 0, -.17},
	)

}

func main() {
	app.Init(win.w, win.h, false, "Run For Your Life!")
	defer app.Terminate()
	createUI()

	app.MainLoop(func(speed float64) {
		//NOOP
	}, func() {
		rythm.EnqueueForDrawing(app.UIContext)
	})
}
