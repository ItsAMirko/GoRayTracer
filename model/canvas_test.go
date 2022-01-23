package model

import (
	"strings"
	"testing"
)

func TestCanCreateACanvas(t *testing.T) {
	canvas := CreateCanvas(10, 20)

	if canvas.width != 10 {
		t.Error("Expected for width:", 10, "Got:", canvas.width)
	}
	if canvas.height != 20 {
		t.Error("Expected for height:", 20, "Got:", canvas.height)
	}

	expected := CreateColor(0, 0, 0)
	var actual Color
	var i, j uint64

	for i = 0; i < 10; i++ {
		for j = 0; j < 20; j++ {
			actual = canvas.pixels[i][j]
			if actual.EqualTo(expected) == false {
				t.Error("Iteration:", i, "Expected:", expected, "Got:", actual)
			}
		}
	}
}

func TestCanSetAPixel(t *testing.T) {
	canvas := CreateCanvas(1, 1)

	expected := CreateColor(0, 0, 0)
	canvas.SetPixel(0, 0, expected)
	canvas.SetPixel(1, 1, expected)
	actual := canvas.pixels[0][0]

	if actual.EqualTo(expected) == false {
		t.Error("Expected:", expected, "Got:", actual)
	}
}

func TestCanCreatePPMFileContent(t *testing.T) {
	canvas := CreateCanvas(5, 3)
	canvas.SetPixel(0, 0, CreateColor(1.5, 0, 0))
	canvas.SetPixel(2, 1, CreateColor(0, 0.5, 0))
	canvas.SetPixel(4, 2, CreateColor(-0.5, 0, 1))

	expectedArray := []string{
		"P3",
		"5 3",
		"255",
		"255 0 0 0 0 0 0 0 0 0 0 0 0 0 0 ",
		"0 0 0 0 0 0 0 128 0 0 0 0 0 0 0 ",
		"0 0 0 0 0 0 0 0 0 0 0 0 0 0 255 ",
		"",
	}
	actualArray := strings.Split(strings.ReplaceAll(canvas.ToPPM(), "\r\n", "\n"), "\n")

	for index, actual := range actualArray {
		expected := expectedArray[index]
		if expected != actual {
			t.Error("Index:", index, "Expected:", expected, "Got:", actual)
		}
	}
}

func TestRespectsMaxCharsPerLineWhenCreationPPMFileContent(t *testing.T) {
	canvas := CreateCanvas(10, 1)
	for x := 0; x < 10; x++ {
		canvas.SetPixel(uint64(x), 0, CreateColor(1, 0.66, 0.33))
	}

	expectedArray := []string{
		"P3",
		"10 1",
		"255",
		"255 168 84 255 168 84 255 168 84 255 168 84 255 168 84 255 168 84 ",
		"255 168 84 255 168 84 255 168 84 255 168 84 ",
		"",
	}
	actualArray := strings.Split(strings.ReplaceAll(canvas.ToPPM(), "\r\n", "\n"), "\n")

	for index, actual := range actualArray {
		expected := expectedArray[index]
		if expected != actual {
			t.Error("Index:", index, "Expected:", expected, "Got:", actual)
		}
	}
}
