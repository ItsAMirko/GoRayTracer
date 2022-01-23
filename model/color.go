package model

import (
	"math"
	"strconv"
)

type Color struct {
	red   float64
	green float64
	blue  float64
}

func CreateColor(red float64, green float64, blue float64) Color {
	return Color{red: red, green: green, blue: blue}
}

func (c Color) AddColor(c2 Color) Color {
	return CreateColor(
		c.red+c2.red,
		c.green+c2.green,
		c.blue+c2.blue)
}

func (c Color) SubtractColor(c2 Color) Color {
	return CreateColor(
		c.red-c2.red,
		c.green-c2.green,
		c.blue-c2.blue)
}

func (c Color) Multiply(multiplier float64) Color {
	return CreateColor(
		c.red*multiplier,
		c.green*multiplier,
		c.blue*multiplier)
}

func (c Color) MultiplyColor(c2 Color) Color {
	return CreateColor(
		c.red*c2.red,
		c.green*c2.green,
		c.blue*c2.blue)
}

func (c Color) ToHexString() string {
	red := calcHexValue(c.red)
	green := calcHexValue(c.green)
	blue := calcHexValue(c.blue)

	return strconv.Itoa(red) + " " + strconv.Itoa(green) + " " + strconv.Itoa(blue)
}

func (c Color) EqualTo(c2 Color) bool {
	return round(c.red, 8) == round(c2.red, 8) &&
		round(c.green, 8) == round(c2.green, 8) &&
		round(c.blue, 8) == round(c2.blue, 8)
}

func calcHexValue(floatValue float64) int {
	hex := math.Round(floatValue * 255)
	if hex > 255 {
		hex = 255
	} else if hex < 0 {
		hex = 0
	}

	return int(hex)
}
