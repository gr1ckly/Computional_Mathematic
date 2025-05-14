package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func DrawGraphic(euler, rungeKutta, milne Diffur, filename string) error {
	p := plot.New()
	p.Title.Text = "Решение дифференциального уравнения"
	p.X.Label.Text = "x"
	p.Y.Label.Text = "y"
	makeLine := func(x, y []float64) plotter.XYs {
		pts := make(plotter.XYs, len(x))
		for i := range x {
			pts[i].X = x[i]
			pts[i].Y = y[i]
		}
		return pts
	}
	eulerLine := makeLine(euler.x, euler.y)
	rkLine := makeLine(rungeKutta.x, rungeKutta.y)
	milneLine := makeLine(milne.x, milne.y)
	var exactLine plotter.XYs
	if euler.exactFn != nil {
		exactLine = make(plotter.XYs, len(euler.x))
		for i, x := range euler.x {
			exactLine[i].X = x
			exactLine[i].Y = euler.exactFn(x)
		}
	}
	err := plotutil.AddLines(p,
		"Эйлер", eulerLine,
		"Рунге-Кутта", rkLine,
		"Милн", milneLine,
	)
	if err != nil {
		return err
	}
	if euler.exactFn != nil {
		exact, err := plotter.NewLine(exactLine)
		if err != nil {
			return err
		}
		exact.Color = plotutil.Color(3)
		exact.Width = vg.Points(2)
		exact.Dashes = []vg.Length{vg.Points(5), vg.Points(3)}
		p.Add(exact)
		p.Legend.Add("Точное решение", exact)
	}
	return p.Save(6*vg.Inch, 4*vg.Inch, filename)
}
