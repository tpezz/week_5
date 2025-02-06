package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"gonum.org/v1/gonum/stat"
)

// Read CSV file and extract data in the right format
func readCSV(filePath string) ([][]float64, [][]float64) {
	file, _ := os.Open(filePath)
	defer file.Close()

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	// Create slices to store X and Y values separately
	xVals := make([][]float64, 4)
	yVals := make([][]float64, 4)

	//Iterate through rows
	for i, row := range records {

		// Skip header row
		if i == 0 {
			continue
		}

		//For each row, convert each column to a float
		x123, _ := strconv.ParseFloat(row[0], 64)
		y1, _ := strconv.ParseFloat(row[1], 64)
		y2, _ := strconv.ParseFloat(row[2], 64)
		y3, _ := strconv.ParseFloat(row[3], 64)
		x4, _ := strconv.ParseFloat(row[4], 64)
		y4, _ := strconv.ParseFloat(row[5], 64)

		//append varible to store them
		xVals[0] = append(xVals[0], x123)
		xVals[1] = append(xVals[1], x123)
		xVals[2] = append(xVals[2], x123)
		xVals[3] = append(xVals[3], x4)

		yVals[0] = append(yVals[0], y1)
		yVals[1] = append(yVals[1], y2)
		yVals[2] = append(yVals[2], y3)
		yVals[3] = append(yVals[3], y4)
	}

	return xVals, yVals
}

func main() {
	//Set file path
	filePath := "Anscombe_quartet_data.csv"

	// Read CSV and get datasets
	xVals, yVals := readCSV(filePath)

	// create slices with lenth of 4 for the results and timing
	results := make([][2]float64, 4)
	timings := make([]float64, 4)

	//Iterate through each row
	for i := range yVals {
		//start time
		start := time.Now()

		//perform linear regression
		slope, intercept := stat.LinearRegression(xVals[i], yVals[i], nil, false)
		//end time
		elapsed := time.Since(start).Seconds()
		//store results
		results[i] = [2]float64{slope, intercept}
		timings[i] = elapsed
	}

	// Print results
	for i, res := range results {
		fmt.Printf("Dataset %d: Slope = %.5f, Intercept = %.5f\n", i+1, res[0], res[1])
		fmt.Printf("Time = %.9f seconds\n", timings[i])
	}
}
