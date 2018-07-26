package main

import (
	"github.com/go-gl/mathgl/mgl32"
	g "github.com/markov/gojira2d/pkg/graphics"
)

type Scene struct {
	position              mgl32.Vec3
	quad                  *g.Primitive2D
}

func NewScene() *Scene {
	p := &Scene{}
	p.quad = g.NewQuadPrimitive(mgl32.Vec3{0,0, 0}, mgl32.Vec2{0, 0})
	p.quad.SetAnchorToCenter()
	t := g.NewTextureFromFile("bojack/sprites/bg/background.png")
	p.quad.SetTexture(t)
	p.quad.SetSizeFromTexture()
	return p
}

func (p *Scene) Update(speed float64) {
}

func (p *Scene) Draw(ctx *g.Context) {
	p.quad.EnqueueForDrawing(ctx)
}
