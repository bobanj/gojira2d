package main

import (
	"github.com/go-gl/mathgl/mgl32"
	g "github.com/markov/gojira2d/pkg/graphics"
	"github.com/go-gl/glfw/v3.2/glfw"
	"container/list"
	"math/rand"
	"math"
)

const (
	windowOfOpportunity = 0.2
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
	win              = window{w: 1920, h: 1080}
	buttonPressed    *g.Primitive2D
	buttonReleased   *g.Primitive2D
	bars             *list.List
	sizeInterpolator = float32(win.w-80) / 3
	barStart         = float32(0)
	barHeight        = float32(274)
	barEnd           = 3 * sizeInterpolator
	gogoQuad         *g.Primitive2D
	gogoAnim         float64

	FragmentShaderTexture = `
       #version 410 core

       in vec2 uv_out;
       out vec4 color;

       uniform sampler2D tex;

       void main() {
           color = texture(tex, uv_out);
            color.a = color.a * 0.5;
       }
       ` + "\x00"
)

func createHud() {
	buttonPressed = g.NewQuadPrimitive(
		mgl32.Vec3{barEnd - 48, 1080 - barHeight/2 - 40, -1},
		mgl32.Vec2{96, 80},
	)
	buttonPressed.SetTexture(g.NewTextureFromFile("bojack/sprites/button/button_pressed.png"))

	buttonReleased = g.NewQuadPrimitive(
		mgl32.Vec3{barEnd - 48, 1080 - barHeight/2 - 40, -1},
		mgl32.Vec2{96, 80},
	)
	buttonReleased.SetTexture(g.NewTextureFromFile("bojack/sprites/button/button_unpressed.png"))
	bars = list.New()
}

func createGoGoGo() {
	gogoQuad = g.NewQuadPrimitive(mgl32.Vec3{1920/2, 300, 0.1}, mgl32.Vec2{0, 0})
	gogoQuad.SetTexture(g.NewTextureFromFile("bojack/sprites/bg/gogogo.png"))
	gogoQuad.SetSizeFromTexture()
	gogoQuad.SetAnchorToCenter()
	gogoQuad.SetScale(mgl32.Vec2{0.7, 0.7})
}

func drawGoGoGo(ctx *g.Context, player *Player, scene *Scene) {
	//gogoQuad.EnqueueForDrawing(ctx)
	if player.canStart && scene.X() == 0 {
		gogoQuad.EnqueueForDrawing(ctx)
	}
}

func updateHud() {
	time := float32(glfw.GetTime())
	if bars.Len() == 0 || bars.Front().Value.(bar).endTime < time-windowOfOpportunity && rand.Int31n(100) > 92 {
		duration := rand.Float32()*(0.9-windowOfOpportunity) + windowOfOpportunity
		size := duration * sizeInterpolator
		newBar := bar{
			time,
			time + duration,
			size,
			g.NewQuadPrimitive(
				mgl32.Vec3{0, 10, 0.6},
				mgl32.Vec2{size, barHeight},
			),
		}
		newBar.quad.SetTexture(g.NewTextureFromFile("bojack/sprites/colors/blue.png"))
		newBar.quad.SetShader(g.NewShaderProgram(g.VertexShaderPrimitive2D, "", FragmentShaderTexture))
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
			barHeight,
		})
		bar.quad.SetPosition(mgl32.Vec3{
			barX,
			1080 - barHeight,
			0.6,
		})
	}
	gogoAnim += 0.1
	gogoScale := 0.5 + float32(math.Abs( math.Sin(gogoAnim/2)/2))
	gogoQuad.SetScale(mgl32.Vec2{gogoScale, gogoScale})

}

func shouldPress() bool {
	if bars.Len() == 0 {
		return false
	}
	lastBar := bars.Back().Value.(bar)
	endTime := float32(glfw.GetTime()) - 3
	return lastBar.creationTime < endTime && lastBar.endTime > endTime
}

func pressOpportunity() bool {
	if bars.Len() == 0 {
		return false
	}
	lastBar := bars.Back().Value.(bar)
	endTime := float32(glfw.GetTime()) - 3
	return mgl32.Abs(lastBar.creationTime-endTime) < windowOfOpportunity
}

func releaseOpportunity() bool {
	if bars.Len() == 0 {
		return false
	}
	lastBar := bars.Back().Value.(bar)
	endTime := float32(glfw.GetTime()) - 3
	return mgl32.Abs(lastBar.endTime-endTime) < windowOfOpportunity
}

func drawHud(ctx *g.Context) {
	if bars.Back() == nil {
		buttonReleased.EnqueueForDrawing(ctx)
		return
	}

	if shouldPress() {
		buttonPressed.EnqueueForDrawing(ctx)
	} else {
		buttonReleased.EnqueueForDrawing(ctx)
	}
}

func drawBars(ctx *g.Context) {
	for e := bars.Front(); e != nil; e = e.Next() {
		e.Value.(bar).quad.EnqueueForDrawing(ctx)
	}
}
