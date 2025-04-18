package main

import (
	"bufio"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func DrawApproximations(points []Point, approximations []*ApproxFunc, filename string, writer *bufio.Writer) error {
	p := plot.New()
	p.Title.Text = "Аппроксимации"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	pts := make(plotter.XYs, len(points))
	for i := range points {
		pts[i].X = points[i].X
		pts[i].Y = points[i].Y
	}
	scatter, _ := plotter.NewScatter(pts)
	scatter.GlyphStyle.Radius = vg.Points(3)
	p.Add(scatter)
	p.Legend.Add("Эксперимент", scatter)
	addFunc := func(name string, f func(float64, []float64) float64, coeffs []float64, colorIdx int) {
		line := plotter.NewFunction(func(x float64) float64 { return f(x, coeffs) })
		line.Width = vg.Points(1)
		line.Color = plotutil.Color(colorIdx)
		p.Add(line)
		p.Legend.Add(name, line)
	}
	for idx, approx := range approximations {
		coeffs, err := approx.Approximate(points)
		if err != nil {
			_, err = writer.WriteString(err.Error())
			if err != nil {
				return err
			}
			continue
		}
		approx.Coefficients = coeffs
		addFunc(approx.Name, approx.Func, coeffs, idx)
	}
	return p.Save(8*vg.Inch, 6*vg.Inch, filename)
}
