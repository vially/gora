package gora

import "testing"

func TestStopEvents(t *testing.T) {
	var stopEventTests = []struct {
		in   stopEvent
		out  string
		desc string
	}{
		{stopEvent{clear: true}, "\r\033[K", "clear event"},
		{stopEvent{message: "foo", symbol: "bar"}, "\r\033[Kbar foo\n", "stop and persist event"},
	}

	for _, tt := range stopEventTests {
		t.Run(tt.desc, func(t *testing.T) {
			s := tt.in.String()
			if s != tt.out {
				t.Errorf("got %q, want %q", s, tt.out)
			}
		})
	}
}

func TestStatusLineRender(t *testing.T) {
	symbol := "A"
	text := "foo"
	expected := "\rA foo\x1b[K"
	out := fullStatusLine(symbol, text)

	if out != expected {
		t.Errorf("got %q, want %q", out, expected)
	}
}
