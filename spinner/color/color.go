package color

import (
	"errors"
	"fmt"
)

type Color interface {
	ANSI() string
}

type RGBColor struct {
	R, G, B int
}

func NewRGBColor(r, g, b int) (*RGBColor, error) {
	if (r < 0 || r > 255) || (g < 0 || g > 255) || (b < 0 || b > 255) {
		return nil, errors.New("RGB values must be between 0 and 255")
	}

	return &RGBColor{r, g, b}, nil
}

func (rgb *RGBColor) ANSI() string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", rgb.R, rgb.G, rgb.B)
}

func Red() Color {
	red, _ := NewRGBColor(255, 0, 0)
	return red
}

func Green() Color {
	green, _ := NewRGBColor(0, 255, 0)
	return green
}

func Blue() Color {
	blue, _ := NewRGBColor(0, 0, 255)
	return blue
}