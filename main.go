package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Line struct {
	X1, Y1    int
	Magnitude int
	Degrees   float64
	color.Color
}

func DrawLine(screen *ebiten.Image, x1, y1, x2, y2 float64, c color.Color) {
	if x2 < x1 {
		x1, x2 = x2, x1
		y1, y2 = y2, y1
	}
	Δx, Δy := x2-x1, y2-y1
	A, B, C := Δy, -Δx, Δx*y1-Δy*x1
	y := y1
	for x := x1; x < x2; x++ {
		s := A*x + B*y + C
		if s > 0 {
			y += 1
		}
		screen.Set(int(x), int(y), c)
	}
	// Δy := 1
}

func ToRadians(Degrees float64) float64 {
	return Degrees * math.Pi / float64(180)
}

type game struct{ l Line }

func (*game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }
func (g *game) Update() error {
	g.l.Degrees += 1
	return nil
}
func (g *game) Draw(screen *ebiten.Image) {
	x := float64(g.l.Magnitude) * math.Cos(ToRadians(g.l.Degrees))
	y := float64(g.l.Magnitude) * math.Sin(ToRadians(g.l.Degrees))
	x2, y2 := x+float64(g.l.X1), y+float64(g.l.Y1)
	DrawLine(screen, float64(g.l.X1), float64(g.l.Y1), x2, y2, g.l.Color)
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g := game{Line{320, 240, 100, 0, color.RGBA{255, 1, 1, 255}}}
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}