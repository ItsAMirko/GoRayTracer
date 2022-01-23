package model

import (
	"testing"
)

func TestCanCreateAPoint(t *testing.T) {
	point := CreatePoint(4, -4, 3)

	if point.x != 4 {
		t.Error("Expected for x:", 4, "Got:", point.x)
	}
	if point.y != -4 {
		t.Error("Expected for y:", -4, "Got:", point.y)
	}
	if point.z != 3 {
		t.Error("Expected for z:", 3, "Got:", point.z)
	}
	if point.IsPoint() != true {
		t.Error("Expected a point")
	}
}

func TestCanSubtractTwoPoints(t *testing.T) {
	point1 := CreatePoint(3, 2, 1)
	point2 := CreatePoint(5, 6, 7)

	expected := CreateVector(-2, -4, -6)
	actual := point1.SubtractPoint(point2)

	if actual.EqualTo(expected) == false {
		t.Error("Expected:", expected, "Got:", actual)
	}
}

func TestCanAddAVectorFromAPoint(t *testing.T) {
	point1 := CreatePoint(3, -2, 1)
	vector := CreateVector(1, 1, -1)

	expected := CreatePoint(4, -1, 0)
	actual := point1.AddVector(vector)

	if actual.EqualTo(expected) == false {
		t.Error("Expected:", expected, "Got:", actual)
	}
}

func TestCanSubtractAVectorFromAPoint(t *testing.T) {
	point1 := CreatePoint(3, 2, 1)
	vector := CreateVector(5, 6, 7)

	expected := CreatePoint(-2, -4, -6)
	actual := point1.SubtractVector(vector)

	if actual.EqualTo(expected) == false {
		t.Error("Expected:", expected, "Got:", actual)
	}
}
