package main

import (
	"github.com/go-gl/mathgl/mgl32"
	g "github.com/markov/gojira2d/pkg/graphics"
	"fmt"
	"github.com/go-gl/glfw/v3.2/glfw"
	"math"
)

const (
	playersStopAtX = float32(550)
	maxSpeed       = 9

	FragmentShaderPlayerDead = `
       #version 410 core

       in vec2 uv_out;
       out vec4 color;

       uniform sampler2D tex;

       void main() {
			if(texture(tex, uv_out).a != 1.0f)
			{
        		discard;
    		}
			float grayScale = dot(texture(tex, uv_out).rgb, vec3(0.299, 0.587, 0.114));
    		color = vec4(grayScale, grayScale, grayScale, 1.0);
       }
       ` + "\x00"
)

type Player struct {
	quad               *g.Primitive2D
	shadowQuad         *g.Primitive2D
	speed              float32
	key                glfw.Key
	lastKeyInteraction float64
	position           mgl32.Vec3
	runningSprites     []*g.Texture
	numberOfFrames     int
	currentFrameIndex  int
	animationSpeed     float32
	canStart           bool
	isDead             bool
	offsetXStartLine   float32
	playerName         string
	mugshotTexturePath string
	deathShader        *g.ShaderProgram
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
	p.isDead = false
	p.offsetXStartLine = offsetXStartLine
	p.runningSprites = make([]*g.Texture, 0, numberOfFrames+1)
	for i := 0; i < numberOfFrames; i++ {
		p.runningSprites = append(
			p.runningSprites,
			g.NewTextureFromFile(fmt.Sprintf("bojack/sprites/%s/%s_%02d.png", playerName, playerName, i)))
	}
	p.playerName = playerName
	p.mugshotTexturePath = fmt.Sprintf("bojack/sprites/mugshots/%s.png", playerName)
	p.speed = 1.9
	p.key = key
	p.position = position
	p.numberOfFrames = numberOfFrames
	p.currentFrameIndex = 0
	p.quad = g.NewQuadPrimitive(position, mgl32.Vec2{0, 0})
	p.quad.SetTexture(p.runningSprites[p.currentFrameIndex])
	p.quad.SetSizeFromTexture()
	p.quad.SetScale(scale)
	p.quad.SetAnchorToBottomCenter()

	p.shadowQuad = g.NewQuadPrimitive(position, mgl32.Vec2{0, 0})
	p.shadowQuad.SetTexture(g.NewTextureFromFile("bojack/sprites/shadow.png"))
	p.shadowQuad.SetSizeFromTexture()
	p.shadowQuad.SetScale(mgl32.Vec2{0.8, 0.6})
	p.shadowQuad.SetAnchorToCenter()

	p.deathShader = g.NewShaderProgram(g.VertexShaderPrimitive2D, "", FragmentShaderPlayerDead)

	return p
}

func (p *Player) Update(scene *Scene) {
	time := glfw.GetTime()
	if !p.canStart {
		if p.position.X() >= playersStopAtX {
			p.canStart = true
			p.speed = 0
		}
	} else if p.lastKeyInteraction+1 < time {
		p.slowDown()
		p.lastKeyInteraction = time
	}
	p.updateSprite(scene)
}

func (p *Player) updateSprite(scene *Scene) {
	if p.position.X() < scene.X()+100 && p.canStart == true && !p.isDead {
		p.position = mgl32.Vec3{scene.X() + 100, p.position.Y(), p.position.Z()}
		p.speed = 0.8
	}
	p.animationSpeed += float32(math.Min(float64(p.speed), 3))
	p.currentFrameIndex = int(p.animationSpeed/10) % p.numberOfFrames
	absPos := p.position
	absPos = absPos.Add(mgl32.Vec3{p.speed, 0, 0})
	p.position = absPos
	p.quad.SetPosition(p.position.Sub(mgl32.Vec3{scene.X(), 0, 0}))
	p.shadowQuad.SetPosition(p.position.Sub(mgl32.Vec3{scene.X(), 0, -0.05}))
	scene.UpdatePlayerPos(p)
	if p.isDead {
		p.quad.SetShader(p.deathShader)
		return
	}

	p.quad.SetTexture(p.runningSprites[p.currentFrameIndex])
}

func (p *Player) speedUp() {
	if !p.canStart || p.isDead {
		return
	}
	p.speed += 1.5
	if p.speed > maxSpeed {
		p.speed = maxSpeed
	}
}

func (p *Player) slowDown() {
	if !p.canStart {
		return
	}
	p.speed *= 0.75
}

func (p *Player) Draw(ctx *g.Context) {
	p.quad.Draw(ctx)
	p.shadowQuad.Draw(ctx)
}
