package main

import (
	"github.com/go-gl/mathgl/mgl32"
	g "github.com/markov/gojira2d/pkg/graphics"
)

type Player struct {
	quad     *g.Primitive2D
	position mgl32.Vec3
	currentSpritePosition int
	sprites []*g.Texture
}


func NewPlayer(position mgl32.Vec3) *Player {
	p := &Player{}
	p.sprites = make([]*g.Texture, 0, 4)
	p.sprites = append(p.sprites, g.NewTextureFromFile("bojack/sprites/bojack/bojack_00.png"))
	p.sprites = append(p.sprites, g.NewTextureFromFile("bojack/sprites/bojack/bojack_01.png"))
	p.sprites = append(p.sprites, g.NewTextureFromFile("bojack/sprites/bojack/bojack_02.png"))
	p.sprites = append(p.sprites, g.NewTextureFromFile("bojack/sprites/bojack/bojack_03.png"))
	//p.sprites[0] = g.NewTextureFromFile("bojack/sprites/bojack/bojack_00.png")
	//p.sprites[1] = g.NewTextureFromFile("bojack/sprites/bojack/bojack_01.png")
	//p.sprites[2] = g.NewTextureFromFile("bojack/sprites/bojack/bojack_02.png")
	//p.sprites[3] = g.NewTextureFromFile("bojack/sprites/bojack/bojack_03.png")
	p.currentSpritePosition = 0
	p.quad = g.NewQuadPrimitive(mgl32.Vec3{position.X(), position.Y(), 0}, mgl32.Vec2{0, 0})
	p.quad.SetAnchorToCenter()
	p.quad.SetTexture(p.sprites[p.currentSpritePosition])
	p.quad.SetSizeFromTexture()
	p.quad.SetScale(mgl32.Vec2{0.15, 0.15})
	return p
}

func (p *Player) Update() {
	if p.currentSpritePosition == 3 {
		p.currentSpritePosition = 1
	} else {
		p.currentSpritePosition = p.currentSpritePosition + 1
	}
	absPos := p.position
	absPos = absPos.Add(mgl32.Vec3{10, 0, 0})
	p.position = absPos
	p.quad.SetPosition(p.position)
	p.quad.SetTexture(p.sprites[p.currentSpritePosition])
}

func (p *Player) Draw(ctx *g.Context) {
	p.quad.EnqueueForDrawing(ctx)
}
