package main

import (
	"fmt"
	a "gojira2d/pkg/app"
	box2d "gojira2d/pkg/physics/box2d"
	"sort"
)

func main() {
	app := a.InitApp(800, 600, false, "Physics")
	defer a.TerminateApp()

	// Define the gravity vector.
	gravity := box2d.MakeB2Vec2(0.0, -10.0)

	// Construct a world object, which will hold and simulate the rigid bodies.
	world := box2d.MakeB2World(gravity)

	characters := make(map[string]*box2d.B2Body)

	// Ground body
	{
		bd := box2d.MakeB2BodyDef()
		ground := world.CreateBody(&bd)

		shape := box2d.MakeB2EdgeShape()
		shape.Set(box2d.MakeB2Vec2(-20.0, 0.0), box2d.MakeB2Vec2(20.0, 0.0))
		ground.CreateFixture(&shape, 0.0)
		characters["00_ground"] = ground
	}

	// Circle character
	{
		bd := box2d.MakeB2BodyDef()
		bd.Position.Set(3.0, 5.0)
		bd.Type = box2d.B2BodyType.B2_dynamicBody
		bd.FixedRotation = true
		bd.AllowSleep = false

		body := world.CreateBody(&bd)

		shape := box2d.MakeB2CircleShape()
		shape.M_radius = 0.5

		fd := box2d.MakeB2FixtureDef()
		fd.Shape = &shape
		fd.Density = 20.0
		body.CreateFixtureFromDef(&fd)
		characters["09_circlecharacter1"] = body
	}

	// Circle character
	{
		bd := box2d.MakeB2BodyDef()
		bd.Position.Set(-7.0, 6.0)
		bd.Type = box2d.B2BodyType.B2_dynamicBody
		bd.AllowSleep = false

		body := world.CreateBody(&bd)

		shape := box2d.MakeB2CircleShape()
		shape.M_radius = 0.25

		fd := box2d.MakeB2FixtureDef()
		fd.Shape = &shape
		fd.Density = 20.0
		fd.Friction = 1.0
		body.CreateFixtureFromDef(&fd)

		characters["10_circlecharacter2"] = body
	}

	// Prepare for simulation. Typically we use a time step of 1/60 of a
	// second (60Hz) and 10 iterations. This provides a high quality simulation
	// in most game scenarios.
	timeStep := 1.0 / 60.0
	velocityIterations := 8
	positionIterations := 3

	output := ""

	characterNames := make([]string, 0)
	for k, _ := range characters {
		characterNames = append(characterNames, k)
	}
	i := 0

	sort.Strings(characterNames)

	app.MainLoop(func(speed float64) {
		world.Step(timeStep, velocityIterations, positionIterations)
		i++
	}, func() {
		for _, name := range characterNames {
			character := characters[name]
			position := character.GetPosition()
			angle := character.GetAngle()
			msg := fmt.Sprintf("%v(%s): %4.3f %4.3f %4.3f\n", i, name, position.X, position.Y, angle)
			fmt.Print(msg)
			output += msg
		}
	})
}
