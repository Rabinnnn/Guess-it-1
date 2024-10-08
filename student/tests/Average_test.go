package functions

import "testing"

// TestMean checks if the Mean function calculates the average correctly
func TestMean(t *testing.T) {
	data := []float64{2, 4, 3}
	expected := 3.0
	result := Mean(data)

	if result != expected {
		t.Errorf("Mean(%v) = %v; want %v", data, result, expected)
	}
}
