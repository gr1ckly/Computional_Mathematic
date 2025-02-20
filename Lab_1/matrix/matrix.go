package matrix

import (
	"fmt"
	"math"
)

type Matrix struct {
	size, maxIter int
	data          [][]float64
	vecB          []float64
	accuracy      float64
}

func BuildMatrix(size, maxIter int, data [][]float64, vecB []float64, accuracy float64) *Matrix {
	matrix := &Matrix{size, maxIter, data, vecB, accuracy}
	return matrix
}

func (matrix *Matrix) IsDiagonallyDominant() bool {
	for i := 0; i < matrix.size; i++ {
		sum := 0.0
		for j := 0; j < matrix.size; j++ {
			if i != j {
				sum += math.Abs(matrix.data[i][j])
			}
		}
		if math.Abs(matrix.data[i][i]) <= sum {
			return false
		}
	}
	return true
}

func (matrix *Matrix) MakeDiagonallyDominant() bool {
	for i := 0; i < matrix.size; i++ {
		maxRow := i
		maxVal := math.Abs(matrix.data[i][i])
		for k := i + 1; k < matrix.size; k++ {
			if math.Abs(matrix.data[k][i]) > maxVal {
				maxVal = math.Abs(matrix.data[k][i])
				maxRow = k
			}
		}
		if maxRow != i {
			matrix.data[i], matrix.data[maxRow] = matrix.data[maxRow], matrix.data[i]
			matrix.vecB[i], matrix.vecB[maxRow] = matrix.vecB[maxRow], matrix.vecB[i]
		}
	}
	if matrix.IsDiagonallyDominant() {
		return true
	} else {
		return false
	}
}

func (matrix *Matrix) StringMatrix() string {
	ans := ""
	for i := 0; i < matrix.size; i++ {
		for j := 0; j < matrix.size; j++ {
			ans += fmt.Sprintf("%v", matrix.data[i][j]) + " "
		}
		ans += fmt.Sprintf("%v", matrix.vecB[i]) + string('\n')
	}
	return ans[:len(ans)-1]
}

func (matrix *Matrix) GaussSeidel() ([]float64, int, []float64, error) {
	x := make([]float64, matrix.size)
	for idx, _ := range x {
		x[idx] = matrix.vecB[idx] / matrix.data[idx][idx]
	}
	errors := make([]float64, matrix.size)

	for iter := 0; iter < matrix.maxIter; iter++ {
		maxDiff := 0.0
		for i := 0; i < matrix.size; i++ {
			sum := 0.0
			for j := 0; j < matrix.size; j++ {
				if j != i {
					sum += matrix.data[i][j] * x[j]
				}
			}
			newX := float64(matrix.vecB[i]-sum) / float64(matrix.data[i][i])
			errors[i] = math.Abs(newX - x[i])
			if errors[i] > maxDiff {
				maxDiff = errors[i]
			}
			x[i] = newX
		}
		if maxDiff <= matrix.accuracy {
			return x, iter + 1, errors, nil
		}
	}
	return nil, matrix.maxIter, errors, fmt.Errorf("Решения за %d итераций при указанной точности не найдено\n", matrix.maxIter)
}

func (matrix *Matrix) Norm() float64 {
	maxNorm := 0.0
	for _, row := range matrix.data {
		sum := 0.0
		for _, val := range row {
			sum += math.Abs(val)
		}
		if sum > maxNorm {
			maxNorm = sum
		}
	}
	return maxNorm
}
