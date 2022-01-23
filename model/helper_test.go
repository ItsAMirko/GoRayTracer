package model

import (
	"testing"
)

func TestCanRoundFloatValue(t *testing.T) {
	float := 1.23456789

	var actual float64
	var expectedValues [10]float64
	expectedValues[0] = 1
	expectedValues[1] = 1.2
	expectedValues[2] = 1.23
	expectedValues[3] = 1.235
	expectedValues[4] = 1.2346
	expectedValues[5] = 1.23457
	expectedValues[6] = 1.234568
	expectedValues[7] = 1.2345679
	expectedValues[8] = 1.23456789
	expectedValues[9] = 1.23456789

	for i := 0; i <= 9; i++ {
		actual = round(float, i)

		if actual != expectedValues[i] {
			t.Error("Iteration:", i, "Expected:", expectedValues[i], "Got:", actual)
		}
	}
}
