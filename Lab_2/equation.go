package main

import (
	"errors"
	"math"
)

type Equation struct {
	fn               func(float64) float64
	diff             func(float64) float64
	fi               func(float64) float64
	checkConvergence func(float64, float64) bool
	checkRoots       func(float64, float64) bool
	roots            []float64
}

func GetEquation(number int) (*Equation, error) {
	switch number {
	case 1:
		{
			roots := []float64{-1.15624, 0.629971, 2.41627}
			fn := func(x float64) float64 {
				return math.Pow(x, 3.0) - 1.89*math.Pow(x, 2.0) - 2.0*x + 1.76
			}
			diff := func(x float64) float64 {
				return 2*math.Pow(x, 2.0) - 3.78*x - 2
			}
			fi := func(x float64) float64 {
				return math.Pow(x, 3.0)*0.5 - 0.945*math.Pow(x, 2.0) + 1.76
			}
			check := func(a float64, b float64) bool {
				if a > b || a <= -1*math.Sqrt(2/3*2.89) || b >= math.Sqrt(2/3*2.89) {
					return false
				}
				return true
			}
			checkRoot := func(a float64, b float64) bool {
				countRoots := 0
				for _, x := range roots {
					if a <= x && b >= x {
						countRoots += 1
					}
				}
				if countRoots != 1 {
					return false
				}
				return true
			}
			return &Equation{fn, diff, fi, check, checkRoot, roots}, nil
		}
	case 2:
		roots := []float64{-1.24967, 0.508787, 3.86589}
		fn := func(x float64) float64 {
			return math.Pow(x, 3.0) - 3.125*math.Pow(x, 2.0) - 3.5*x + 2.458
		}
		diff := func(x float64) float64 {
			return 3*math.Pow(x, 2.0) - 6.25*x - 3.5
		}
		fi := func(x float64) float64 {
			return (math.Pow(x, 3.0) - 3.125*math.Pow(x, 2.0) + 2.458) / 3.5
		}
		check := func(a float64, b float64) bool {
			if a > b || a <= -0.46 || b >= 2.54 {
				return false
			}
			return true
		}
		checkRoot := func(a float64, b float64) bool {
			countRoots := 0
			for _, x := range roots {
				if a <= x && b >= x {
					countRoots += 1
				}
			}
			if countRoots != 1 {
				return false
			}
			return true
		}
		return &Equation{fn, diff, fi, check, checkRoot, roots}, nil
	case 3:
		{
			roots := []float64{-1.35672, 0.254703, 2.47424}
			fn := func(x float64) float64 {
				return 1.8*math.Pow(x, 3.0) - 2.47*math.Pow(x, 2.0) - 5.53*x + 1.539
			}
			diff := func(x float64) float64 {
				return 5.4*math.Pow(x, 2.0) - 4.94*x - 5.53
			}
			fi := func(x float64) float64 {
				return 1.8/5.53*math.Pow(x, 3.0) - 2.47/5.53*math.Pow(x, 2.0) + 1.539/5.53
			}
			check := func(a float64, b float64) bool {
				if a > b || a <= -1.31 || b >= 3.14 {
					return true
				}
				return true
			}
			checkRoot := func(a float64, b float64) bool {
				countRoots := 0
				for _, x := range roots {
					if a <= x && b >= x {
						countRoots += 1
					}
				}
				if countRoots != 1 {
					return false
				}
				return true
			}
			return &Equation{fn, diff, fi, check, checkRoot, roots}, nil
		}
	}
	return nil, errors.New("Неверный номер уравнения")
}
func GetEqMap() map[int]string {
	return map[int]string{
		1: "x^3 - 1.89x^2 - 2x + 1.76 = 0",
		2: "x^3 - 3.125x^2 - 3.5x + 2.458 = 0",
		3: "1.8x^3 - 2.47x^2 - 5.53x + 1.539 = 0",
	}
}
