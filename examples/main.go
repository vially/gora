package main

import (
	"time"

	"github.com/logrusorgru/aurora"
	"github.com/vially/gora"
	"github.com/vially/gora/symbols"
)

func main() {
	spinner := gora.New("Loading stuff")
	spinner.Start()
	time.Sleep(3 * time.Second)
	spinner.UpdateColor(aurora.YellowFg)
	spinner.UpdateText("Loading more stuff")
	time.Sleep(3 * time.Second)
	spinner.StopAndPersist("Task 1 finished", symbols.ColoredSuccess)

	spinner2 := gora.New("Starting another task")
	spinner2.Start()
	time.Sleep(3 * time.Second)
	spinner2.StopAndPersist("Task 2 finished", symbols.ColoredSuccess)

	spinner3 := gora.New("Cleaning up")
	spinner3.Start()
	time.Sleep(3 * time.Second)
	spinner3.Stop()
}
