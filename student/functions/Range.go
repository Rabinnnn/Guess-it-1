package functions

// Predict the next number range based on the mean and standard deviation
func Range(numbers []float64) (float64, float64) {
	m := Mean(numbers)
	sd := StandardDeviation(numbers)
	// use a prediction range of mean ± 2*standard deviation
	return m - 2*sd, m + 2*sd
}