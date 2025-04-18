package main

import (
	"math"
	"sort"
)

func CalcStdDeviation(points []Point, f func(float64, []float64) float64, coeffs []float64) float64 {
	var s float64
	for _, p := range points {
		yPred := f(p.X, coeffs)
		s += math.Pow(p.Y-yPred, 2)
	}
	return math.Sqrt(s / float64(len(points)))
}

func CalcDetermination(points []Point, f func(float64, []float64) float64, coeffs []float64) float64 {
	var sTot, sRes float64
	var meanY float64
	for _, p := range points {
		meanY += p.Y
	}
	meanY /= float64(len(points))

	for _, p := range points {
		yPred := f(p.X, coeffs)
		sTot += math.Pow(p.Y-meanY, 2)
		sRes += math.Pow(p.Y-yPred, 2)
	}
	return 1 - (sRes / sTot)
}

func PearsonCorrelation(points []Point) float64 {
	var sumX, sumY float64
	n := float64(len(points))
	for _, p := range points {
		sumX += p.X
		sumY += p.Y
	}
	meanX := sumX / n
	meanY := sumY / n
	var num, denomX, denomY float64
	for _, p := range points {
		dx := p.X - meanX
		dy := p.Y - meanY
		num += dx * dy
		denomX += dx * dx
		denomY += dy * dy
	}
	if denomX == 0 || denomY == 0 {
		return 0
	}
	return num / math.Sqrt(denomX*denomY)
}

func LeastSquaresMethod(points []Point, approximate *ApproxFunc) float64 {
	sum := 0.0
	for _, point := range points {
		sum += math.Pow(approximate.Func(point.X, approximate.Coefficients)-point.Y, 2.0)
	}
	return sum
}

func GetBestApproximation(points []Point, approximations []*ApproxFunc) (float64, []*ApproxFunc) {
	mp := make(map[float64][]*ApproxFunc)
	keys := []float64{}
	for _, appr := range approximations {
		num := LeastSquaresMethod(points, appr)
		mp[num] = append(mp[num], appr)
		keys = append(keys, num)
	}
	sort.Float64s(keys)
	return keys[0], mp[keys[0]]
}
