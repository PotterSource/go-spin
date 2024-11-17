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

func NewColor(r, g, b int) (Color, error) {
	color, err := NewRGBColor(r, g, b)
	if err != nil {
		return nil, err
	}
	return color, nil
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

func Yellow() Color {
	yellow, _ := NewRGBColor(255, 255, 0)
	return yellow
}

func Cyan() Color {
	cyan, _ := NewRGBColor(0, 255, 255)
	return cyan
}

func Magenta() Color {
	magenta, _ := NewRGBColor(255, 0, 255)
	return magenta
}

func Orange() Color {
	orange, _ := NewRGBColor(255, 165, 0)
	return orange
}

func White() Color {
	white, _ := NewRGBColor(255, 255, 255)
	return white
}

func Black() Color {
	black, _ := NewRGBColor(0, 0, 0)
	return black
}
