package main

import (
	"errors"
	"math"
)

func getPartition(integral *Integral, start float64, h float64, count int) []float64 {
	yArr := make([]float64, count)
	for i := 0; i < count; i++ {
		yArr[i] = integral.fn(start + float64(i)*h)
	}
	return yArr
}

func countH(integral *Integral, iter int) float64 {
	return (integral.end - integral.start) / float64(iter)
}

func rectangle(yArr []float64, h float64) float64 {
	i0 := 0.0
	for _, y := range yArr {
		i0 += y
	}
	return i0 * h
}

func SolveIntegral(integral *Integral, accuracy float64, fn func(*Integral, int) float64) (float64, int) {
	iter := 4
	i0 := fn(integral, iter)
	iter *= 2
	i1 := fn(integral, iter)
	for math.Abs(i1-i0) >= accuracy {
		iter *= 2
		i0 = i1
		i1 = fn(integral, iter)
	}
	return i1, iter
}

func LeftRectangle(integral *Integral, startIter int) float64 {
	h := countH(integral, startIter)
	yArr := getPartition(integral, integral.start, h, startIter)
	return rectangle(yArr, h)
}

func RightRectangle(integral *Integral, iter int) float64 {
	h := countH(integral, iter)
	yArr := getPartition(integral, integral.start+h, h, iter)
	return rectangle(yArr, h)
}

func MidRectangle(integral *Integral, iter int) float64 {
	h := countH(integral, iter)
	yArr := getPartition(integral, integral.start+h/2, h, iter)
	return rectangle(yArr, h)
}

func Trapeze(integral *Integral, iter int) float64 {
	h := countH(integral, iter)
	yArr := getPartition(integral, integral.start, h, iter+1)
	i0 := yArr[0] + yArr[len(yArr)-1]
	i0 /= 2
	for i := 1; i < len(yArr)-1; i++ {
		i0 += yArr[i]
	}
	return i0 * h
}

func Simpson(integral *Integral, iter int) float64 {
	h := countH(integral, iter)
	yArr := getPartition(integral, integral.start, h, iter+1)
	i0 := yArr[0] + yArr[len(yArr)-1]
	for i := 1; i < len(yArr)-1; i += 2 {
		i0 += 4 * yArr[i]
	}
	for i := 2; i < len(yArr)-1; i += 2 {
		i0 += 2 * yArr[i]
	}
	return i0 * h / 3
}

func GetMethodDescription() map[int]string {
	return map[int]string{
		1: "Метод левых прямоугольников",
		2: "Метод правых прямоугольников",
		3: "Метод средних прямоугольников",
		4: "Метод трапеции",
		5: "Метод Симпсона",
	}
}

func GetMethod(num int) (func(*Integral, int) float64, error) {
	var fn func(*Integral, int) float64
	switch num {
	case 1:
		fn = LeftRectangle
	case 2:
		fn = RightRectangle
	case 3:
		fn = MidRectangle
	case 4:
		fn = Trapeze
	case 5:
		fn = Simpson
	default:
		return nil, errors.New("Метод выбран некорректно")
	}
	return fn, nil
}
