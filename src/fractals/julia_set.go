package fractals

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math"
	"math/cmplx"

	"github.com/B3zaleel/fractage/src/helpers"
	"github.com/llgcode/draw2d/draw2dimg"
)

// Properties of a Julia set image.
type JuliaSet struct {
	Width         int
	Height        int
	ColorPalette  helpers.ColorPalette
	MaxIterations int
	C             complex128
	BailOut       float64
	Region        helpers.Rect
	Background    color.RGBA
}

// Writes the Julia set image to the given output.
func (props *JuliaSet) WriteImage(output io.Writer) error {
	viewport := image.Rect(0, 0, props.Width, props.Height)
	img := image.NewRGBA(viewport)
	gc := draw2dimg.NewGraphicContext(img)
	helpers.FillImage(img, props.Background)
	err := props.render(gc)
	if err != nil {
		return err
	}
	err = png.Encode(output, img)
	if err != nil {
		return err
	}
	return nil
}

// Helper function for rendering the Julia set.
func (props *JuliaSet) render(gc *draw2dimg.GraphicContext) error {
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
			tmp := math.Exp(-cmplx.Abs(Z))
			for (n < props.MaxIterations) && (cmplx.Abs(Z) < props.BailOut) {
				Z = Z*Z + props.C
				tmp += math.Exp(-cmplx.Abs(Z))
				n++
			}
			if n < props.MaxIterations {
				pixelColor, err = props.ColorPalette.GetColor(tmp / float64(props.MaxIterations))
				if err != nil {
					return err
				}
			} else {
				pixelColor, err = props.ColorPalette.GetColor(1)
				if err != nil {
					return err
				}
			}
			helpers.PutPixel(gc, float64(x), float64(y), pixelColor)
		}
	}
	return nil
}