package main

import (
	g "github.com/markov/gojira2d/pkg/graphics"
	"container/list"
	"math/rand"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Track struct {
	buttonPressed       *g.Primitive2D
	buttonReleased      *g.Primitive2D
	creationTime        float32
	endTime             float32
	quad                *g.Primitive2D
	bars                *list.List
	win                 window
	barHeight           float32
	barStart            float32
	barEnd              float32
	bottomOffset        float32
	sizeInterpolator    float32
	windowOfOpportunity float32
}

func (track *Track) Update() {
	time := float32(glfw.GetTime())
	if track.bars.Len() == 0 || track.bars.Front().Value.(bar).endTime < time-track.windowOfOpportunity && rand.Int31n(100) > 92 {
		duration := rand.Float32()*(0.9-track.windowOfOpportunity) + track.windowOfOpportunity
		size := duration * track.sizeInterpolator
		newBar := bar{
			time,
			time + duration,
			size,
			g.NewQuadPrimitive(
				mgl32.Vec3{0, 10, 0.6},
				mgl32.Vec2{size, track.barHeight},
			),
		}
		newBar.quad.SetTexture(g.NewTextureFromFile("bojack/sprites/colors/blue.png"))
		newBar.quad.SetShader(g.NewShaderProgram(g.VertexShaderPrimitive2D, "", FragmentShaderTexture))
		track.bars.PushFront(newBar)
	}

	for e := track.bars.Front(); e != nil; e = e.Next() {
		bar := e.Value.(bar)
		if bar.endTime+3 < time {
			track.bars.Remove(e)
			continue
		}

		barX := (time-bar.creationTime)*track.sizeInterpolator - bar.size
		barCutOff := float32(0)
		if barX < track.barStart {
			barCutOff = track.barStart - barX
			barX = track.barStart
		}
		barWidth := bar.size - barCutOff
		if barX+barWidth > track.barEnd {
			barCutOff = barX + barWidth - track.barEnd
		}
		barWidth = bar.size - barCutOff
		if barCutOff > bar.size {
			barWidth = 0
		}
		bar.quad.SetSize(mgl32.Vec2{
			barWidth,
			track.barHeight,
		})
		bar.quad.SetPosition(mgl32.Vec3{
			barX,
			1080 - track.barHeight - track.bottomOffset,
			0.6,
		})
	}
}

func (track *Track) DrawBars(ctx *g.Context) {
	for e := track.bars.Front(); e != nil; e = e.Next() {
		e.Value.(bar).quad.EnqueueForDrawing(ctx)
	}
}
func (track *Track) DrawButton(ctx *g.Context) {
	if track.isEmpty() {
		track.buttonReleased.EnqueueForDrawing(ctx)
		return
	}

	if track.shouldPress() {
		track.buttonPressed.EnqueueForDrawing(ctx)
	} else {
		track.buttonReleased.EnqueueForDrawing(ctx)
	}
}

func (track *Track) shouldPress() bool {
	if track.bars.Len() == 0 {
		return false
	}
	lastBar := track.bars.Back().Value.(bar)
	endTime := float32(glfw.GetTime()) - 3
	return lastBar.creationTime < endTime && lastBar.endTime > endTime
}

func (track *Track) pressOpportunity() bool {
	if track.bars.Len() == 0 {
		return false
	}
	lastBar := track.bars.Back().Value.(bar)
	endTime := float32(glfw.GetTime()) - 3
	return mgl32.Abs(lastBar.creationTime-endTime) < track.windowOfOpportunity
}

func (track *Track) releaseOpportunity() bool {
	if track.bars.Len() == 0 {
		return false
	}
	lastBar := track.bars.Back().Value.(bar)
	endTime := float32(glfw.GetTime()) - 3
	return mgl32.Abs(lastBar.endTime-endTime) < track.windowOfOpportunity
}

func (track *Track) isEmpty() bool {
	return track.bars.Back() == nil
}

func NewTrack(win window, barHeight float32, bottomOffset float32, windowOfOpportunity float32) Track {
	track := Track{}
	track.win = win
	track.bars = list.New()
	track.barStart = float32(0)
	track.barHeight = barHeight
	track.bottomOffset = bottomOffset
	track.sizeInterpolator = float32(win.w-80) / 3
	track.barEnd = 3 * track.sizeInterpolator
	track.windowOfOpportunity = windowOfOpportunity

	track.buttonPressed = g.NewQuadPrimitive(
		mgl32.Vec3{track.barEnd - 48, 1080 - track.barHeight/2 - 40 - bottomOffset, -1},
		mgl32.Vec2{96, 80},
	)
	track.buttonPressed.SetTexture(g.NewTextureFromFile("bojack/sprites/button/button_pressed.png"))

	track.buttonReleased = g.NewQuadPrimitive(
		mgl32.Vec3{track.barEnd - 48, 1080 - track.barHeight/2 - 40 - bottomOffset, -1},
		mgl32.Vec2{96, 80},
	)
	track.buttonReleased.SetTexture(g.NewTextureFromFile("bojack/sprites/button/button_unpressed.png"))

	return track
}
