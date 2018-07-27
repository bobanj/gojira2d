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
	createGoGoGo()
	scene := NewScene()

	players := []*Player{
		NewPlayer(
			mgl32.Vec3{-200, 940, 0.3},
			mgl32.Vec2{0.35, 0.35},
			"bojack",
			4,
			glfw.KeyQ,
			0,
			),
		NewPlayer(
			mgl32.Vec3{-350, 1000, 0.2},
			mgl32.Vec2{0.4, 0.4},
			"monkey",
			4,
			glfw.KeyP,
			150,
			),
		NewPlayer(
			mgl32.Vec3{-500, 1060, 0.1},
			mgl32.Vec2{0.34, 0.34},
			"todd",
			4,
			glfw.KeyB,
			300,
			),
	}
	zombies := []*Zombie{
		NewZombie(mgl32.Vec3{-500, 970, 0.25}, mgl32.Vec2{0.33, 0.33}, "zombie", 3),
		NewZombie(mgl32.Vec3{-650, 1030, 0.15}, mgl32.Vec2{0.32, 0.32}, "other_zombie", 3),
	}

	RegisterKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		HandleKeyPress(key, action, players)
	})

	app.MainLoop(func(speed float64) {
		scene.Update(speed)
		updateHud()
		for _, player := range players {
			player.Update(scene)
		}
		for _, zombie := range zombies {
			zombie.Update(scene)
		}

	}, func() {
		scene.Draw(app.Context)
		for _, player := range players {
			player.Draw(app.Context)
		}
		for _, zombie := range zombies {
			zombie.Draw(app.Context)
		}
		drawHud(app.UIContext)
		drawGoGoGo(app.UIContext, players[2], scene)
	})
}
