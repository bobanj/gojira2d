package main

import (
	"github.com/go-gl/mathgl/mgl32"
	g "github.com/markov/gojira2d/pkg/graphics"
	"github.com/go-gl/glfw/v3.2/glfw"
	"container/list"
	"math/rand"
)

type window struct {
	w, h int
}

type bar struct {
	creationTime float32
	endTime      float32
	size         float32
	quad         *g.Primitive2D
}

var (
	win              = window{w: 800, h: 600}
	buttonPressed    *g.Primitive2D
	buttonReleased   *g.Primitive2D
	bars             *list.List
	sizeInterpolator = float32(win.w - 80) / 3
	barStart         = float32(20)
	barEnd           = 3*sizeInterpolator
)

func createHud() {
	buttonPressed = g.NewQuadPrimitive(
		mgl32.Vec3{float32(win.w - 60), 30, -1},
		mgl32.Vec2{48, 40},
	)
	buttonPressed.SetTexture(g.NewTextureFromFile("bojack/sprites/button/button_pressed.png"))

	buttonReleased = g.NewQuadPrimitive(
		mgl32.Vec3{float32(win.w - 60), 30, -1},
		mgl32.Vec2{48, 40},
	)
	buttonReleased.SetTexture(g.NewTextureFromFile("bojack/sprites/button/button_unpressed.png"))
	bars = list.New()
}

func updateHud() {
	time := float32(glfw.GetTime())
	if bars.Len() == 0 || bars.Front().Value.(bar).endTime < time && rand.Int31n(100) > 95 {
		duration := rand.Float32()
		size := duration * sizeInterpolator
		newBar := bar{
			time,
			time + duration,
			size,
			g.NewQuadPrimitive(
				mgl32.Vec3{0, 10, -1},
				mgl32.Vec2{size, 60},
			),
		}
		newBar.quad.SetTexture(g.NewTextureFromFile("bojack/sprites/colors/blue.png"))
		bars.PushFront(newBar)
	}

	for e := bars.Front(); e != nil; e = e.Next() {
		bar := e.Value.(bar)
		if bar.endTime+3 < time {
			bars.Remove(e)
			continue
		}

		barX := (time-bar.creationTime)*sizeInterpolator - bar.size
		barCutOff := float32(0)
		if barX < barStart {
			barCutOff = barStart - barX
			barX = barStart
		}
		barWidth := bar.size - barCutOff
		if barX+barWidth > barEnd {
			barCutOff = barX + barWidth - barEnd
		}
		barWidth = bar.size - barCutOff
		if barCutOff > bar.size {
			barWidth = 0
		}
		bar.quad.SetSize(mgl32.Vec2{
			barWidth,
			60,
		})
		bar.quad.SetPosition(mgl32.Vec3{
			barX,
			10,
			-1,
		})
	}
}

func shouldPress() bool {
	if bars.Len() == 0 {
		return false
	}
	lastBar := bars.Back().Value.(bar)
	endTime := float32(glfw.GetTime()) - 3
	return lastBar.creationTime < endTime && lastBar.endTime > endTime
}

func drawHud(ctx *g.Context) {
	if bars.Back() == nil {
		buttonReleased.EnqueueForDrawing(ctx)
		return
	}
	for e := bars.Front(); e != nil; e = e.Next() {
		e.Value.(bar).quad.EnqueueForDrawing(ctx)
	}

	if shouldPress() {
		buttonPressed.EnqueueForDrawing(ctx)
	} else {
		buttonReleased.EnqueueForDrawing(ctx)
	}
}
