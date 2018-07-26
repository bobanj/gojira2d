package main

import (
				"github.com/go-gl/mathgl/mgl32"
	g "github.com/markov/gojira2d/pkg/graphics"
	)

type Player struct {
	quad      *g.Primitive2D
	position   mgl32.Vec3
}

func NewPlayer(position mgl32.Vec3) *Player {
	p := &Player{}
	p.position = position
	p.quad = g.NewQuadPrimitive(mgl32.Vec3{position.X(), position.Y(), 0}, mgl32.Vec2{30, 30})
	p.quad.SetAnchorToCenter()
	p.quad.SetTexture(g.NewTextureFromFile("examples/assets/texture.png"))
	return p
}

func (p *Player) Update() {
	absPos := p.position
	absPos = absPos.Add(mgl32.Vec3{10, 0, 0})
	p.position = absPos
	p.quad.SetPosition(p.position)
}

func (p *Player) Draw(ctx *g.Context) {
	p.quad.EnqueueForDrawing(ctx)
}

