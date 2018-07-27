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
	s := &Scene{}
	s.shouldShowWinner = false
	s.quad = g.NewQuadPrimitive(mgl32.Vec3{0, 0, 0.6}, mgl32.Vec2{0, 0})
	s.quad.SetAnchorToCenter()
	t := g.NewTextureFromFile("bojack/sprites/bg/background.png")
	s.quad.SetTexture(t)
	s.quad.SetSizeFromTexture()
	s.winnerQuad = g.NewQuadPrimitive(mgl32.Vec3{100, 100, 0.01}, mgl32.Vec2{0, 0})
	return s
}

func (s *Scene) Update(speed float64) {
	if s.x + 1920 >= 8884 {
		s.x = 8884 - 1920
	}
	s.quad.SetPosition(mgl32.Vec3{-s.x, 0, 1.0})
}

func (s *Scene)UpdatePlayerPos(player *Player) {
	if !s.shouldShowWinner && player.position.X() >= s.quad.GetSize().X() {
		player.isWinner = true
		s.winnerQuad.SetTexture(g.NewTextureFromFile(player.mugshotTexturePath))
		s.winnerQuad.SetSizeFromTexture()
		s.winnerQuad.SetScale(mgl32.Vec2{0.6, 0.6})
		s.winnerQuad.SetAnchorToCenter()
		s.winnerQuad.SetPosition(mgl32.Vec3{float32(win.w) / 2, float32(win.h)/2, 0.01})
		s.shouldShowWinner = true
	}
	if player.position.X() > s.x + 1600 {
		s.x = player.position.X() - 1600
	}
}

func (s *Scene) Draw(ctx *g.Context)  {
	s.quad.Draw(ctx)
	if s.shouldShowWinner {
		s.winnerQuad.Draw(ctx)
	}
}

func (s * Scene) X() float32 {
	return s.x
}