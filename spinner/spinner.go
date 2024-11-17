package spinner

import (
	"fmt"
	"github.com/PotterSource/go-spin/spinner/color"
	"time"
)

type Frames []string
type Type map[string]Frames

type Spinner struct {
	frames Frames
	delay  time.Duration
	stop   chan bool
	color  color.Color
}

var spinnerTypes = Type{
	"bar": {
		"▁", "▂", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃", "▂",
	},
	"dots": {
		"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏",
	},
	"dots2": {
		"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷",
	},
	"line": {
		"-", "\\", "|", "/",
	},
}

func NewSpinner(
	spinnerType string,
	delay time.Duration,
	rgbColor color.Color,
) *Spinner {
	frames, ok := spinnerTypes[spinnerType]
	if !ok {
		return nil
	}
	return &Spinner{
		frames: frames,
		delay:  delay,
		stop:   make(chan bool),
		color:  rgbColor,
	}
}

func (s *Spinner) Start() {
	go func() {
		for {
			for _, frame := range s.frames {
				select {
				case <-s.stop:
					return
				default:
					// Ensuring the color is converted to a string properly
					fmt.Printf("\r%s%s\033[0m", s.color.ANSI(), frame)
					time.Sleep(s.delay)
				}
			}
		}
	}()
}

func (s *Spinner) Stop() {
	s.stop <- true
	close(s.stop)
	fmt.Println("\rSpinner stopped.")
}
