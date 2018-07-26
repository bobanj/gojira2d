package main

import (
		"github.com/markov/gojira2d/pkg/app"
		)

func main() {
	app.Init(640, 480, false, "Run For Your Life!")
	defer app.Terminate()

	app.MainLoop(func(speed float64) {
		//NOOP
	}, func() {
		//NOOP
	})
}