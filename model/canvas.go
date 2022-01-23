package model

import (
	"strconv"
)

type Canvas struct {
	width  uint64
	height uint64
	pixels [][]Color
}

func CreateCanvas(width uint64, height uint64) Canvas {
	pixels := make([][]Color, width)
	for i := uint64(0); i < width; i++ {
		pixels[i] = make([]Color, height)
		for j := uint64(0); j < height; j++ {
			pixels[i][j] = CreateColor(0, 0, 0)
		}
	}

	return Canvas{width: width, height: height, pixels: pixels}
}

func (c Canvas) SetPixel(x uint64, y uint64, color Color) {
	if x >= c.width || y >= c.height {
		return
	}

	c.pixels[x][y] = color
}

func (c Canvas) ToPPM() string {
	header := "P3\n" +
		strconv.FormatInt(int64(c.width), 10) + " " +
		strconv.FormatInt(int64(c.height), 10) +
		"\n255\n"

	body, line := "", ""
	for j := uint64(0); j < c.height; j++ {
		for i := uint64(0); i < c.width; i++ {
			if len(line+c.pixels[i][j].ToHexString()) > 70 {
				body = body + line + "\n"
				line = ""
			}
			line = line + c.pixels[i][j].ToHexString() + " "
		}
		body = body + line + "\n"
		line = ""
	}

	return header + body
}
