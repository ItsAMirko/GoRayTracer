package model

import "math"

func round(valued float64, floatingPoints int) float64 {
	return math.Round(valued*math.Pow10(floatingPoints)) / math.Pow10(floatingPoints)
}
