package main

import (
	"bufio"
	"fmt"
	"guess-it-1/functions"
	"os"
	"strconv"
)

func main() {
	const windowSize = 5 // Number of immediate previous numbers to consider for prediction
	var data []float64
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		value, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("Error")
			continue
		}
		data = append(data, value)

		if len(data) >= windowSize {
			data = data[len(data)-windowSize:]
		}

		if len(data) > 1 {
			lower, upper := functions.Range(data)
			fmt.Printf("%d %d\n", int(lower), int(upper))
		}
	}
}