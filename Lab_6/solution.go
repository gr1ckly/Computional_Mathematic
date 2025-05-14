package main

import (
	"errors"
	"fmt"
	"math"
)

func Euler(d Diffur) (Diffur, error) {
	h := d.x[1] - d.x[0]
	result := Diffur{
		fn:       d.fn,
		exactFn:  d.exactFn,
		x:        d.x,
		accuracy: d.accuracy,
		y:        make([]float64, len(d.x)),
	}
	result.y[0] = d.y[0]
	for i := 1; i < len(result.x); i++ {
		result.y[i] = result.y[i-1] + h*d.fn(result.x[i-1], result.y[i-1])
	}
	return result, nil
}

func RungeKutta4(d Diffur) (Diffur, error) {
	if len(d.x) < 2 {
		return Diffur{}, errors.New("Недостаточно точек x для вычисления")
	}
	h := d.x[1] - d.x[0]
	result := Diffur{
		fn:       d.fn,
		exactFn:  d.exactFn,
		x:        d.x,
		accuracy: d.accuracy,
		y:        make([]float64, len(d.x)),
	}
	result.y[0] = d.y[0]
	for i := 1; i < len(result.x); i++ {
		k1 := h * d.fn(result.x[i-1], result.y[i-1])
		k2 := h * d.fn(result.x[i-1]+h/2, result.y[i-1]+k1/2)
		k3 := h * d.fn(result.x[i-1]+h/2, result.y[i-1]+k2/2)
		k4 := h * d.fn(result.x[i-1]+h, result.y[i-1]+k3)
		result.y[i] = result.y[i-1] + (k1+2*k2+2*k3+k4)/6
	}
	return result, nil
}

func Milne(d Diffur) (Diffur, float64, error) {
	if len(d.x) < 5 {
		return Diffur{}, 0, errors.New("Метод Милна требует минимум 5 точек")
	}
	h := d.x[1] - d.x[0]
	result := Diffur{
		fn:       d.fn,
		exactFn:  d.exactFn,
		x:        d.x,
		accuracy: d.accuracy,
		y:        make([]float64, len(d.x)),
	}
	rk, err := RungeKutta4(Diffur{
		fn:       d.fn,
		exactFn:  d.exactFn,
		x:        d.x[:4],
		accuracy: d.accuracy,
		y:        d.y[:4],
	})
	if err != nil {
		return Diffur{}, 0, fmt.Errorf("Ошибка в подготовке начальных точек: %v", err)
	}
	copy(result.y[:4], rk.y[:4])
	for i := 4; i < len(result.x); i++ {
		pred := result.y[i-4] + 4*h*(2*d.fn(result.x[i-3], result.y[i-3])-d.fn(result.x[i-2], result.y[i-2])+2*d.fn(result.x[i-1], result.y[i-1]))/3
		result.y[i] = result.y[i-2] + h*(d.fn(result.x[i-2], result.y[i-2])+4*d.fn(result.x[i-1], result.y[i-1])+d.fn(result.x[i], pred))/3
	}
	maxError := 0.0
	for i, x := range result.x {
		exactY := d.exactFn(x)
		currErr := math.Abs(result.y[i] - exactY)
		if currErr > maxError {
			maxError = currErr
		}
	}

	return result, maxError, nil
}

func estimateRungeError(d Diffur, method func(Diffur) (Diffur, error), coef float64) (float64, error) {
	if len(d.x) < 2 {
		return 0, errors.New("Недостаточно точек для оценки")
	}
	resH, err := method(d)
	if err != nil {
		return 0, err
	}
	var xHalf []float64
	h := d.x[1] - d.x[0]
	for i := 0; i < len(d.x)-1; i++ {
		xHalf = append(xHalf, d.x[i])
		xHalf = append(xHalf, d.x[i]+h/2)
	}
	xHalf = append(xHalf, d.x[len(d.x)-1])
	dHalf := Diffur{
		fn:       d.fn,
		x:        xHalf,
		accuracy: d.accuracy,
		y:        d.y,
		exactFn:  d.exactFn,
	}
	resHalf, err := method(dHalf)
	if err != nil {
		return 0, err
	}
	var yHalfInPoints []float64
	for _, x := range d.x {
		for i, xh := range resHalf.x {
			if math.Abs(xh-x) < 1e-9 {
				yHalfInPoints = append(yHalfInPoints, resHalf.y[i])
				break
			}
		}
	}
	if len(yHalfInPoints) != len(resH.y) {
		return 0, errors.New("несоответствие размеров")
	}
	maxError := 0.0
	for i := range resH.y {
		error := math.Abs(resH.y[i] - yHalfInPoints[i])
		if error > maxError {
			maxError = error
		}
	}
	return maxError / coef, nil
}
