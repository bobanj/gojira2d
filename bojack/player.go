package main

import (
	"github.com/go-gl/mathgl/mgl32"
	g "github.com/markov/gojira2d/pkg/graphics"
	"fmt"
	"github.com/go-gl/glfw/v3.2/glfw"
	"math"
	)

var (
	playersStopAtX = float32(550)
)

type Player struct {
	quad              *g.Primitive2D
	speed             float32
	key               glfw.Key
	keyPressed        bool
	position          mgl32.Vec3
	runningSprites    []*g.Texture
	numberOfFrames    int
	currentFrameIndex int
	animationSpeed    float32
	canStart          bool
	offsetXStartLine float32
	playerName string
}

func NewPlayer(
	position mgl32.Vec3,
	scale mgl32.Vec2,
	playerName string,
	numberOfFrames int,
	key glfw.Key,
	offsetXStartLine float32) *Player {
	p := &Player{}
	p.canStart = false
	p.offsetXStartLine = offsetXStartLine
	p.runningSprites = make([]*g.Texture, 0, numberOfFrames+1)
	for i := 0; i < numberOfFrames; i++ {
		var spriteNumber string
		if i < 10 {
			spriteNumber = fmt.Sprintf("0%d", i)
		} else {
			spriteNumber = fmt.Sprintf("%d", i)
		}
		p.runningSprites = append(
			p.runningSprites,
			g.NewTextureFromFile(fmt.Sprintf("bojack/sprites/%s/%s_%s.png", playerName, playerName, spriteNumber)))

	}
	p.playerName = playerName
	p.key = key
	p.position = position
	p.numberOfFrames = numberOfFrames
	p.currentFrameIndex = 0
	p.quad = g.NewQuadPrimitive(position, mgl32.Vec2{0, 0})
	p.quad.SetTexture(p.runningSprites[p.currentFrameIndex])
	p.quad.SetSizeFromTexture()
	p.quad.SetScale(scale)
	p.quad.SetAnchorToBottomCenter()
	return p
}

func (p *Player) UpdateIntro(scene *Scene) {
	p.updateSprite(scene)
	//log.Printf("%s calc: %f:", p.playerName, p.position.X() + p.offsetXStartLine)
	if p.position.X() + p.offsetXStartLine >= playersStopAtX {
		p.canStart = true
	}
}

func (p *Player) Update(scene *Scene) {
	if !p.canStart {
		p.speed = 1.9
		p.UpdateIntro(scene)
	} else {
		p.RunRunRun(scene)
	}
}

func (p *Player) RunRunRun(scene *Scene) {
	if p.keyPressed {
		if p.speed < 9 {
			p.speed += 0.1
		}
	} else {
		p.speed /= 2
	}

	p.updateSprite(scene)
	//log.Printf("CURRENT POSITION #%d:", p.currentFrameIndex)
}

func (p *Player)updateSprite(scene *Scene) {
	p.animationSpeed += float32(math.Min(float64(p.speed), 3))
	p.currentFrameIndex = int(p.animationSpeed/10) % p.numberOfFrames
	absPos := p.position
	absPos = absPos.Add(mgl32.Vec3{p.speed, 0, 0})
	p.position = absPos
	p.quad.SetPosition(p.position.Sub(mgl32.Vec3{scene.X(), 0, 0}))
	p.quad.SetTexture(p.runningSprites[p.currentFrameIndex])
	scene.UpdatePlayerPos(p.position.X())
}

func (p *Player) Draw(ctx *g.Context) {
	p.quad.EnqueueForDrawing(ctx)
}
