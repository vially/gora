package gora

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/logrusorgru/aurora"
)

// Spinner is an elegant terminal spinner
type Spinner struct {
	writer    io.Writer
	text      string
	animation *animation
	stop      chan bool
}

// New creates a new spinner
func New(text string) *Spinner {
	return &Spinner{
		text:      text,
		animation: newDefaultAnimation(),
		writer:    os.Stdout,
		stop:      make(chan bool),
	}
}

// Start the spinner
func (s *Spinner) Start() {
	go func() {
		fmt.Fprint(s.writer, fullStatusLine(s.animation.NextFrame()+hideCursor, s.text))
		for {
			select {
			case <-s.stop:
				return
			case <-time.After(100 * time.Millisecond):
				fmt.Fprint(s.writer, fullStatusLine(s.animation.NextFrame(), s.text))
			}
		}
	}()
}

// UpdateText replaces the spinner text
func (s *Spinner) UpdateText(text string) {
	s.text = text
}

// UpdateColor replaces the spinner color
func (s *Spinner) UpdateColor(color aurora.Color) {
	s.animation.color = color
}

// StopAndPersist stops the spinner and changes the text and symbol
func (s *Spinner) StopAndPersist(text, symbol string) {
	s.stopWithEvent(&stopEvent{message: text, symbol: symbol})
}

// Stop the spinner and clear the text
func (s *Spinner) Stop() {
	s.stopWithEvent(&stopEvent{clear: true})
}

func (s *Spinner) stopWithEvent(event *stopEvent) {
	s.stop <- true
	fmt.Fprint(s.writer, event.String()+showCursor)
}

func fullStatusLine(animationFrame, text string) string {
	return fmt.Sprintf("\r%s %s%s", animationFrame, text, eraseToEndOfLine)
}

type stopEvent struct {
	message string
	symbol  string
	clear   bool
}

func (e *stopEvent) String() string {
	if e.clear {
		return eraseCurrentLine
	}
	return fmt.Sprintf("%s%s %s\n", eraseCurrentLine, e.symbol, e.message)
}

var eraseToEndOfLine = "\033[K"
var eraseCurrentLine = "\r" + eraseToEndOfLine
var hideCursor = "\033[?25l"
var showCursor = "\033[?25h"
