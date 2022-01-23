package model

type Point struct {
	*Tuple
}

func CreatePoint(x float64, y float64, z float64) Point {
	return Point{&Tuple{x: x, y: y, z: z, w: 1}}
}

func (p Point) SubtractPoint(p2 Point) Vector {
	return CreateVector(p.x-p2.x, p.y-p2.y, p.z-p2.z)
}

func (p Point) AddVector(v Vector) Point {
	return CreatePoint(p.x+v.x, p.y+v.y, p.z+v.z)
}

func (p Point) SubtractVector(v Vector) Point {
	return CreatePoint(p.x-v.x, p.y-v.y, p.z-v.z)
}

func (p Point) EqualTo(p2 Point) bool {
	return round(p.x, 8) == round(p2.x, 8) &&
		round(p.y, 8) == round(p2.y, 8) &&
		round(p.z, 8) == round(p2.z, 8) &&
		round(p.w, 8) == round(p2.w, 8)
}
