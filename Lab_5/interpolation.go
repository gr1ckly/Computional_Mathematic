package main

func LagrangeInterpolation(points []Point, x float64) float64 {
	ans := 0.0
	for i := 0; i < len(points); i++ {
		curr := points[i].Y
		for j := 0; j < len(points); j++ {
			if i != j {
				curr *= (x - points[j].X) / (points[i].X - points[j].X)
			}
		}
		ans += curr
	}
	return ans
}

func NewtonInterpolation(points []Point, x float64) float64 {
	table := DividedDifferences(points)
	ans := table[0][0]
	curr := 1.0
	for i := 1; i < len(points); i++ {
		curr *= (x - points[i-1].X)
		ans += table[0][i] * curr
	}
	return ans
}

func GaussInterpolation(points []Point, x0 float64) float64 {
	h := points[1].X - points[0].X
	t := (x0 - points[len(points)/2].X) / h
	table := FiniteDifferences(points)
	ans := table[len(points)/2][0]
	curr := 1.0
	fact := 1.0
	isForward := x0 >= points[len(points)/2].X
	for i := 1; i < len(points); i++ {
		var idx int
		fact *= float64(i)
		if isForward {
			if i%2 == 1 {
				curr *= (t + float64(i/2))
			} else {
				curr *= (t - float64(i/2))
			}
			idx = len(points)/2 - (i / 2)
		} else {
			if i%2 == 1 {
				curr *= (t - float64(i/2))
			} else {
				curr *= (t + float64(i/2))
			}
			idx = len(points)/2 - ((i + 1) / 2)
		}
		ans += curr * table[idx][i] / fact
	}
	return ans
}
