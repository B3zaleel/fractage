package fractals

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math"

	"github.com/B3zaleel/fractage/src/helpers"
	"github.com/llgcode/draw2d/draw2dimg"
)

// Properties of a Sierpinski triangle image.
type Hopalong struct {
	Width           int
	Height          int
	Color           color.RGBA
	UseRandomColors bool
	Iterations      int
	A               float64
	B               float64
	C               float64
	Resolution      int
	Background      color.RGBA
}

// Writes the Hopalong image to the given output.
func (props *Hopalong) WriteImage(output io.Writer) {
	viewport := image.Rect(0, 0, props.Width, props.Height)
	img := image.NewRGBA(viewport)
	gc := draw2dimg.NewGraphicContext(img)
	bounds := props.render(nil, float64(props.Resolution))
	xScale := float64(props.Width) / bounds.Width
	yScale := float64(props.Height) / bounds.Height
	scale := math.Min(xScale, yScale)
	helpers.FillImage(img, props.Background)
	props.render(gc, scale)
	err := png.Encode(output, img)
	if err != nil {
		panic(err)
	}
}

// Helper function for rendering the Hopalong.
func (props *Hopalong) render(gc *draw2dimg.GraphicContext, scale float64) helpers.Rect {
	x, y, xMax, yMax, t := 0.0, 0.0, 0.0, 0.0, 0.0
	midX, midY := float64(props.Width)/2.0, float64(props.Height)/2.0
	ptColor := props.Color
	if props.UseRandomColors && gc != nil {
		ptColor = helpers.RandomColor()
	}
	for i := 0; i < props.Iterations; i++ {
		for j := 0; j < props.Resolution; j++ {
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
				if props.UseRandomColors && i%50 == 0 && gc != nil {
					ptColor = helpers.RandomColor()
				}
				xMax = math.Max(xMax, x)
				yMax = math.Max(yMax, y)
				if gc != nil {
					helpers.PutPixel(gc, midX+x*scale, midY-y*scale, ptColor)
				}
			}
		}
	}
	bounds := helpers.Rect{
		X:      -xMax,
		Y:      -yMax,
		Width:  2 * xMax,
		Height: 2 * yMax,
	}
	return bounds
}
