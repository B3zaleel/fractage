package fractals

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math"

	"github.com/B3zaleel/fractage/src/helpers"
)

// Properties of a Hopalong image.
type Hopalong struct {
	Width           int
	Height          int
	Color           color.RGBA
	UseRandomColors bool
	A               float64
	B               float64
	C               float64
	D               float64
	X               float64
	Y               float64
	Scale           float64
	Resolution      int
	Background      color.RGBA
}

// Writes the Hopalong image to the given output.
func (props *Hopalong) WriteImage(output io.Writer) {
	viewport := image.Rect(0, 0, props.Width, props.Height)
	img := image.NewRGBA(viewport)
	helpers.FillImage(img, props.Background)
	props.render(img)
	err := png.Encode(output, img)
	if err != nil {
		panic(err)
	}
}

// Helper function for rendering the Hopalong.
func (props *Hopalong) render(img *image.RGBA) {
	x, y, t := props.X, props.Y, 0.0
	midX, midY := float64(props.Width)/2.0, float64(props.Height)/2.0
	ptColor := props.Color
	if props.UseRandomColors {
		ptColor = helpers.RandomColor()
	}
	for i := 0; i < props.Width; i++ {
		for j := 0; j < props.Height; j++ {
			for k := 0; k < props.Resolution; k++ {
				xSign := 0
				if x < 0 {
					xSign = -1
				} else if x > 0 {
					xSign = 1
				}
				t = x
				x = y - float64(xSign)*math.Sqrt(math.Abs(float64(props.B)*x-float64(props.C)))
				y = float64(props.A) - t
				if props.UseRandomColors && i%50 == 0 {
					ptColor = helpers.RandomColor()
				}
				img.Set(int(midX+x*props.Scale), int(midY-y*props.Scale), ptColor)
			}
		}
	}
}
