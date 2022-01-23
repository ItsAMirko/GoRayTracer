package main

import (
	"fmt"
	"math"
	"os"
	. "raytracer/model"
)

type Projectile struct {
	Position Point
	Velocity Vector
}

type Environment struct {
	Gravity Vector
	Wind    Vector
}

func (proj *Projectile) tick(env Environment) {
	proj.Position = proj.Position.AddVector(proj.Velocity)
	proj.Velocity = proj.Velocity.AddVector(env.Gravity).AddVector(env.Wind)
}

func main() {
	startPosition := CreatePoint(0, 1, 0)
	startVelocity := CreateVector(1, 1.8, 0).Normalize().Multiply(8)
	gravity := CreateVector(0, -0.05, 0)
	wind := CreateVector(-0.005, 0, 0)

	projectile := Projectile{startPosition, startVelocity}
	environment := Environment{gravity, wind}
	canvas := CreateCanvas(900, 500)

	for projectile.Position.Y() > 0 {
		canvas.SetPixel(uint64(math.Round(projectile.Position.X())),
			uint64(500-math.Round(projectile.Position.Y())),
			CreateColor(0, 1, 0))

		projectile.tick(environment)
	}

	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
	}

	err = os.WriteFile(path+"/chapters/chapter2/test.ppm", []byte(canvas.ToPPM()), 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
}
