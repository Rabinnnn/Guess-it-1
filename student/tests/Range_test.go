package functions

import "testing"

// TestMean checks if the Mean function calculates the average correctly
func TestRange(t *testing.T) {
	data := []float64{23, 45}
	expected1 := 12
	expected2 := 56
	result1, result2 := Range(data)

	if int(result1) != expected1 || int(result2) != expected2 {
		t.Errorf("Range(%v) = %v %v; want %v = %v %v", data, result1, result2, data, expected1, expected2)
	}
}
