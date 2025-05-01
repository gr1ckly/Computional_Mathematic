package main

import "math"

type Point struct {
	X, Y float64
}

func GetFuncsDescription() map[string]string {
	return map[string]string{
		"1": "sin(x)",
		"2": "cos(x)",
		"3": "exp(x)",
	}
}

func GetFuncs() map[string]func(float64) float64 {
	return map[string]func(float64) float64{
		"1": math.Sin,
		"2": math.Cos,
		"3": math.Exp,
	}
}

func FiniteDifferences(points []Point) [][]float64 {
	ans := make([][]float64, len(points))
	for i := range ans {
		ans[i] = make([]float64, len(points))
	}
	for i := 0; i < len(ans); i++ {
		ans[i][0] = points[i].Y
	}
	for j := 1; j < len(ans); j++ {
		for i := 0; i < len(ans)-j; i++ {
			ans[i][j] = ans[i+1][j-1] - ans[i][j-1]
		}
	}
	return ans
}

func DividedDifferences(points []Point) [][]float64 {
	ans := make([][]float64, len(points))
	for i := range ans {
		ans[i] = make([]float64, len(points))
		ans[i][0] = points[i].Y
	}
	for j := 1; j < len(ans); j++ {
		for i := 0; i < len(ans)-j; i++ {
			ans[i][j] = (ans[i+1][j-1] - ans[i][j-1]) / (points[i+j].X - points[i].X)
		}
	}
	return ans
}
