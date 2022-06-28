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
	JULIA_SET_MAX_ITERATIONS        = 500_000
	JULIA_SET_DEFAULT_ITERATIONS    = 250
	JULIA_SET_DEFAULT_COLOR_PALETTE = "orange_blue"
	JULIA_SET_DEFAULT_C             = -0.5 + 0.6i
	JULIA_SET_DEFAULT_BAIL_OUT      = 2
	JULIA_SET_DEFAULT_REGION        = "-1.5, -1.5, 3, 3"
)

func GetJuliaSet(ctx iris.Context) {
	query := ctx.Request().URL.Query()
	fractal := fractals.JuliaSet{
		Width:         DEFAULT_WIDTH,
		Height:        DEFAULT_HEIGHT,
		C:             JULIA_SET_DEFAULT_C,
		MaxIterations: JULIA_SET_DEFAULT_ITERATIONS,
		BailOut:       JULIA_SET_DEFAULT_BAIL_OUT,
		Background:    color.RGBA{255, 255, 255, 255},
	}
	colorPaletteValue := JULIA_SET_DEFAULT_COLOR_PALETTE
	regionValue := JULIA_SET_DEFAULT_REGION
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
	if query.Has("c") {
		c, err := strconv.ParseComplex(query.Get("c"), 64)
		if err != nil {
			ctx.Text(err.Error())
			return
		}
		fractal.C = c
	}
	if query.Has("color_palette") {
		colorPaletteValue = query.Get("color_palette")
	}
	if query.Has("iterations") {
		iterations, err := strconv.Atoi(query.Get("iterations"))
		if err != nil {
			ctx.Text(err.Error())
			return
		}
		if iterations < 0 || iterations > JULIA_SET_MAX_ITERATIONS {
			ctx.Text(fmt.Sprintf("Too many iterations. Max: %d\n", JULIA_SET_MAX_ITERATIONS))
			return
		}
		fractal.MaxIterations = iterations
	}
	if query.Has("region") {
		regionValue = query.Get("region")
	}
	if query.Has("bail_out") {
		bailOut, err := strconv.ParseFloat(query.Get("bail_out"), 32)
		if err != nil {
			ctx.Text(err.Error())
			return
		}
		fractal.BailOut = bailOut
	}
	if query.Has("background") {
		background, err := helpers.ParseColor(query.Get("background"))
		if err != nil {
			ctx.Text(err.Error())
			return
		}
		fractal.Background = background
	}
	region, err := helpers.ParseRect(regionValue)
	if err != nil {
		ctx.Text(err.Error())
		return
	}
	colorPalette, err := helpers.ParseColorPalette(colorPaletteValue)
	if err != nil {
		ctx.Text(err.Error())
		return
	}
	fractal.Region = region
	fractal.ColorPalette = colorPalette
	ctx.ContentType("image/png")
	err = fractal.WriteImage(ctx.ResponseWriter())
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}
