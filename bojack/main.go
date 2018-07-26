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
	scene := NewScene()

	playerOne := NewPlayer(mgl32.Vec3{40, 210, 0.1}, mgl32.Vec2{0.15, 0.15}, "bojack", 3, glfw.KeyQ)
	playerTwo := NewPlayer(mgl32.Vec3{40, 400, 0.2}, mgl32.Vec2{0.17, 0.17}, "todd", 3, glfw.KeyB)
	playerThree := NewPlayer(mgl32.Vec3{40, 580, 0.3}, mgl32.Vec2{0.25, 0.25}, "monkey", 3, glfw.KeyP)

	//players := []*Player{
	//	NewPlayer(mgl32.Vec3{40, 940, 0.3}, mgl32.Vec2{0.35, 0.35}, "bojack", glfw.KeyQ, 3),
	//	NewPlayer(mgl32.Vec3{40, 1000, 0.2}, mgl32.Vec2{0.4, 0.4}, "monkey", glfw.KeyP, 3),
	//	NewPlayer(mgl32.Vec3{40, 1060, 0.1}, mgl32.Vec2{0.34, 0.34}, "todd", glfw.KeyB, 3),
	//}
	RegisterKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		switch action {
		case glfw.Press:
			switch key {
			case playerOne.key:
				playerOne.keyPressed = true
			case playerTwo.key:
				playerTwo.keyPressed = true
			case playerThree.key:
				playerThree.keyPressed = true
			}
		case glfw.Release:

			switch key {
			case playerOne.key:
				playerOne.keyPressed = false
			case playerTwo.key:
				playerTwo.keyPressed = false
			case playerThree.key:
				playerThree.keyPressed = false
			}
		}
	})

	app.MainLoop(func(speed float64) {
		scene.Update(speed)
		updateHud()
		playerOne.Update()
		playerTwo.Update()
		playerThree.Update()
	}, func() {
		scene.Draw(app.Context)
		playerOne.Draw(app.Context)
		playerTwo.Draw(app.Context)
		playerThree.Draw(app.Context)
		drawHud(app.UIContext)
	})
}
