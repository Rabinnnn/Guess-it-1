package functions

import "math"

// Calculate the standard deviation 
func StandardDeviation(numbers []float64) float64 {
	m := Mean(numbers)
	var sum float64
	for _, number := range numbers {
		sum += (number - m) * (number - m)
	}
	return math.Sqrt(sum / float64(len(numbers)))
}