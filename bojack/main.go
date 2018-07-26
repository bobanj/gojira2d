package main

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/markov/gojira2d/pkg/app"
	"log"
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
	scene := NewScene()

	playerOne := NewPlayer(mgl32.Vec3{40, 210, 0.1}, mgl32.Vec2{0.15, 0.15}, "bojack", 3)
	playerTwo := NewPlayer(mgl32.Vec3{40, 400, 0.2}, mgl32.Vec2{0.17, 0.17}, "todd", 3)
	playerThree := NewPlayer(mgl32.Vec3{40, 580, 0.3}, mgl32.Vec2{0.25, 0.25}, "monkey", 3)
	RegisterKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		if action == glfw.Release {
			log.Printf("#%d key:", key)
		}
	})

	playerSpeed := float32(0.6)
	app.MainLoop(func(speed float64) {
		scene.Update(speed)
		updateHud()
			playerOne.Update(playerSpeed)
			playerTwo.Update(playerSpeed)
			playerThree.Update(playerSpeed)
	}, func() {
		scene.Draw(app.Context)
		playerOne.Draw(app.Context)
		playerTwo.Draw(app.Context)
		playerThree.Draw(app.Context)
		drawHud(app.UIContext)
	})
}
