package model

import (
	"testing"
)

func TestCanCreateATuple(t *testing.T) {
	point := Tuple{x: 1, y: 2, z: 3, w: 4}

	if point.X() != 1 {
		t.Error("Expected for x:", 1, "Got:", point.x)
	}
	if point.Y() != 2 {
		t.Error("Expected for y:", 2, "Got:", point.y)
	}
	if point.Z() != 3 {
		t.Error("Expected for z:", 3, "Got:", point.z)
	}
	if point.W() != 4 {
		t.Error("Expected for w:", 4, "Got:", point.w)
	}
}

func TestTupleCanBeAPointOrVector(t *testing.T) {
	vector := Tuple{x: 1, y: 2, z: 3, w: 0}
	point := Tuple{x: 1, y: 2, z: 3, w: 1}

	if vector.IsVector() != true {
		t.Error("Expected tuple to be a vector")
	}
	if point.IsPoint() != true {
		t.Error("Expected tuple to be a point")
	}
}

func TestCanAddUpTwoTuples(t *testing.T) {
	tuple1 := Tuple{x: 3, y: -2, z: 5, w: 1}
	tuple2 := Tuple{x: -2, y: 3, z: 1, w: 0}

	expected := Tuple{x: 1, y: 1, z: 6, w: 1}
	actual := tuple1.AddTuple(tuple2)

	if actual.EqualTo(expected) == false {
		t.Error("Expected:", expected, "Got:", actual)
	}
}

func TestCanNegateATuple(t *testing.T) {
	tuple := Tuple{x: 1, y: -2, z: 3, w: -4}

	expected := Tuple{x: -1, y: 2, z: -3, w: 4}
	actual := tuple.Negate()

	if actual.EqualTo(expected) == false {
		t.Error("Expected:", expected, "Got:", actual)
	}
}

func TestCanMultiplyATupleByAScalar(t *testing.T) {
	tuple1 := Tuple{x: 1, y: -2, z: 3, w: -4}
	scalar := 3.5

	expected := Tuple{x: 3.5, y: -7, z: 10.5, w: -14}
	actual := tuple1.Multiply(scalar)

	if actual.EqualTo(expected) == false {
		t.Error("Expected:", expected, "Got:", actual)
	}
}

func TestCanMultiplyATupleByAFraction(t *testing.T) {
	tuple := Tuple{x: 1, y: -2, z: 3, w: -4}
	fraction := 0.5

	expected := Tuple{x: 0.5, y: -1, z: 1.5, w: -2}
	actual := tuple.Multiply(fraction)

	if actual.EqualTo(expected) == false {
		t.Error("Expected:", expected, "Got:", actual)
	}
}

func TestCanDivideATupleByAScalar(t *testing.T) {
	tuple := Tuple{x: 1, y: -2, z: 3, w: -4}
	scalar := 2.0

	expected := Tuple{x: 0.5, y: -1, z: 1.5, w: -2}
	actual := tuple.Divide(scalar)

	if actual.EqualTo(expected) == false {
		t.Error("Expected:", expected, "Got:", actual)
	}
}
