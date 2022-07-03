package fractals

import (
	"errors"
	"image"
	"image/color"
	"image/png"
	"io"
	"math"
	"math/cmplx"
	"strconv"
	"strings"

	"github.com/B3zaleel/fractage/src/helpers"
)

// Properties of a Julia set image.
type JuliaSet struct {
	Width              int
	Height             int
	ColorPalette       helpers.ColorPalette
	MaxIterations      int
	C                  complex128
	Variables          map[rune]complex128
	BailOut            float64
	Region             helpers.Rect
	SeriesFunctionName string
	Background         color.RGBA
}

// Creates a function that computes the sum of c and the absolute value of
// the 4th power of a trigonometric function for a given z.
func absTrig(props *JuliaSet, trigFxn func(complex128) complex128) func(complex128) complex128 {
	abs := func(c complex128) complex128 {
		return complex(math.Abs(real(c)), math.Abs(imag(c)))
	}
	return func(z complex128) complex128 {
		return abs(cmplx.Pow(trigFxn(z), 4)) + props.C
	}
}

// Creates a function that computes the product of c and the trigonometric value for a given z.
func cTrig(props *JuliaSet, trigFxn func(complex128) complex128) func(complex128) complex128 {
	return func(z complex128) complex128 {
		return props.C * trigFxn(z)
	}
}

var (
	JULIA_SET_SERIES = map[string]func(*JuliaSet) func(complex128) complex128{
		"classic": func(props *JuliaSet) func(complex128) complex128 {
			return func(z complex128) complex128 { return z*z + props.C }
		},
		"csin":       func(props *JuliaSet) func(complex128) complex128 { return cTrig(props, cmplx.Sin) },
		"ccos":       func(props *JuliaSet) func(complex128) complex128 { return cTrig(props, cmplx.Cos) },
		"ctan":       func(props *JuliaSet) func(complex128) complex128 { return cTrig(props, cmplx.Tan) },
		"abs_sin4":   func(props *JuliaSet) func(complex128) complex128 { return absTrig(props, cmplx.Sin) },
		"abs_cos4":   func(props *JuliaSet) func(complex128) complex128 { return absTrig(props, cmplx.Cos) },
		"abs_tan4":   func(props *JuliaSet) func(complex128) complex128 { return absTrig(props, cmplx.Tan) },
		"abs_cot4":   func(props *JuliaSet) func(complex128) complex128 { return absTrig(props, cmplx.Cot) },
		"abs_sinh4":  func(props *JuliaSet) func(complex128) complex128 { return absTrig(props, cmplx.Sinh) },
		"abs_cosh4":  func(props *JuliaSet) func(complex128) complex128 { return absTrig(props, cmplx.Cosh) },
		"abs_tanh4":  func(props *JuliaSet) func(complex128) complex128 { return absTrig(props, cmplx.Tanh) },
		"abs_asinh4": func(props *JuliaSet) func(complex128) complex128 { return absTrig(props, cmplx.Asinh) },
		"abs_acosh4": func(props *JuliaSet) func(complex128) complex128 { return absTrig(props, cmplx.Acosh) },
		"abs_atanh4": func(props *JuliaSet) func(complex128) complex128 { return absTrig(props, cmplx.Atanh) },
	}
)

// Writes the Julia set image to the given output.
func (props *JuliaSet) WriteImage(output io.Writer) error {
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

// Helper function for rendering the Julia set.
func (props *JuliaSet) render(img *image.RGBA) error {
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
	seriesFunction := JULIA_SET_SERIES[props.SeriesFunctionName](props)
	for y := 0; y < int(height); y++ {
		for x := 0; x < int(width); x++ {
			n = 0
			Z := complex(xOffset+float64(x)*step, yOffset+float64(y)*step)
			seriesValue := math.Exp(-cmplx.Abs(Z))
			for (n < props.MaxIterations) && (cmplx.Abs(Z) < props.BailOut) {
				Z = seriesFunction(Z)
				seriesValue += math.Exp(-cmplx.Abs(Z))
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

// Checks if a function name exists in the set of JULIA_SET_SERIES names.
func IsValidJuliaSetSeriesFunction(txt string) bool {
	fxnName := strings.Trim(txt, helpers.WHITESPACE_CUTSET)
	for name, _ := range JULIA_SET_SERIES {
		if name == fxnName {
			return true
		}
	}
	return false
}

// Retrieves the value of a variable and returns a default value if the variable doesn't exist.
func (props *JuliaSet) GetVaraible(c rune, defaultValue complex128) complex128 {
	for key, value := range props.Variables {
		if key == c {
			return value
		}
	}
	return defaultValue
}

// Converts a comma-separated list of variable assignments to a map of runes and complex numbers.
func ParseJuliaSetVariables(txt string) (map[rune]complex128, error) {
	values, err := helpers.GetCSV(txt)
	if err != nil {
		return nil, err
	}
	variables := make(map[rune]complex128, len(values))
	for i := 0; i < len(values); i++ {
		rule := strings.Trim(values[i], helpers.WHITESPACE_CUTSET)
		before, after, found := strings.Cut(rule, "=")
		if found {
			variable := []rune(strings.Trim(before, helpers.WHITESPACE_CUTSET))
			if len(variable) == 1 {
				valueText := strings.Trim(after, helpers.WHITESPACE_CUTSET)
				value, err := strconv.ParseComplex(valueText, 128)
				if err != nil {
					return nil, err
				}
				variables[variable[0]] = value
			} else {
				return nil, errors.New("A variable must be a single character")
			}
		} else {
			return nil, errors.New("Invalid variable assignment. It must be of the form variable=value")
		}
	}
	return variables, nil
}
