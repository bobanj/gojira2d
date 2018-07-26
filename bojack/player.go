package main

import (
	"github.com/go-gl/mathgl/mgl32"
	g "github.com/markov/gojira2d/pkg/graphics"
	"fmt"
	"log"
)

type Player struct {
	quad                  *g.Primitive2D
	position              mgl32.Vec3
	currentSpritePosition int
	runningSprites        []*g.Texture
}

func NewPlayer(position mgl32.Vec3, playerName string, numberOfFrames int) *Player {
	p := &Player{}
	p.runningSprites = make([]*g.Texture, 0, numberOfFrames)
	for i := 0; i <= numberOfFrames; i++ {
		spriteNumber := fmt.Sprintf("%d", i)
		if i < 10 {
			spriteNumber = fmt.Sprintf("0%d", i)
		}
		log.Printf(fmt.Sprintf("bojack/sprites/%s/%s_%s.png", playerName, playerName, spriteNumber))
		p.runningSprites = append(
			p.runningSprites,
			g.NewTextureFromFile(fmt.Sprintf("bojack/sprites/%s/%s_%s.png", playerName, playerName, spriteNumber)))

	}
	p.currentSpritePosition = 0
	p.quad = g.NewQuadPrimitive(mgl32.Vec3{position.X(), position.Y(), 0}, mgl32.Vec2{0, 0})
	p.quad.SetAnchorToCenter()
	p.quad.SetTexture(p.runningSprites[p.currentSpritePosition])
	p.quad.SetSizeFromTexture()
	p.quad.SetScale(mgl32.Vec2{0.15, 0.15})
	return p
}

func (p *Player) Update(speed float32) {
	if p.currentSpritePosition == 3 {
		p.currentSpritePosition = 1
	} else {
		p.currentSpritePosition = p.currentSpritePosition + 1
	}
	absPos := p.position
	absPos = absPos.Add(mgl32.Vec3{speed, 0, 0})
	p.position = absPos
	p.quad.SetPosition(p.position)
	p.quad.SetTexture(p.runningSprites[p.currentSpritePosition])
}

func (p *Player) Draw(ctx *g.Context) {
	p.quad.EnqueueForDrawing(ctx)
}
