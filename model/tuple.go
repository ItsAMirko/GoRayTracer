package model

type Tuple struct {
	x float64
	y float64
	z float64
	w float64
}

func (t Tuple) IsVector() bool {
	return t.w == 0
}

func (t Tuple) IsPoint() bool {
	return t.w == 1
}

func (t Tuple) X() float64 {
	return t.x
}

func (t Tuple) Y() float64 {
	return t.y
}

func (t Tuple) Z() float64 {
	return t.z
}

func (t Tuple) W() float64 {
	return t.w
}

func (t Tuple) AddTuple(t2 Tuple) Tuple {
	return Tuple{
		x: t.x + t2.x,
		y: t.y + t2.y,
		z: t.z + t2.z,
		w: t.w + t2.w}
}

func (t Tuple) Negate() Tuple {
	return Tuple{
		x: -t.x,
		y: -t.y,
		z: -t.z,
		w: -t.w}
}

func (t Tuple) Multiply(multiplier float64) Tuple {
	return Tuple{
		x: t.x * multiplier,
		y: t.y * multiplier,
		z: t.z * multiplier,
		w: t.w * multiplier}
}

func (t Tuple) Divide(divider float64) Tuple {
	return Tuple{
		x: t.x / divider,
		y: t.y / divider,
		z: t.z / divider,
		w: t.w / divider}
}

func (t Tuple) EqualTo(t2 Tuple) bool {
	return round(t.x, 8) == round(t2.x, 8) &&
		round(t.y, 8) == round(t2.y, 8) &&
		round(t.z, 8) == round(t2.z, 8) &&
		round(t.w, 8) == round(t2.w, 8)
}
