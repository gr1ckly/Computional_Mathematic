package main

import "math"

type SystemEquation struct {
	fn1    func(float64, float64) float64
	fn2    func(float64, float64) float64
	diff11 func(float64, float64) float64
	diff12 func(float64, float64) float64
	diff21 func(float64, float64) float64
	diff22 func(float64, float64) float64
}

func GetSystemEquation(number int) (*SystemEquation, error) {
	switch number {
	case 1:
		{
			fn1 := func(x float64, y float64) float64 {
				return math.Tan(x*y) - x*x
			}
			fn2 := func(x float64, y float64) float64 {
				return 0.5*x*x + 2*y*y - 1
			}
			diff11 := func(x float64, y float64) float64 {
				return y/math.Pow(math.Cos(x*y), 2.0) - 2*x
			}
			diff12 := func(x float64, y float64) float64 {
				return x / math.Pow(math.Cos(x*y), 2.0)
			}
			diff21 := func(x float64, y float64) float64 {
				return x
			}
			diff22 := func(x float64, y float64) float64 {
				return 4 * y
			}
			return &SystemEquation{fn1, fn2, diff11, diff12, diff21, diff22}, nil
		}
	case 2:
		{
			fn1 := func(x float64, y float64) float64 {
				return x + math.Sin(y) + 0.4
			}
			fn2 := func(x float64, y float64) float64 {
				return 2*y - math.Cos(x+1)
			}
			diff11 := func(x float64, y float64) float64 {
				return 1
			}
			diff12 := func(x float64, y float64) float64 {
				return math.Cos(y)
			}
			diff21 := func(x float64, y float64) float64 {
				return math.Sin(x + 1)
			}
			diff22 := func(x float64, y float64) float64 {
				return 2
			}
			return &SystemEquation{fn1, fn2, diff11, diff12, diff21, diff22}, nil
		}
	}
	return nil, nil
}

func GetSysEqMap() map[int]string {
	return map[int]string{
		1: "| tg(xy) = x^2\n" +
			"<\n" +
			"| 0.5x^2 + 2y^2 = 1",
		2: "| x + sin(y) = -0.4\n" +
			"<\n" +
			"| 2y - cos(x + 1) = 0",
	}
}
