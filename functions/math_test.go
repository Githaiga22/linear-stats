package functions

import (
	"math"
	"testing"
)

func floatEquals(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestFindLnReg(t *testing.T) {
	tests := []struct {
		name       string
		xValues    []float64
		yValues    []float64
		expectedM  float64
		expectedB  float64
		tolerance  float64
	}{
		{
			name:      "Test case 1: Basic linear regression",
			xValues:   []float64{1, 2, 3, 4, 5},
			yValues:   []float64{2, 4, 6, 8, 10},
			expectedM: 2.0, // Slope
			expectedB: 0.0, // Intercept
			tolerance: 1e-6,
		},
		{
			name:      "Test case 2: Non-zero intercept",
			xValues:   []float64{1, 2, 3, 4, 5},
			yValues:   []float64{3, 5, 7, 9, 11},
			expectedM: 2.0,
			expectedB: 1.0,
			tolerance: 1e-6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, b := FindLnReg(tt.xValues, tt.yValues)
			if !floatEquals(m, tt.expectedM, tt.tolerance) {
				t.Errorf("FindLnReg() slope = %v, expected %v", m, tt.expectedM)
			}
			if !floatEquals(b, tt.expectedB, tt.tolerance) {
				t.Errorf("FindLnReg() intercept = %v, expected %v", b, tt.expectedB)
			}
		})
	}
}

func TestFindPeaCorr(t *testing.T) {
	tests := []struct {
		name      string
		xValues   []float64
		yValues   []float64
		expectedR float64
		tolerance float64
	}{
		{
			name:      "Test case 1: Strong positive correlation",
			xValues:   []float64{1, 2, 3, 4, 5},
			yValues:   []float64{2, 4, 6, 8, 10},
			expectedR: 1.0,
			tolerance: 1e-6,
		},
		{
			name:      "Test case 2: Strong negative correlation",
			xValues:   []float64{1, 2, 3, 4, 5},
			yValues:   []float64{10, 8, 6, 4, 2},
			expectedR: -1.0,
			tolerance: 1e-6,
		},
		{
			name:      "Test case 3: Weak correlation",
			xValues:   []float64{1, 2, 3, 4, 5},
			yValues:   []float64{5, 6, 7, 8, 9},
			expectedR: 1, 
			tolerance: 1e-6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := FindPeaCorr(tt.xValues, tt.yValues)
			if !floatEquals(r, tt.expectedR, tt.tolerance) {
				t.Errorf("FindPeaCorr() = %v, expected %v", r, tt.expectedR)
			}
		})
	}
}
