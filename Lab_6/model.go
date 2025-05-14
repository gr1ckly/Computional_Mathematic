package main

import (
	"errors"
	"math"
)

var fns = map[string]func(float64, float64) float64{
	"1": func(x, y float64) float64 {
		return x + y
	},
	"2": func(x, y float64) float64 {
		return x*x - 2*y
	},
	"3": func(x, y float64) float64 {
		return math.Sin(x) + y
	},
}

var fnsDescription = map[string]string{
	"1": "x + y",
	"2": "x^2 - 2y",
	"3": "sin(x) + y",
}

func GetExactSolutionFunction(eqChoice string, x0, y0 float64) (func(float64) float64, error) {
	switch eqChoice {
	case "1":
		C := (y0 + x0 + 1) * math.Exp(-x0)
		return func(x float64) float64 {
			return C*math.Exp(x) - x - 1
		}, nil
	case "2":
		C := (y0 - (2*x0*x0-2*x0+1)/4) / math.Exp(-2*x0)
		return func(x float64) float64 {
			return C*math.Exp(-2*x) + (2*x*x-2*x+1)/4
		}, nil
	case "3":
		C := (y0 + (math.Sin(x0)+math.Cos(x0))/2) / math.Exp(x0)
		return func(x float64) float64 {
			return C*math.Exp(x) - (math.Sin(x)+math.Cos(x))/2
		}, nil
	default:
		return nil, errors.New("неизвестное уравнение")
	}
}

type Diffur struct {
	fn       func(float64, float64) float64
	exactFn  func(float64) float64
	x        []float64
	y        []float64
	accuracy float64
}

func BuildDiffur(fn func(float64, float64) float64, exactFn func(float64) float64, x []float64, y0 float64, accuracy float64) Diffur {
	y := make([]float64, len(x))
	y[0] = y0
	return Diffur{fn, exactFn, x, y, accuracy}
}
