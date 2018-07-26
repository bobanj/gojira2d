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

