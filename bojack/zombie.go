package main

import (
	"github.com/go-gl/mathgl/mgl32"
	g "github.com/markov/gojira2d/pkg/graphics"
	"fmt"
	"math"
)

type Zombie struct {
	quad                  *g.Primitive2D
	shadowQuad            *g.Primitive2D
	speed                 float32
	position              mgl32.Vec3
	runningSprites        []*g.Texture
	numberOfFrames        int
	currentSpritePosition float64
	animationSpeed        float64
	mugshotTexturePath	string
	isWinner	bool
}

func NewZombie(position mgl32.Vec3, scale mgl32.Vec2, playerName string, numberOfFrames int) *Zombie {
	zombie := &Zombie{}
	zombie.runningSprites = make([]*g.Texture, 0, numberOfFrames+1)
	for i := 0; i < numberOfFrames; i++ {
		zombie.runningSprites = append(
			zombie.runningSprites,
			g.NewTextureFromFile(fmt.Sprintf("bojack/sprites/%s/%s_%02d.png", playerName, playerName, i)))
	}
	zombie.mugshotTexturePath = fmt.Sprintf("bojack/sprites/mugshots/%s.png", playerName)
	zombie.isWinner = false
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

	zombie.shadowQuad = g.NewQuadPrimitive(position, mgl32.Vec2{0, 0})
	zombie.shadowQuad.SetTexture(g.NewTextureFromFile("bojack/sprites/shadow.png"))
	zombie.shadowQuad.SetSizeFromTexture()
	zombie.shadowQuad.SetScale(mgl32.Vec2{0.8, 0.6})
	zombie.shadowQuad.SetAnchorToCenter()
	return zombie
}

func (zombie *Zombie) Update(scene *Scene) {
	if zombie.position.X() < scene.X()-120 && scene.X() > 0 {
		zombie.position = mgl32.Vec3{scene.X() - 120, zombie.position.Y(), zombie.position.Z()}
	}

	zombie.currentSpritePosition = math.Mod(zombie.currentSpritePosition+zombie.animationSpeed, float64(zombie.numberOfFrames))
	absPos := zombie.position
	absPos = absPos.Add(mgl32.Vec3{zombie.speed, 0, 0})
	zombie.position = absPos
	zombie.quad.SetPosition(zombie.position.Sub(mgl32.Vec3{scene.X(), 0, 0}))
	zombie.shadowQuad.SetPosition(zombie.position.Sub(mgl32.Vec3{scene.X() + 40, 0, -0.05}))
	zombie.quad.SetTexture(zombie.runningSprites[int(zombie.currentSpritePosition)])

	zombie.speed = 1 + scene.X() * 0.000475
	scene.UpdateZombiePos(zombie)
}

func (zombie *Zombie) Draw(ctx *g.Context) {
	zombie.shadowQuad.Draw(ctx)
	zombie.quad.Draw(ctx)
}
