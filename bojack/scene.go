package main

import (
	"github.com/go-gl/mathgl/mgl32"
	g "github.com/markov/gojira2d/pkg/graphics"
	"github.com/markov/gojira2d/pkg/ui"
	"math/rand"
	"fmt"
	)

type Scene struct {
	position mgl32.Vec3
	quad     *g.Primitive2D
	x        float32
}

func NewScene() *Scene {
	p := &Scene{}
	p.quad = g.NewQuadPrimitive(mgl32.Vec3{0, 0, 0.4}, mgl32.Vec2{0, 0})
	p.quad.SetAnchorToCenter()
	t := g.NewTextureFromFile("bojack/sprites/bg/background.png")
	p.quad.SetTexture(t)
	p.quad.SetSizeFromTexture()
	return p
}

func (p *Scene) Update(speed float64) {
	if p.x + 1920 >= 8884 {
		p.x = 8884 - 1920
	}
	p.quad.SetPosition(mgl32.Vec3{-p.x, 0, 0.4})
}

func (p *Scene)UpdatePlayerPos(x float32) {
	if x > p.x + 1600 {
		p.x = x - 1600
	}
}

func (p *Scene)ShowAndFadeGoGoGo() {
	font := ui.NewFontFromFiles(
		"regular",
		"examples/assets/fonts/roboto-regular.fnt",
		"examples/assets/fonts/roboto-regular.png",
	)
	color := g.Color{
		rand.Float32(),
		rand.Float32(),
		rand.Float32(),
		0.6 + 0.4*rand.Float32(),
	}
	tc := ui.NewText(
		"The quick brown fox jumps over the lazy dog",
		font,
		mgl32.Vec3{0, 0 + 30, 0},
		mgl32.Vec2{10, 30},
		color,
		mgl32.Vec4{0, 0, 0, -.17},
	)
	tc.SetText(fmt.Sprintf("GO GO GO"))
}

func (p *Scene) Draw(ctx *g.Context)  {
	p.quad.EnqueueForDrawing(ctx)
}

func (p* Scene) X() float32 {
	return p.x
}