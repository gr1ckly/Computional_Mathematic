package main

import (
	"errors"
	"math"
)

func Bisection(eq *Equation, a, b, eps float64) (float64, int, error) {
	if eq.fn(a)*eq.fn(b) > 0 {
		return 0, 0, errors.New("Функция на концах отрезка имеет одинаковый знак.")
	}
	var c float64
	iterations := 0
	for (b-a)/2 > eps {
		c = (a + b) / 2
		if eq.fn(c) == 0 {
			break
		} else if eq.fn(a)*eq.fn(c) < 0 {
			b = c
		} else {
			a = c
		}
		iterations++
	}
	return c, iterations, nil
}

func Newton(eq *Equation, x0, eps float64) (float64, int, error) {
	x := x0
	iterations := 0

	for math.Abs(eq.fn(x)) > eps {
		if eq.diff(x) == 0 {
			return 0, 0, errors.New("Производная равна нулю, метод не применим")
		}
		x = x - eq.fn(x)/eq.diff(x)
		iterations++
	}
	return x, iterations, nil
}

func SimpleIteration(eq *Equation, x0, eps float64) (float64, int) {
	x := x0
	iterations := 0
	xPrev := x0 + 1
	for math.Abs(x-xPrev) > eps {
		xPrev = x
		x = eq.fi(xPrev)
		iterations++
	}
	return x, iterations
}

func NewtonMethod(sysEq *SystemEquation, x0, y0 float64, eps float64, maxIter int) (float64, float64, int, float64, error) {
	x, y := x0, y0
	for i := 1; i <= maxIter; i++ {
		F1 := sysEq.fn1(x, y)
		F2 := sysEq.fn2(x, y)
		diff11 := sysEq.diff11(x, y)
		diff12 := sysEq.diff12(x, y)
		diff21 := sysEq.diff21(x, y)
		diff22 := sysEq.diff22(x, y)
		detJ := diff11*diff22 - diff12*diff21
		if math.Abs(detJ) < 1e-10 {
			return 0, 0, i, 0, errors.New("Якобиан вырожден, метод Ньютона не применим")
		}
		dx := (F2*diff12 - F1*diff22) / detJ
		dy := (F1*diff21 - F2*diff11) / detJ
		xNew := x + dx
		yNew := y + dy
		errX := math.Abs(xNew - x)
		errY := math.Abs(yNew - y)
		errorVal := math.Max(math.Abs(dx), math.Abs(dy))
		if errX < eps && errY < eps {
			return xNew, yNew, i, errorVal, nil
		}
		x, y = xNew, yNew
	}

	return x, y, maxIter, math.Max(math.Abs(x-x0), math.Abs(y-y0)), errors.New("Метод не сошелся за указанное число итераций")
}
