package controllers

import (
	"fmt"
	"image/color"
	"strconv"

	"github.com/B3zaleel/fractage/src/fractals"
	"github.com/B3zaleel/fractage/src/helpers"
	"github.com/kataras/iris/v12"
)

const (
	HOPALONG_MAX_ITERATIONS     = 2_000_000_000
	HOPALONG_MAX_RESOLUTION     = 15
	HOPALONG_DEFAULT_ITERATIONS = 10_000_000
	HOPALONG_DEFAULT_RESOLUTION = 2
	HOPALONG_DEFAULT_A          = 2
	HOPALONG_DEFAULT_B          = 1
	HOPALONG_DEFAULT_C          = 0
)

func GetHopalong(ctx iris.Context) {
	query := ctx.Request().URL.Query()
	fractal := fractals.Hopalong{
		Width:           DEFAULT_WIDTH,
		Height:          DEFAULT_HEIGHT,
		A:               HOPALONG_DEFAULT_A,
		B:               HOPALONG_DEFAULT_B,
		C:               HOPALONG_DEFAULT_C,
		UseRandomColors: true,
		Iterations:      HOPALONG_DEFAULT_ITERATIONS,
		Resolution:      HOPALONG_DEFAULT_RESOLUTION,
		Background:      color.RGBA{255, 255, 255, 255},
	}
	if query.Has("width") {
		width, err := strconv.Atoi(query.Get("width"))
		if err != nil {
			ctx.Text(err.Error())
			return
		}
		fractal.Width = width
	}
	if query.Has("height") {
		height, err := strconv.Atoi(query.Get("height"))
		if err != nil {
			ctx.Text(err.Error())
			return
		}
		fractal.Height = height
	}
	if query.Has("color") {
		color, err := helpers.ParseColor(query.Get("color"))
		if err != nil {
			ctx.Text(err.Error())
			return
		}
		fractal.Color = color
		fractal.UseRandomColors = false
	}
	if query.Has("iterations") {
		iterations, err := strconv.Atoi(query.Get("iterations"))
		if err != nil {
			ctx.Text(err.Error())
			return
		}
		if iterations < 0 || iterations > HOPALONG_MAX_ITERATIONS {
			ctx.Text(fmt.Sprintf("Too many iterations. Max: %d\n", HOPALONG_MAX_ITERATIONS))
			return
		}
		fractal.Iterations = iterations
	}
	if query.Has("resolution") {
		resolution, err := strconv.Atoi(query.Get("resolution"))
		if err != nil {
			ctx.Text(err.Error())
			return
		}
		if resolution < 0 || resolution > HOPALONG_MAX_RESOLUTION {
			ctx.Text(fmt.Sprintf("Resolution is too high. Max: %d\n", HOPALONG_MAX_RESOLUTION))
			return
		}
		fractal.Resolution = resolution
	}
	if query.Has("a") {
		a, err := strconv.ParseFloat(query.Get("a"), 32)
		if err != nil {
			ctx.Text(err.Error())
			return
		}
		fractal.A = a
	}
	if query.Has("b") {
		b, err := strconv.ParseFloat(query.Get("b"), 32)
		if err != nil {
			ctx.Text(err.Error())
			return
		}
		fractal.B = b
	}
	if query.Has("c") {
		c, err := strconv.ParseFloat(query.Get("c"), 32)
		if err != nil {
			ctx.Text(err.Error())
			return
		}
		fractal.C = c
	}
	if query.Has("background") {
		background, err := helpers.ParseColor(query.Get("background"))
		if err != nil {
			ctx.Text(err.Error())
			return
		}
		fractal.Background = background
	}
	ctx.ContentType("image/png")
	fractal.WriteImage(ctx.ResponseWriter())
}
