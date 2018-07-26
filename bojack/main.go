package main

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/markov/gojira2d/pkg/app"
	"github.com/go-gl/mathgl/mgl32"
)

var (
	keyCallbackFunc glfw.KeyCallback
)

func main() {
	app.Init(win.w, win.h, false, "Run For Your Life!", false)
	defer app.Terminate()
	defer UnregisterKeyCallback()
	createHud()
	scene := NewScene()

	players := []*Player{
		NewPlayer(mgl32.Vec3{40, 940, 0.3}, mgl32.Vec2{0.35, 0.35}, "bojack", 3, glfw.KeyQ),
		NewPlayer(mgl32.Vec3{40, 1000, 0.2}, mgl32.Vec2{0.4, 0.4}, "monkey", 3, glfw.KeyP),
		NewPlayer(mgl32.Vec3{40, 1060, 0.1}, mgl32.Vec2{0.34, 0.34}, "todd", 3, glfw.KeyB),
	}
	RegisterKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		switch action {
		case glfw.Press:
			switch key {
			case players[0].key:
				players[0].keyPressed = true
			case players[1].key:
				players[1].keyPressed = true
			case players[2].key:
				players[2].keyPressed = true
			}
		case glfw.Release:

			switch key {
			case players[0].key:
				players[0].keyPressed = false
			case players[1].key:
				players[1].keyPressed = false
			case players[2].key:
				players[2].keyPressed = false
			}
		}
	})

	app.MainLoop(func(speed float64) {
		scene.Update(speed)
		updateHud()
		for _, player := range players {
			player.Update(scene)
		}
	}, func() {
		scene.Draw(app.Context)
		for _, player := range players {
			player.Draw(app.Context)
		}
		drawHud(app.UIContext)
	})
}
