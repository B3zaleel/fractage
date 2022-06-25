package helpers

import (
	"errors"
	"image/color"
	"math"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	NIL_COLOR = color.RGBA{0, 0, 0, 0}
)

// Represents a color palette.
type ColorPalette struct {
	Name        string       `yaml:"name"`
	Transitions []Transition `yaml:"transitions"`
}

// Represents a color transition.
type Transition struct {
	Color    string `yaml:"color"`
	_Color   *color.RGBA
	Position float32 `yaml:"position"`
}

// Gets the color value of a given position in this ColorPalette.
func (palette *ColorPalette) GetColor(pos float64) (color.RGBA, error) {
	if len(palette.Transitions) == 0 {
		return NIL_COLOR, errors.New("ColorPalette has no color transitions")
	}
	value := math.Max(0, math.Min(1.0, pos))
	transitionIdx := 0
	for i := 0; i < len(palette.Transitions); i++ {
		if value >= float64(palette.Transitions[i].Position) {
			break
		}
		transitionIdx++
	}
	if transitionIdx >= len(palette.Transitions)-1 {
		if palette.Transitions[transitionIdx]._Color != nil {
			return *palette.Transitions[transitionIdx]._Color, nil
		}
		return ParseColor(palette.Transitions[transitionIdx].Color)
	}
	curTransition := palette.Transitions[transitionIdx]
	nextTransition := palette.Transitions[transitionIdx+1]
	var err error
	var curColor color.RGBA
	if curTransition._Color != nil {
		curColor = *curTransition._Color
	} else {
		curColor, err = ParseColor(curTransition.Color)
		if err != nil {
			return NIL_COLOR, err
		}
	}
	var nextColor color.RGBA
	if nextTransition._Color != nil {
		nextColor = *nextTransition._Color
	} else {
		nextColor, err = ParseColor(nextTransition.Color)
		if err != nil {
			return NIL_COLOR, err
		}
	}
	grad := (value - float64(curTransition.Position))
	grad /= (float64(nextTransition.Position) - float64(curTransition.Position))
	posColor := color.RGBA{
		R: curColor.R + uint8(grad)*(nextColor.R-curColor.R),
		G: curColor.G + uint8(grad)*(nextColor.G-curColor.G),
		B: curColor.B + uint8(grad)*(nextColor.B-curColor.B),
		A: 255,
	}
	return posColor, nil
}

// Translate the value of the color transitions for this color palette.
func (palette *ColorPalette) TranslateColorTransitions() error {
	for i := 0; i < len(palette.Transitions); i++ {
		transitionColor, err := ParseColor(palette.Transitions[i].Color)
		if err != nil {
			return err
		}
		palette.Transitions[i]._Color = &transitionColor
	}
	return nil
}

// Returns the color value of a predetermined color palette that
// matches the given name.
func ParseNameColorPalette(name string) (ColorPalette, error) {
	var palettes []ColorPalette
	nilPalette := ColorPalette{
		Name:        "",
		Transitions: nil,
	}
	file, err := os.ReadFile("src/data/color_palettes.yaml")
	if err != nil {
		return nilPalette, err
	}
	err = yaml.Unmarshal(file, &palettes)
	if err != nil {
		return nilPalette, err
	}
	for _, palette := range palettes {
		if palette.Name == name {
			return palette, nil
		}
	}
	return nilPalette, errors.New("Palette not found")
}
