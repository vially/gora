package symbols

import (
	"github.com/logrusorgru/aurora"
)

var (
	// Info represents an information message
	Info = "ℹ"
	// Success represents a success message
	Success = "✔"
	// Warning represents a warning message
	Warning = "⚠"
	// Error represents an error message
	Error = "✖"

	// ColoredInfo represents a colored information message
	ColoredInfo = aurora.White(Info).String()
	// ColoredSuccess represents a colored success message
	ColoredSuccess = aurora.Green(Success).String()
	// ColoredWarning represents a colored warning message
	ColoredWarning = aurora.Yellow(Warning).String()
	// ColoredError represents a colored error message
	ColoredError = aurora.Red(Error).String()
)
