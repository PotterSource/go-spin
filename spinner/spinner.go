package spinner

import (
	"errors"
	"fmt"
	"github.com/PotterSource/go-spin/spinner/color"
	"time"
)

type Frames []string
type Type map[string]Frames

type Spinner struct {
	frames Frames
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
	rgbColor color.Color,
) (*Spinner, error) {
	frames, ok := spinnerTypes[spinnerType]
	if !ok {
		return nil, errors.New("invalid spinner type")
	}

	return &Spinner{
		frames: frames,
		stop:   make(chan bool),
		color:  rgbColor,
	}, nil
}

func (s *Spinner) Start() {
	go func() {
		for {
			for _, frame := range s.frames {
				select {
				case <-s.stop:
					return
				default:
					fmt.Printf("\r%s%s\033[0m", s.color.ANSI(), frame)
					time.Sleep(100 * time.Millisecond)
				}
			}
		}
	}()
}

func (s *Spinner) Stop(messages ...string) {
	s.stop <- true
	close(s.stop)
	if len(messages) > 0 {
		fmt.Println("\r" + messages[0])
	}
}
