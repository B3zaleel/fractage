package helpers

import (
	"image"
	"image/color"

	"github.com/llgcode/draw2d/draw2dimg"
)

// Fills an image with color.
//  *image*: The image to fill.
//  *color*: The color to fill the image with.
func FillImage(image *image.RGBA, color color.RGBA) {
	width := image.Bounds().Dx()
	height := image.Bounds().Dy()
	for i := 0; i <= width; i++ {
		for j := 0; j <= height; j++ {
			image.SetRGBA(i, j, color)
		}
	}
}

// Draws a rectangle in an image.
//  *image*: The image to draw the rectangle in.
//  *x*: The horizontal offset of the rectangle.
//  *y*: The vertical offset of the rectangle.
//  *width*: The width of the rectangle.
//  *height*: The height of the rectangle.
//  *color*: The color to fill the rectangle with.
func DrawRectangle(gc *draw2dimg.GraphicContext, x, y, width, height float64, color color.RGBA) {
	gc.SetStrokeColor(color)
	gc.SetLineWidth(0.2)
	gc.BeginPath()
	gc.MoveTo(x, y)
	gc.LineTo(x+width, y)
	gc.LineTo(x+width, y+height)
	gc.LineTo(x, y+height)
	gc.LineTo(x, y)
	gc.Close()
	gc.Stroke()
}

// Draws a filled rectangle in an image.
//  *image*: The image to draw the rectangle in.
//  *x*: The horizontal offset of the rectangle.
//  *y*: The vertical offset of the rectangle.
//  *width*: The width of the rectangle.
//  *height*: The height of the rectangle.
//  *color*: The color to fill the rectangle with.
func FillRectangle(gc *draw2dimg.GraphicContext, x, y, width, height float64, color color.RGBA) {
	gc.SetFillColor(color)
	gc.SetStrokeColor(color)
	gc.SetLineWidth(0.2)
	gc.BeginPath()
	gc.MoveTo(x, y)
	gc.LineTo(x+width, y)
	gc.LineTo(x+width, y+height)
	gc.LineTo(x, y+height)
	gc.LineTo(x, y)
	gc.Close()
	gc.FillStroke()
}
