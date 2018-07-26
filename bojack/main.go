package main

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/markov/gojira2d/pkg/app"
	"log"
	"github.com/go-gl/mathgl/mgl32"
	g "github.com/markov/gojira2d/pkg/graphics"
	"github.com/markov/gojira2d/pkg/ui"
)

var (
	keyCallbackFunc glfw.KeyCallback
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
		g.Color{0.8, 0.7, 0.6, 1},
		mgl32.Vec4{0, 0, 0, -.17},
	)

}

func main() {
	app.Init(win.w, win.h, false, "Run For Your Life!")
	defer app.Terminate()
	defer UnregisterKeyCallback()
	createUI()

	player := NewPlayer(mgl32.Vec3{15, 15, 0})
	RegisterKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		if action == glfw.Release {
			log.Printf("#%d key:", key)
			player.Update()
		}
	})
	app.MainLoop(func(speed float64) {
		//NOOP
	}, func() {
		player.Draw(app.Context)
		rythm.EnqueueForDrawing(app.UIContext)
	})
}
