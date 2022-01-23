package model

import "math"

type Vector struct {
	*Tuple
}

func CreateVector(x float64, y float64, z float64) Vector {
	return Vector{&Tuple{x: x, y: y, z: z, w: 0}}
}

func (v Vector) AddVector(v2 Vector) Vector {
	return CreateVector(v.x+v2.x, v.y+v2.y, v.z+v2.z)
}

func (v Vector) SubtractVector(v2 Vector) Vector {
	return CreateVector(v.x-v2.x, v.y-v2.y, v.z-v2.z)
}

func (v Vector) Negate() Vector {
	return CreateVector(-v.x, -v.y, -v.z)
}

func (v Vector) Multiply(multiplier float64) Vector {
	return CreateVector(v.x*multiplier, v.y*multiplier, v.z*multiplier)
}

func (v Vector) Divide(divider float64) Vector {
	return CreateVector(v.x/divider, v.y/divider, v.z/divider)
}

func (v Vector) Magnitude() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

func (v Vector) Normalize() Vector {
	magnitude := v.Magnitude()

	return CreateVector(v.x/magnitude, v.y/magnitude, v.z/magnitude)
}

func (v Vector) Dot(v2 Vector) float64 {
	return v.x*v2.x + v.y*v2.y + v.z*v2.z + v.w*v2.w
}

func (v Vector) Cross(v2 Vector) Vector {
	return CreateVector(v.y*v2.z-v.z*v2.y,
		v.z*v2.x-v.x*v2.z,
		v.x*v2.y-v.y*v2.x)
}

func (v Vector) EqualTo(v2 Vector) bool {
	return round(v.x, 8) == round(v2.x, 8) &&
		round(v.y, 8) == round(v2.y, 8) &&
		round(v.z, 8) == round(v2.z, 8) &&
		round(v.w, 8) == round(v2.w, 8)
}
