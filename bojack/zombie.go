package main

import (
	"github.com/go-gl/mathgl/mgl32"
	g "github.com/markov/gojira2d/pkg/graphics"
	"fmt"
	"math"
)

type Zombie struct {
	quad                  *g.Primitive2D
	speed                 float32
	position              mgl32.Vec3
	runningSprites        []*g.Texture
	numberOfFrames        int
	currentSpritePosition float64
	animationSpeed        float64
}

func NewZombie(position mgl32.Vec3, scale mgl32.Vec2, playerName string, numberOfFrames int) *Zombie {
	zombie := &Zombie{}
	zombie.runningSprites = make([]*g.Texture, 0, numberOfFrames+1)
	for i := 0; i < numberOfFrames; i++ {
		zombie.runningSprites = append(
			zombie.runningSprites,
			g.NewTextureFromFile(fmt.Sprintf("bojack/sprites/%s/%s_%02d.png", playerName, playerName, i)))
	}
	zombie.speed = 1
	zombie.position = position
	zombie.numberOfFrames = numberOfFrames
	zombie.currentSpritePosition = 0
	zombie.quad = g.NewQuadPrimitive(position, mgl32.Vec2{0, 0})
	zombie.quad.SetTexture(zombie.runningSprites[0])
	zombie.quad.SetSizeFromTexture()
	zombie.quad.SetScale(scale)
	zombie.quad.SetAnchorToBottomCenter()
	zombie.animationSpeed = 0.05
	return zombie
}

func (zombie *Zombie) Update(scene *Scene)  {
	if zombie.position.X() < scene.X() - 120 {
		zombie.position = mgl32.Vec3{scene.X() - 120, zombie.position.Y(), zombie.position.Z()}
	}

	zombie.currentSpritePosition = math.Mod(zombie.currentSpritePosition+zombie.animationSpeed, float64(zombie.numberOfFrames))
	absPos := zombie.position
	absPos = absPos.Add(mgl32.Vec3{zombie.speed, 0, 0})
	zombie.position = absPos
	zombie.quad.SetPosition(zombie.position.Sub(mgl32.Vec3{scene.X(), 0, 0}))
	zombie.quad.SetTexture(zombie.runningSprites[int(zombie.currentSpritePosition)])
}

func (zombie *Zombie) Draw(ctx *g.Context) {
	zombie.quad.EnqueueForDrawing(ctx)
}
