package main

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func DrawGraphic(points []Point, x, lagrange, newton, gauss float64, filename string) error {
	p := plot.New()
	p.Title.Text = "Интерполяция функции"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	pts := make(plotter.XYs, len(points))
	for i, pt := range points {
		pts[i].X = pt.X
		pts[i].Y = pt.Y
	}
	scatter, err := plotter.NewScatter(pts)
	if err != nil {
		return err
	}
	scatter.GlyphStyle.Color = plotutil.Color(4)
	scatter.GlyphStyle.Radius = vg.Points(3)
	scatter.GlyphStyle.Shape = draw.CircleGlyph{}
	p.Add(scatter)
	p.Legend.Add("Исходные точки", scatter)
	minX, maxX := points[0].X, points[len(points)-1].X
	nPoints := 100
	step := (maxX - minX) / float64(nPoints-1)
	lagrangePts := make(plotter.XYs, nPoints)
	for i := 0; i < nPoints; i++ {
		xVal := minX + float64(i)*step
		lagrangePts[i].X = xVal
		lagrangePts[i].Y = LagrangeInterpolation(points, xVal)
	}
	lagrangeLine, err := plotter.NewLine(lagrangePts)
	if err != nil {
		return err
	}
	lagrangeLine.Color = plotutil.Color(1)
	p.Add(lagrangeLine)
	p.Legend.Add("Лагранж", lagrangeLine)
	newtonPts := make(plotter.XYs, nPoints)
	for i := 0; i < nPoints; i++ {
		xVal := minX + float64(i)*step
		newtonPts[i].X = xVal
		newtonPts[i].Y = NewtonInterpolation(points, xVal)
	}
	newtonLine, err := plotter.NewLine(newtonPts)
	if err != nil {
		return err
	}
	newtonLine.Color = plotutil.Color(2)
	p.Add(newtonLine)
	p.Legend.Add("Ньютон", newtonLine)
	gaussPts := make(plotter.XYs, nPoints)
	for i := 0; i < nPoints; i++ {
		xVal := minX + float64(i)*step
		gaussPts[i].X = xVal
		gaussPts[i].Y = GaussInterpolation(points, xVal)
	}
	gaussLine, err := plotter.NewLine(gaussPts)
	if err != nil {
		return err
	}
	gaussLine.Color = plotutil.Color(3)
	p.Add(gaussLine)
	p.Legend.Add("Гаусс", gaussLine)
	lagrangeScatter, err := plotter.NewScatter(plotter.XYs{{X: x, Y: lagrange}})
	if err != nil {
		return err
	}
	lagrangeScatter.GlyphStyle.Color = plotutil.Color(1)
	lagrangeScatter.GlyphStyle.Shape = draw.CircleGlyph{}
	lagrangeScatter.GlyphStyle.Radius = vg.Points(4)
	p.Add(lagrangeScatter)
	p.Legend.Add(fmt.Sprintf("Лагранж (x=%v)", x), lagrangeScatter)
	newtonScatter, err := plotter.NewScatter(plotter.XYs{{X: x, Y: newton}})
	if err != nil {
		return err
	}
	newtonScatter.GlyphStyle.Color = plotutil.Color(2)
	newtonScatter.GlyphStyle.Shape = draw.SquareGlyph{}
	newtonScatter.GlyphStyle.Radius = vg.Points(4)
	p.Add(newtonScatter)
	p.Legend.Add(fmt.Sprintf("Ньютон (x=%v)", x), newtonScatter)
	gaussScatter, err := plotter.NewScatter(plotter.XYs{{X: x, Y: gauss}})
	if err != nil {
		return err
	}
	gaussScatter.GlyphStyle.Color = plotutil.Color(3)
	gaussScatter.GlyphStyle.Shape = draw.TriangleGlyph{}
	gaussScatter.GlyphStyle.Radius = vg.Points(6)
	p.Add(gaussScatter)
	p.Legend.Add(fmt.Sprintf("Гаусс (x=%v)", x), gaussScatter)
	err = p.Save(10*vg.Inch, 6*vg.Inch, filename)
	return err
}
