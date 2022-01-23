package main

import (
	"fmt"
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
	startVelocity := CreateVector(1, 1, 0).Normalize()
	gravity := CreateVector(0, -0.1, 0)
	wind := CreateVector(-0.01, 0, 0)

	projectile := Projectile{startPosition, startVelocity}
	environment := Environment{gravity, wind}

	for projectile.Position.Y() > 0 {
		fmt.Printf("%.4f | %.4f\n", projectile.Position.X(), projectile.Position.Y())

		projectile.tick(environment)
	}
}
