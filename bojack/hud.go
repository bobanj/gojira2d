package main

import (
	"github.com/go-gl/mathgl/mgl32"
	g "github.com/markov/gojira2d/pkg/graphics"
	"math"
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
	win      = window{w: 1920, h: 1080}
	gogoQuad *g.Primitive2D
	gogoAnim float64
	track    Track

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
	track = NewTrack(win)
}

func createGoGoGo() {
	gogoQuad = g.NewQuadPrimitive(mgl32.Vec3{1920 / 2, 300, 0.1}, mgl32.Vec2{0, 0})
	gogoQuad.SetTexture(g.NewTextureFromFile("bojack/sprites/bg/gogogo.png"))
	gogoQuad.SetSizeFromTexture()
	gogoQuad.SetAnchorToCenter()
	gogoQuad.SetScale(mgl32.Vec2{0.7, 0.7})
}

func drawGoGoGo(ctx *g.Context, player *Player, scene *Scene) {
	if player.canStart && scene.X() == 0 {
		gogoQuad.EnqueueForDrawing(ctx)
	}
}

func updateHud() {
	track.Update()
	gogoAnim += 0.1
	gogoScale := 0.5 + float32(math.Abs(math.Sin(gogoAnim/2)/2))
	gogoQuad.SetScale(mgl32.Vec2{gogoScale, gogoScale})
}

func shouldPress() bool {
	return track.shouldPress()
}

func pressOpportunity() bool {
	return track.pressOpportunity()
}

func releaseOpportunity() bool {
	return track.releaseOpportunity()
}

func drawHud(ctx *g.Context) {
	track.DrawButton(ctx)
}

func drawBars(ctx *g.Context) {
	track.Draw(ctx)
}
