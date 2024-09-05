package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	
	"linear-stats/functions"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Please provide the path to the data file as an argument.")
	}

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var yValues []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err)
		}
		yValues = append(yValues, value)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Create xValues corresponding to the line numbers
	xValues := make([]float64, len(yValues))
	for i := range xValues {
		xValues[i] = float64(i)
	}

	// Calculate Linear Regression Line
	m, b := functions.FindLnReg(xValues, yValues)
	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", m, b)

	// Calculate Pearson Correlation Coefficient
	r := functions.FindPeaCorr(xValues, yValues)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", r)
}
