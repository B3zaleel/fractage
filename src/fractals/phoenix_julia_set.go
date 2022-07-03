package fractals

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math"
	"math/cmplx"

	"github.com/B3zaleel/fractage/src/helpers"
)

// Properties of a phoenix Julia set image.
type PhoenixJuliaSet struct {
	Width         int
	Height        int
	ColorPalette  helpers.ColorPalette
	MaxIterations int
	C             complex128
	K             complex128
	BailOut       float64
	Region        helpers.Rect
	Background    color.RGBA
}

// Writes the phoenix Julia set image to the given output.
func (props *PhoenixJuliaSet) WriteImage(output io.Writer) error {
	viewport := image.Rect(0, 0, props.Width, props.Height)
	img := image.NewRGBA(viewport)
	helpers.FillImage(img, props.Background)
	err := props.render(img)
	if err != nil {
		return err
	}
	err = png.Encode(output, img)
	if err != nil {
		return err
	}
	return nil
}

// Helper function for rendering the phoenix Julia set.
func (props *PhoenixJuliaSet) render(img *image.RGBA) error {
	width, height := float64(props.Width), float64(props.Height)
	step := math.Max(props.Region.Width/width, props.Region.Height/height)
	xOffset := props.Region.X - (width*step-props.Region.Width)/2.0
	yOffset := props.Region.Y - (height*step-props.Region.Height)/2.0
	err := props.ColorPalette.TranslateColorTransitions()
	if err != nil {
		return err
	}
	var pixelColor color.RGBA
	var n int
	for y := 0; y < int(height); y++ {
		for x := 0; x < int(width); x++ {
			n = 0
			Z := complex(xOffset+float64(x)*step, yOffset+float64(y)*step)
			ZPrev, ZNext := Z, Z
			seriesValue := math.Exp(-cmplx.Abs(ZNext))
			for (n < props.MaxIterations) && (cmplx.Abs(ZNext) < props.BailOut) {
				ZPrev = Z
				Z = ZNext
				ZNext = cmplx.Pow(Z, 2) + props.C + props.K*ZPrev
				seriesValue += math.Exp(-cmplx.Abs(ZNext))
				n++
			}
			if n < props.MaxIterations {
				pixelColor, err = props.ColorPalette.GetColor(seriesValue / float64(props.MaxIterations))
				if err != nil {
					return err
				}
			} else {
				pixelColor, err = props.ColorPalette.GetColor(1)
				if err != nil {
					return err
				}
			}
			img.Set(x, y, pixelColor)
		}
	}
	return nil
}
