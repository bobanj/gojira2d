package main

import (
	"log"
	"github.com/markov/gojira2d/pkg/app"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func RegisterKeyCallback(callback glfw.KeyCallback) {
	if keyCallbackFunc != nil {
		log.Panic("A keyboard key-callback is already registered!")
	}
	keyCallbackFunc = callback
	app.GetWindow().SetKeyCallback(callback)
}

func UnregisterKeyCallback() {
	keyCallbackFunc = nil
	app.GetWindow().SetKeyCallback(nil)
}

func HandleKeyPress(key glfw.Key, action glfw.Action, players []*Player) {
	var keyPressed bool
	if players[0].canStart && players[1].canStart && players[2].canStart {
		switch action {
		case glfw.Press:
			keyPressed = true
		case glfw.Release:
			keyPressed = false
		default:
			return
		}
		for _, p := range players {
			if key == p.key0 {
				p.lastKeyInteraction = glfw.GetTime()
				if track0.pressOpportunity() {
					if keyPressed {
						p.speedUp()
					} else {
						p.slowDown()
					}
				} else if track0.releaseOpportunity() {
					if keyPressed {
						p.slowDown()
					} else {
						p.speedUp()
					}
				} else if track0.shouldPress() != keyPressed {
					p.slowDown()
				}
			}
		}
	}
}
