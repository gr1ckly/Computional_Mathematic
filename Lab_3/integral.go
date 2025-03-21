package main

import (
	"errors"
	"math"
)

type Integral struct {
	fn    func(float64) float64
	start float64
	end   float64
}

func GetIntegral(num int) (*Integral, error) {
	var fn func(float64) float64
	switch num {
	case 1:
		fn = func(x float64) float64 {
			return -1*math.Pow(x, 3.0) - math.Pow(x, 2.0) - 2*x + 1
		}
	case 2:
		fn = func(x float64) float64 {
			return math.Pow(x, 3.0) - 3*math.Pow(x, 2.0) + 6*x - 19
		}
	case 3:
		fn = func(x float64) float64 {
			return math.Pow(x, 2.0) + math.Cos(x)
		}
	case 4:
		fn = func(x float64) float64 {
			return (3*x + 12) * math.Sin(x)
		}
	default:
		return nil, errors.New("Некорректно выбран интеграл")
	}
	return &Integral{fn: fn}, nil
}

func GetIntegralDescription() map[int]string {
	return map[int]string{
		1: "-x^3 - x^2 - 2x + 1",
		2: "x^3 - 3x^2 + 6x - 19",
		3: "x^2 + cos(x)",
		4: "(3x + 12)sin(x)",
	}
}
