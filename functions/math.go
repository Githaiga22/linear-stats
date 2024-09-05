package functions

import "math"

func FindLnReg(xValues, yValues []float64) (float64, float64) {
	n := float64(len(xValues))
	sumX, sumY, sumXY, sumX2 := 0.0, 0.0, 0.0, 0.0

	for i := 0; i < len(xValues); i++ {
		sumX += xValues[i]
		sumY += yValues[i]
		sumXY += xValues[i] * yValues[i]
		sumX2 += xValues[i] * xValues[i]
	}

	// Slope (m)
	m := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)

	// Intercept (b)
	// y = mx + b
	// b = y - mx

	b := (sumY / n) - m*(sumX/n)
	return m, b
}

func FindPeaCorr(xValues, yValues []float64) float64 {
	n := float64(len(xValues))
	sumX, sumY, sumXY, sumX2, sumY2 := 0.0, 0.0, 0.0, 0.0, 0.0

	for i := 0; i < len(xValues); i++ {
		sumX += xValues[i]
		sumY += yValues[i]
		sumXY += xValues[i] * yValues[i]
		sumX2 += xValues[i] * xValues[i]
		sumY2 += yValues[i] * yValues[i]
	}

	// Pearson correlation coefficient (r)
	numerator := (n*sumXY - sumX*sumY)
	denominator := math.Sqrt((n*sumX2 - sumX*sumX) * (n*sumY2 - sumY*sumY))
	r := numerator / denominator

	return r
}
