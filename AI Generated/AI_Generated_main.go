package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func linearRegression(x, y []float64) (float64, float64) {
	n, sumX, sumY, sumXY, sumX2 := float64(len(x)), 0.0, 0.0, 0.0, 0.0
	for i := range x {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumX2 += x[i] * x[i]
	}
	m := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	b := (sumY - m*sumX) / n
	return m, b
}

func main() {
	file, err := os.Open("Anscombe_quartet_data.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	var x, y1, y2, y3, x4, y4 []float64
	for i := 1; i < len(rows); i++ {
		v := rows[i]
		xVal, _ := strconv.ParseFloat(v[0], 64)
		y1Val, _ := strconv.ParseFloat(v[1], 64)
		y2Val, _ := strconv.ParseFloat(v[2], 64)
		y3Val, _ := strconv.ParseFloat(v[3], 64)
		x4Val, _ := strconv.ParseFloat(v[4], 64)
		y4Val, _ := strconv.ParseFloat(v[5], 64)
		x = append(x, xVal)
		y1 = append(y1, y1Val)
		y2 = append(y2, y2Val)
		y3 = append(y3, y3Val)
		x4 = append(x4, x4Val)
		y4 = append(y4, y4Val)
	}

	m1, b1 := linearRegression(x, y1)
	m2, b2 := linearRegression(x, y2)
	m3, b3 := linearRegression(x, y3)
	m4, b4 := linearRegression(x4, y4)

	fmt.Printf("y1: y = %.2fx + %.2f\n", m1, b1)
	fmt.Printf("y2: y = %.2fx + %.2f\n", m2, b2)
	fmt.Printf("y3: y = %.2fx + %.2f\n", m3, b3)
	fmt.Printf("y4: y = %.2fx + %.2f\n", m4, b4)
}
