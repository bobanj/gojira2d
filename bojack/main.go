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

	players := make([]*Player, 0, 32)
	players = append(players,  NewPlayer(mgl32.Vec3{40, 210, 0.001}, mgl32.Vec2{0.15, 0.15}, "bojack", 3))
	players = append(players,  NewPlayer(mgl32.Vec3{40, 400, 0.002}, mgl32.Vec2{0.17, 0.17}, "todd", 3))
	players = append(players,  NewPlayer(mgl32.Vec3{40, 580, 0.001}, mgl32.Vec2{0.25, 0.25}, "monkey", 3))
	RegisterKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		if action == glfw.Release {
			log.Printf("#%d key:", key)
		}
	})
	scene := NewScene()

	playerSpeed := float32(0.4)
	app.MainLoop(func(speed float64) {
		scene.Update(speed)
		updateHud()
		for _, player := range players {
			player.Update(playerSpeed)
		}
	}, func() {
		scene.Draw(app.Context)
		for _, player := range players {
			player.Draw(app.Context)
		}
		drawHud(app.UIContext)
	})
}
