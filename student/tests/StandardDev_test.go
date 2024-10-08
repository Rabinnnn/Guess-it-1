package functions

import (
	"math"
	"testing"
)

// TestStandardDeviation checks if the StandardDeviation function calculates the standard deviation correctly
func TestStandardDeviation(t *testing.T) {
	data := []float64{23, 45, 67, 89, 73}
	expected := 23.03
	result := StandardDeviation(data)

	if math.Abs(result-expected) > 0.01 { // Allowing a small margin of error for floating-point calculations
		t.Errorf("StandardDeviation(%v) = %v; want %v", data, result, expected)
	}
}
