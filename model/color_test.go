package model

import (
	"testing"
)

func TestCanCreateAColor(t *testing.T) {
	color := CreateColor(0.1, 0.2, 0.3)

	if color.red != 0.1 {
		t.Error("Expected for red:", 0.1, "Got:", color.red)
	}
	if color.green != 0.2 {
		t.Error("Expected for green:", 0.2, "Got:", color.green)
	}
	if color.blue != 0.3 {
		t.Error("Expected for blue:", 0.3, "Got:", color.blue)
	}
}

func TestCanAddUpTwoColors(t *testing.T) {
	color1 := CreateColor(0.9, 0.6, 0.75)
	color2 := CreateColor(0.7, 0.1, 0.25)

	expected := CreateColor(1.6, 0.7, 1)
	actual := color1.AddColor(color2)

	if actual.EqualTo(expected) == false {
		t.Error("Expected:", expected, "Got:", actual)
	}
}

func TestCanSubtractTwoColors(t *testing.T) {
	color1 := CreateColor(0.9, 0.6, 0.75)
	color2 := CreateColor(0.7, 0.1, 0.25)

	expected := CreateColor(0.2, 0.5, 0.5)
	actual := color1.SubtractColor(color2)

	if actual.EqualTo(expected) == false {
		t.Error("Expected:", expected, "Got:", actual)
	}
}

func TestCanMultiplyAColorByAScalar(t *testing.T) {
	color1 := CreateColor(0.2, 0.3, 0.4)
	scalar := 2.0

	expected := CreateColor(0.4, 0.6, 0.8)
	actual := color1.Multiply(scalar)

	if actual.EqualTo(expected) == false {
		t.Error("Expected:", expected, "Got:", actual)
	}
}

func TestCanMultiplyAColorByAColor(t *testing.T) {
	color1 := CreateColor(1, 0.2, 0.4)
	color2 := CreateColor(0.9, 1, 0.1)

	expected := CreateColor(0.9, 0.2, 0.04)
	actual := color1.MultiplyColor(color2)

	if actual.EqualTo(expected) == false {
		t.Error("Expected:", expected, "Got:", actual)
	}
}

func TestCanBeConvertedToHexString(t *testing.T) {
	color1 := CreateColor(1, 0.5, 0)
	color2 := CreateColor(-1, 0.75, 2)

	expected1 := "255 128 0"
	actual1 := color1.ToHexString()

	if expected1 != actual1 {
		t.Error("Expected:", expected1, "Got:", actual1)
	}

	expected2 := "0 191 255"
	actual2 := color2.ToHexString()

	if expected2 != actual2 {
		t.Error("Expected:", expected2, "Got:", actual2)
	}
}
