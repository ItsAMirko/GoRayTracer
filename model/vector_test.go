package model

import (
	"math"
	"testing"
)

func TestCanCreateAVector(t *testing.T) {
	vector := CreateVector(4, -4, 3)

	if vector.x != 4 {
		t.Error("Expected for x:", 4, "Got:", vector.x)
	}
	if vector.y != -4 {
		t.Error("Expected for y:", -4, "Got:", vector.y)
	}
	if vector.z != 3 {
		t.Error("Expected for z:", 3, "Got:", vector.z)
	}
	if vector.IsVector() != true {
		t.Error("Expected a point")
	}
}

func TestCanAddAVectorToAVector(t *testing.T) {
	vector1 := CreateVector(3, -2, 1)
	vector2 := CreateVector(1, 1, -1)

	expected := CreateVector(4, -1, 0)
	actual := vector1.AddVector(vector2)

	if actual.EqualTo(expected) == false {
		t.Error("Expected:", expected, "Got:", actual)
	}
}

func TestCanSubtractTwoVectors(t *testing.T) {
	vector1 := CreateVector(3, 2, 1)
	vector2 := CreateVector(5, 6, 7)

	expected := CreateVector(-2, -4, -6)
	actual := vector1.SubtractVector(vector2)

	if actual.EqualTo(expected) == false {
		t.Error("Expected:", expected, "Got:", actual)
	}
}

func TestCanNegateAVector(t *testing.T) {
	vector := CreateVector(1, -2, 3)

	expected := CreateVector(-1, 2, -3)
	actual := vector.Negate()

	if actual.EqualTo(expected) == false {
		t.Error("Expected:", expected, "Got:", actual)
	}
}

func TestCanMultiplyAVectorByAFraction(t *testing.T) {
	vector := CreateVector(1, -2, 3)
	fraction := 0.5

	expected := CreateVector(0.5, -1, 1.5)
	actual := vector.Multiply(fraction)

	if actual.EqualTo(expected) == false {
		t.Error("Expected:", expected, "Got:", actual)
	}
}

func TestCanDivideAVectorByAScalar(t *testing.T) {
	vector := CreateVector(1, -2, 3)
	scalar := 2.0

	expected := CreateVector(0.5, -1, 1.5)
	actual := vector.Divide(scalar)

	if actual.EqualTo(expected) == false {
		t.Error("Expected:", expected, "Got:", actual)
	}
}

func TestCanComputeTheMagnitude(t *testing.T) {
	vector1 := CreateVector(1, 0, 0)
	vector2 := CreateVector(0, 1, 0)
	vector3 := CreateVector(0, 0, 1)
	vector4 := CreateVector(1, 2, 3)
	vector5 := CreateVector(-1, -2, -3)

	if vector1.Magnitude() != 1 {
		t.Error("Expected magnitude of:", 1, "Got:", vector1.Magnitude())
	}
	if vector2.Magnitude() != 1 {
		t.Error("Expected magnitude of:", 1, "Got:", vector2.Magnitude())
	}
	if vector3.Magnitude() != 1 {
		t.Error("Expected magnitude of:", 1, "Got:", vector3.Magnitude())
	}
	if vector4.Magnitude() != math.Sqrt(14) {
		t.Error("Expected magnitude of:", math.Sqrt(14), "Got:", vector4.Magnitude())
	}
	if vector5.Magnitude() != math.Sqrt(14) {
		t.Error("Expected magnitude of:", math.Sqrt(14), "Got:", vector5.Magnitude())
	}
}

func TestCanNormalizingAVector(t *testing.T) {
	vector1 := CreateVector(4, 0, 0)
	vector2 := CreateVector(1, 2, 3)

	expected1 := CreateVector(1, 0, 0)
	actual1 := vector1.Normalize()

	if actual1.EqualTo(expected1) == false {
		t.Error("Expected:", expected1, "Got:", actual1)
	}

	expected2 := CreateVector(1/math.Sqrt(14), 2/math.Sqrt(14), 3/math.Sqrt(14))
	actual2 := vector2.Normalize()

	if actual2.EqualTo(expected2) == false {
		t.Error("Expected:", expected2, "Got:", actual2)
	}
	if actual2.Magnitude() != 1 {
		t.Error("Expected magnitude:", 1, "Got:", actual2.Magnitude())
	}
}

func TestCanCalculateDotProduct(t *testing.T) {
	vector1 := CreateVector(1, 2, 3)
	vector2 := CreateVector(2, 3, 4)

	if vector1.Dot(vector2) != 20 {
		t.Error("Expected dot product:", 20, "Got:", vector1.Dot(vector2))
	}
}

func TestCanCalculateCrossProduct(t *testing.T) {
	vector1 := CreateVector(1, 2, 3)
	vector2 := CreateVector(2, 3, 4)

	expected1 := CreateVector(-1, 2, -1)
	actual1 := vector1.Cross(vector2)

	if actual1.EqualTo(expected1) == false {
		t.Error("Expected:", expected1, "Got:", actual1)
	}

	expected2 := CreateVector(1, -2, 1)
	actual2 := vector2.Cross(vector1)

	if actual2.EqualTo(expected2) == false {
		t.Error("Expected:", expected2, "Got:", actual2)
	}
}
