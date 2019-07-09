package gora

import (
	"fmt"
	"strings"
	"testing"

	"github.com/logrusorgru/aurora"
)

var basicFrameSet = []string{"A", "B", "C"}

type mockColorizer struct {
	colored bool
}

func (m *mockColorizer) Colorize(text string, color aurora.Color) string {
	if !m.colored {
		return text
	}
	return fmt.Sprintf("%d: %s", color, text)
}

func TestAnimationFrames(t *testing.T) {
	var animationTests = []struct {
		animation *animation
		frames    int
		out       string
		desc      string
	}{
		{newAnimation(basicFrameSet), 4, "ABCA", "basic animation"},
		{newAnimation([]string{}), 4, "", "no output on empty frameset"},
	}

	for _, tt := range animationTests {
		t.Run(tt.desc, func(t *testing.T) {
			s := runAnimation(tt.animation, tt.frames)
			if s != tt.out {
				t.Errorf("got %q, want %q", s, tt.out)
			}
		})
	}
}

func TestAnimationDefaults(t *testing.T) {
	defaultAnimation := newDefaultAnimation()
	if defaultAnimation.color != aurora.CyanFg {
		t.Errorf("wrong color for default animation; got %q, want %q", defaultAnimation.color, aurora.CyanFg)
	}
}

func TestAuroraColorizer(t *testing.T) {
	c := &auroraColorizer{}
	out := c.Colorize("foo", aurora.CyanFg)
	expected := "\x1b[36mfoo\x1b[0m"
	if out != expected {
		t.Errorf("got %q, want %q", out, expected)
	}
}

func newAnimation(frameSet []string) *animation {
	return &animation{
		colorizer: &mockColorizer{},
		frameSet:  frameSet,
	}
}

func runAnimation(a *animation, frames int) string {
	var result strings.Builder

	for i := 0; i < frames; i++ {
		result.WriteString(a.NextFrame())
	}

	return result.String()
}
