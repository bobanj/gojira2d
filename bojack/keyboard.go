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
	switch action {
	case glfw.Press:
		keyPressed = true
	case glfw.Release:
		keyPressed = false
	default:
		return
	}
	for _, p := range players {
		if key == p.key {
			p.lastKeyInteraction = glfw.GetTime()
			if pressOpportunity() {
				if keyPressed {
					p.speedUp()
				} else {
					p.slowDown()
				}
			} else if releaseOpportunity() {
				if keyPressed {
					p.slowDown()
				} else {
					p.speedUp()
				}
			} else if shouldPress() != keyPressed {
				p.slowDown()
			}
		}
	}
}
