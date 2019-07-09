# gora

> Elegant terminal spinner for Go inspired by the excellent [ora](https://github.com/sindresorhus/ora)

# Usage

```go
package main

import (
    "time"

    "github.com/logrusorgru/aurora"
    "github.com/vially/gora"
    "github.com/vially/gora/symbols"
)

func main() {
    spinner := gora.New("Loading unicorns")
    spinner.Start()
    time.Sleep(3 * time.Second)

    spinner.UpdateColor(aurora.YellowFg)
    spinner.UpdateText("Loading rainbows")
    time.Sleep(3 * time.Second)

    spinner.StopAndPersist("Done", symbols.ColoredSuccess)
}
```
