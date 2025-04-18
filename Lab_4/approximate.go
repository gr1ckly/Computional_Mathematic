package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

type ApproxFunc struct {
	Name         string
	Func         func(float64, []float64) float64
	Approximate  func([]Point) ([]float64, error)
	Coefficients []float64
}

func NewLinear(points []Point) (*ApproxFunc, error) {
	approx := ApproxFunc{
		Name: "Линейная",
		Func: func(x float64, coeffs []float64) float64 {
			return coeffs[0] + coeffs[1]*x
		},
		Approximate: func(points []Point) ([]float64, error) {
			var sx, sy, sxx, sxy float64
			n := float64(len(points))
			for _, p := range points {
				sx += p.X
				sy += p.Y
				sxx += p.X * p.X
				sxy += p.X * p.Y
			}
			det := n*sxx - sx*sx
			if det == 0 {
				return nil, fmt.Errorf("невозможно аппроксимировать")
			}
			a := (sy*sxx - sx*sxy) / det
			b := (n*sxy - sx*sy) / det
			return []float64{a, b}, nil
		},
	}
	coef, err := approx.Approximate(points)
	if err != nil {
		return nil, err
	}
	approx.Coefficients = coef
	return &approx, nil
}

func NewPoly2(points []Point) (*ApproxFunc, error) {
	approx := ApproxFunc{
		Name: "Полином 2-й степени",
		Func: func(x float64, c []float64) float64 {
			return c[0] + c[1]*x + c[2]*x*x
		},
		Approximate: polyApprox(2),
	}
	coef, err := approx.Approximate(points)
	if err != nil {
		return nil, err
	}
	approx.Coefficients = coef
	return &approx, nil
}

func NewPoly3(points []Point) (*ApproxFunc, error) {
	approx := ApproxFunc{
		Name: "Полином 3-й степени",
		Func: func(x float64, c []float64) float64 {
			return c[0] + c[1]*x + c[2]*x*x + c[3]*x*x*x
		},
		Approximate: polyApprox(3),
	}
	coef, err := approx.Approximate(points)
	if err != nil {
		return nil, err
	}
	approx.Coefficients = coef
	return &approx, nil
}

func NewPower(points []Point) (*ApproxFunc, error) {
	approx := ApproxFunc{
		Name: "Степенная",
		Func: func(x float64, c []float64) float64 {
			return c[0] * math.Pow(x, c[1])
		},
		Approximate: func(points []Point) ([]float64, error) {
			var sumLnX, sumLnY, sumLnXLnX, sumLnXLnY float64
			n := float64(len(points))
			for _, p := range points {
				if p.X <= 0 || p.Y <= 0 {
					return nil, fmt.Errorf("отрицательные значения для логарифма")
				}
				lnx := math.Log(p.X)
				lny := math.Log(p.Y)
				sumLnX += lnx
				sumLnY += lny
				sumLnXLnX += lnx * lnx
				sumLnXLnY += lnx * lny
			}
			b := (n*sumLnXLnY - sumLnX*sumLnY) / (n*sumLnXLnX - sumLnX*sumLnX)
			a := math.Exp((sumLnY - b*sumLnX) / n)
			return []float64{a, b}, nil
		},
	}
	coef, err := approx.Approximate(points)
	if err != nil {
		return nil, err
	}
	approx.Coefficients = coef
	return &approx, nil
}

func NewExp(points []Point) (*ApproxFunc, error) {
	approx := ApproxFunc{
		Name: "Экспоненциальная",
		Func: func(x float64, c []float64) float64 {
			return c[0] * math.Exp(c[1]*x)
		},
		Approximate: func(points []Point) ([]float64, error) {
			var sx, sy, sxx, sxy float64
			n := float64(len(points))
			for _, p := range points {
				if p.Y <= 0 {
					return nil, fmt.Errorf("Y <= 0 для логарифма")
				}
				lnY := math.Log(p.Y)
				sx += p.X
				sy += lnY
				sxx += p.X * p.X
				sxy += p.X * lnY
			}
			det := n*sxx - sx*sx
			b := (n*sxy - sx*sy) / det
			a := math.Exp((sy - b*sx) / n)
			return []float64{a, b}, nil
		},
	}
	coef, err := approx.Approximate(points)
	if err != nil {
		return nil, err
	}
	approx.Coefficients = coef
	return &approx, nil
}

func NewLog(points []Point) (*ApproxFunc, error) {
	approx := ApproxFunc{
		Name: "Логарифмическая",
		Func: func(x float64, c []float64) float64 {
			return c[0] + c[1]*math.Log(x)
		},
		Approximate: func(points []Point) ([]float64, error) {
			var sx, sy, sxx, sxy float64
			n := float64(len(points))
			for _, p := range points {
				if p.X <= 0 {
					return nil, fmt.Errorf("X <= 0 для логарифма")
				}
				lnX := math.Log(p.X)
				sx += lnX
				sy += p.Y
				sxx += lnX * lnX
				sxy += lnX * p.Y
			}
			det := n*sxx - sx*sx
			b := (n*sxy - sx*sy) / det
			a := (sy - b*sx) / n
			return []float64{a, b}, nil
		},
	}
	coef, err := approx.Approximate(points)
	if err != nil {
		return nil, err
	}
	approx.Coefficients = coef
	return &approx, nil
}

func polyApprox(degree int) func([]Point) ([]float64, error) {
	return func(points []Point) ([]float64, error) {
		A := make([][]float64, degree+1)
		B := make([]float64, degree+1)
		for i := range A {
			A[i] = make([]float64, degree+1)
			for j := range A[i] {
				for _, p := range points {
					A[i][j] += math.Pow(p.X, float64(i+j))
				}
			}
			for _, p := range points {
				B[i] += math.Pow(p.X, float64(i)) * p.Y
			}
		}
		return gauss(A, B)
	}
}

func gauss(A [][]float64, B []float64) ([]float64, error) {
	n := len(B)
	for i := 0; i < n; i++ {
		max := math.Abs(A[i][i])
		maxRow := i
		for k := i + 1; k < n; k++ {
			if math.Abs(A[k][i]) > max {
				max = math.Abs(A[k][i])
				maxRow = k
			}
		}
		A[i], A[maxRow] = A[maxRow], A[i]
		B[i], B[maxRow] = B[maxRow], B[i]

		for k := i + 1; k < n; k++ {
			c := A[k][i] / A[i][i]
			for j := i; j < n; j++ {
				A[k][j] -= c * A[i][j]
			}
			B[k] -= c * B[i]
		}
	}

	x := make([]float64, n)
	for i := n - 1; i >= 0; i-- {
		x[i] = B[i]
		for j := i + 1; j < n; j++ {
			x[i] -= A[i][j] * x[j]
		}
		x[i] /= A[i][i]
	}
	return x, nil
}
