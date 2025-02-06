package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gonum.org/v1/gonum/stat"
)

// readCSV reads a CSV file and extracts data into two-dimensional slices for X and Y values.
func readCSV(filePath string) ([][]float64, [][]float64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, nil, err
	}

	// Create slices to store X and Y values separately
	xVals := make([][]float64, 4)
	yVals := make([][]float64, 4)
	for i := range xVals {
		xVals[i] = []float64{}
		yVals[i] = []float64{}
	}

	// Iterate through rows
	for i, row := range records {
		// Skip header row
		if i == 0 {
			continue
		}

		// For each row, convert each column to a float
		x123, err := strconv.ParseFloat(row[0], 64)
		if err != nil {
			return nil, nil, err
		}
		y1, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			return nil, nil, err
		}
		y2, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			return nil, nil, err
		}
		y3, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			return nil, nil, err
		}
		x4, err := strconv.ParseFloat(row[4], 64)
		if err != nil {
			return nil, nil, err
		}
		y4, err := strconv.ParseFloat(row[5], 64)
		if err != nil {
			return nil, nil, err
		}

		// Append variables to store them
		xVals[0] = append(xVals[0], x123)
		xVals[1] = append(xVals[1], x123)
		xVals[2] = append(xVals[2], x123)
		xVals[3] = append(xVals[3], x4)

		yVals[0] = append(yVals[0], y1)
		yVals[1] = append(yVals[1], y2)
		yVals[2] = append(yVals[2], y3)
		yVals[3] = append(yVals[3], y4)
	}

	return xVals, yVals, nil
}

func main() {
	// Set file path
	filePath := "Anscombe_quartet_data.csv"

	// Read CSV and get datasets
	xVals, yVals, err := readCSV(filePath)
	if err != nil {
		log.Fatalf("Failed to read CSV: %v", err)
	}

	// Create slices with length of 4 for the results and timing
	results := make([][2]float64, 4)
	timings := make([]float64, 4)

	// Iterate through each dataset
	for i := range yVals {
		// Start time
		start := time.Now()

		// Perform linear regression
		slope, intercept := stat.LinearRegression(xVals[i], yVals[i], nil, false)

		// End time
		elapsed := time.Since(start).Seconds()

		// Store results
		results[i] = [2]float64{slope, intercept}
		timings[i] = elapsed
	}

	// Print results
	for i, res := range results {
		fmt.Printf("Dataset %d: Slope = %.5f, Intercept = %.5f\n", i+1, res[0], res[1])
		fmt.Printf("Time = %.9f seconds\n", timings[i])
	}
}
