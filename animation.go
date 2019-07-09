package gora

import "github.com/logrusorgru/aurora"

var defaultFrameSet = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
var defaultColorizer = &auroraColorizer{}
var defaultColor = aurora.CyanFg

type colorizer interface {
	Colorize(string, aurora.Color) string
}

type auroraColorizer struct{}

func (*auroraColorizer) Colorize(text string, color aurora.Color) string {
	return aurora.White(text).Colorize(color).String()
}

type animation struct {
	currentFrame int
	colorizer    colorizer
	frameSet     []string
	color        aurora.Color
}

func (a *animation) NextFrame() string {
	if len(a.frameSet) == 0 {
		return ""
	}

	frame := a.frameSet[a.currentFrame%len(a.frameSet)]
	a.currentFrame++

	return a.colorizer.Colorize(frame, a.color)
}

func newDefaultAnimation() *animation {
	return &animation{
		colorizer: defaultColorizer,
		frameSet:  defaultFrameSet,
		color:     defaultColor,
	}
}
