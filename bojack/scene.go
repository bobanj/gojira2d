package main

import (
	"github.com/go-gl/mathgl/mgl32"
	g "github.com/markov/gojira2d/pkg/graphics"
	)

type Scene struct {
	position mgl32.Vec3
	quad     *g.Primitive2D
	winnerQuad     *g.Primitive2D
	x        float32
	shouldShowWinner bool
}

func NewScene() *Scene {
	p := &Scene{}
	p.shouldShowWinner = false
	p.quad = g.NewQuadPrimitive(mgl32.Vec3{0, 0, 0.6}, mgl32.Vec2{0, 0})
	p.quad.SetAnchorToCenter()
	t := g.NewTextureFromFile("bojack/sprites/bg/background.png")
	p.quad.SetTexture(t)
	p.quad.SetSizeFromTexture()
	p.winnerQuad = g.NewQuadPrimitive(mgl32.Vec3{100, 100, 0.01}, mgl32.Vec2{0, 0})
	return p
}

func (p *Scene) Update(speed float64) {
	if p.x + 1920 >= 8884 {
		p.x = 8884 - 1920
	}
	p.quad.SetPosition(mgl32.Vec3{-p.x, 0, 0.4})
}

func (p *Scene)UpdatePlayerPos(player *Player) {
	if player.position.X() >= p.quad.GetSize().X() {
		p.winnerQuad.SetTexture(g.NewTextureFromFile(player.mugshotTexturePath))
		p.winnerQuad.SetSizeFromTexture()
		p.winnerQuad.SetAnchorToCenter()
		p.shouldShowWinner = true
	}
	if player.position.X() > p.x + 1600 {
		p.x = player.position.X() - 1600
	}
}

func (p *Scene) Draw(ctx *g.Context)  {
	p.quad.EnqueueForDrawing(ctx)
	if p.shouldShowWinner {
		p.winnerQuad.EnqueueForDrawing(ctx)
	}
}

func (p* Scene) X() float32 {
	return p.x
}