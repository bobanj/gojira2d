package main

import (
	"github.com/go-gl/mathgl/mgl32"
	g "github.com/markov/gojira2d/pkg/graphics"
	"fmt"
)

type Zombie struct {
	quad                  *g.Primitive2D
	speed                 float32
	position              mgl32.Vec3
	runningSprites        []*g.Texture
	numberOfFrames        int
	currentSpritePosition int
}

func NewZombie(position mgl32.Vec3, scale mgl32.Vec2, playerName string, numberOfFrames int) *Zombie {
	zombie := &Zombie{}
	zombie.runningSprites = make([]*g.Texture, 0, numberOfFrames+1)
	for i := 0; i < numberOfFrames; i++ {
		var spriteNumber string
		if i < 10 {
			spriteNumber = fmt.Sprintf("0%d", i)
		} else {
			spriteNumber = fmt.Sprintf("%d", i)
		}
		zombie.runningSprites = append(
			zombie.runningSprites,
			g.NewTextureFromFile(fmt.Sprintf("bojack/sprites/%s/%s_%s.png", playerName, playerName, spriteNumber)))

	}
	zombie.speed = 1
	zombie.position = position
	zombie.numberOfFrames = numberOfFrames
	zombie.currentSpritePosition = 0
	zombie.quad = g.NewQuadPrimitive(position, mgl32.Vec2{0, 0})
	zombie.quad.SetTexture(zombie.runningSprites[zombie.currentSpritePosition])
	zombie.quad.SetSizeFromTexture()
	zombie.quad.SetScale(scale)
	zombie.quad.SetAnchorToBottomCenter()
	return zombie
}

func (zombie *Zombie) Update() {
	zombie.currentSpritePosition = (zombie.currentSpritePosition + 1) % zombie.numberOfFrames
	absPos := zombie.position
	absPos = absPos.Add(mgl32.Vec3{zombie.speed, 0, 0})
	zombie.position = absPos
	zombie.quad.SetPosition(zombie.position)
	zombie.quad.SetTexture(zombie.runningSprites[zombie.currentSpritePosition])
}

func (zombie *Zombie) Draw(ctx *g.Context) {
	zombie.quad.EnqueueForDrawing(ctx)
}
