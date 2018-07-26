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
	app.Init(win.w, win.h, false, "Run For Your Life!")
	defer app.Terminate()
	defer UnregisterKeyCallback()
	createHud()

	players := []*Player{
		NewPlayer(mgl32.Vec3{40, 210, 0.001}, mgl32.Vec2{0.15, 0.15}, "bojack", glfw.KeyQ, 3),
		NewPlayer(mgl32.Vec3{40, 400, 0.002}, mgl32.Vec2{0.17, 0.17}, "todd", glfw.KeyB, 3),
		NewPlayer(mgl32.Vec3{40, 580, 0.001}, mgl32.Vec2{0.25, 0.25}, "monkey", glfw.KeyP, 3),
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
	scene := NewScene()

	app.MainLoop(func(speed float64) {
		scene.Update(speed)
		updateHud()
		for _, player := range players {
			if player.keyPressed && shouldPress() {
				player.speed += 0.1
			} else {
				player.speed /= 2
			}
			player.Update()
		}
	}, func() {
		scene.Draw(app.Context)
		for _, player := range players {
			player.Draw(app.Context)
		}
		drawHud(app.UIContext)
	})
}
