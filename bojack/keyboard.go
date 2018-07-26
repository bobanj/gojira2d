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
	switch action {
	case glfw.Press:
		switch key {
		case players[0].key:
			players[0].keyPressed = true
		case players[1].key:
			players[1].keyPressed = true
		case players[2].key:
			players[2].keyPressed = true
		}
	case glfw.Release:
		switch key {
		case players[0].key:
			players[0].keyPressed = false
		case players[1].key:
			players[1].keyPressed = false
		case players[2].key:
			players[2].keyPressed = false
		}
	}
}
