package main

import (
	"fmt"
	"git.sr.ht/~sbinet/gg"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
	"log"
	"math"
	"strconv"
)

const width, height = 600, 600

func DrawGraphic(cancel <-chan struct{}, eq ...*Equation) {
	dc := gg.NewContext(width, height)
	dc.SetColor(color.White)
	dc.Clear()
	dc.SetColor(color.Black)
	dc.SetLineWidth(1)
	dc.DrawLine(0, height/2, width, height/2)
	dc.DrawLine(width/2, 0, width/2, height)
	dc.Stroke()
	const tickCount = 10
	xStep := float64(width) / float64(tickCount)
	yStep := float64(height) / float64(tickCount)
	for i := 0; i <= tickCount; i++ {
		x := float64(i) * xStep
		label := strconv.Itoa(i - tickCount/2)
		dc.DrawStringAnchored(label, x, height/2+15, 0.5, 0.5)
	}
	for i := 0; i <= tickCount; i++ {
		y := float64(i) * yStep
		label := strconv.Itoa(tickCount/2 - i)
		dc.DrawStringAnchored(label, width/2-15, y, 0.5, 0.5)
	}
	colors := []color.Color{
		color.RGBA{255, 0, 0, 255},
		color.RGBA{0, 255, 0, 255},
		color.RGBA{0, 0, 255, 255},
		color.RGBA{255, 165, 0, 255},
	}
	for i, equation := range eq {
		dc.SetColor(colors[i%len(colors)])
		dc.SetLineWidth(2)
		prevX, prevY := 0.0, 0.0
		firstPoint := true
		for px := 1; px <= width; px++ {
			x := (float64(px) * 10 / float64(width)) - 5
			y := equation.fn(x)
			if math.IsNaN(y) || math.IsInf(y, 0) {
				firstPoint = true
				continue
			}
			screenX := float64(px)
			screenY := height/2 - (y * 40)
			if firstPoint {
				prevX, prevY = screenX, screenY
				firstPoint = false
				continue
			}
			dc.DrawLine(prevX, prevY, screenX, screenY)
			prevX, prevY = screenX, screenY
		}
		dc.Stroke()
		dc.SetColor(color.RGBA{0, 0, 0, 255})
		for _, root := range equation.roots {
			px := int(((root + 5) / 10) * float64(width))
			py := int(height / 2)
			dc.DrawCircle(float64(px), float64(py), 3)
			dc.Fill()
			label := fmt.Sprintf("(%.2f, 0)", root)
			dc.SetColor(color.Black)
			dc.DrawStringAnchored(label, float64(px)+10, float64(py)-10, 0.5, 0.5)
		}
	}

	pixelgl.Run(func() {
		cfg := pixelgl.WindowConfig{
			Title:  "График уравнения",
			Bounds: pixel.R(0, 0, width, height),
			VSync:  true,
		}
		win, err := pixelgl.NewWindow(cfg)
		if err != nil {
			log.Fatal(err)
			return
		}
		pic := pixel.PictureDataFromImage(dc.Image())
		sprite := pixel.NewSprite(pic, pic.Bounds())
		for !win.Closed() {
			select {
			case <-cancel:
				win.SetClosed(true)
				break
			default:
				win.Clear(color.White)
				sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
				win.Update()
			}
		}
	})
}

func DrawSystemGraphic(cancel <-chan struct{}, eq *SystemEquation) {
	dc := gg.NewContext(width, height)
	dc.SetColor(color.White)
	dc.Clear()

	dc.SetColor(color.Black)
	dc.SetLineWidth(1)
	dc.DrawLine(0, height/2, width, height/2)
	dc.DrawLine(width/2, 0, width/2, height)
	dc.Stroke()

	const tickCount = 10
	xStep := float64(width) / float64(tickCount)
	yStep := float64(height) / float64(tickCount)
	for i := 0; i <= tickCount; i++ {
		x := float64(i) * xStep
		label := strconv.Itoa(i - tickCount/2)
		dc.DrawStringAnchored(label, x, height/2+15, 0.5, 0.5)
	}

	for i := 0; i <= tickCount; i++ {
		y := float64(i) * yStep
		label := strconv.Itoa(tickCount/2 - i)
		dc.DrawStringAnchored(label, width/2-15, y, 0.5, 0.5)
	}

	colors := []color.Color{
		color.RGBA{255, 0, 0, 255},
		color.RGBA{0, 0, 255, 255},
	}

	dc.SetColor(colors[0])
	dc.SetLineWidth(1)
	for px := 0; px < width; px++ {
		for py := 0; py < height; py++ {
			x := (float64(px)/width)*10 - 5
			y := (float64(py)/height)*10 - 5
			z := eq.fn1(x, y)
			if z > -0.05 && z < 0.05 {
				dc.SetPixel(px, height-int((y+5)/10*float64(height)))
			}
		}
	}
	dc.Stroke()

	dc.SetColor(colors[1])
	dc.SetLineWidth(1)
	for px := 0; px < width; px++ {
		for py := 0; py < height; py++ {
			x := (float64(px)/width)*10 - 5
			y := (float64(py)/height)*10 - 5
			z := eq.fn2(x, y)
			if z > -0.05 && z < 0.05 {
				dc.SetPixel(px, height-int((y+5)/10*float64(height)))
			}
		}
	}
	dc.Stroke()

	dc.SetColor(color.Black)
	dc.SetLineWidth(2)
	for _, root := range eq.roots {
		px := int(((root[0] + 5) / 10) * float64(width))
		py := height - int(((root[1]+5)/10)*float64(height))
		dc.DrawCircle(float64(px), float64(py), 3)
		dc.Fill()

		label := fmt.Sprintf("(%.2f, %.2f)", root[0], root[1])
		dc.SetColor(color.Black)
		dc.DrawStringAnchored(label, float64(px)+8, float64(py)-8, 0, 0)
	}

	pixelgl.Run(func() {
		cfg := pixelgl.WindowConfig{
			Title:  "График системы нелинейных уравнений",
			Bounds: pixel.R(0, 0, width, height),
			VSync:  true,
		}

		win, err := pixelgl.NewWindow(cfg)
		if err != nil {
			log.Fatal(err)
			return
		}

		pic := pixel.PictureDataFromImage(dc.Image())
		sprite := pixel.NewSprite(pic, pic.Bounds())

		for !win.Closed() {
			select {
			case <-cancel:
				win.SetClosed(true)
				return
			default:
				win.Clear(color.White)
				sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
				win.Update()
			}
		}
	})
}
