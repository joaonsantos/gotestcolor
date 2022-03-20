// Package color provides definitions to colorize terminal output.
package color

import "fmt"

type Code int

const (
	Reset Code = 0
	Black Code = iota + 30
	Red
	Green
	Yellow
	Blue
	Purple
	Cyan
	Gray
	White
)

func Colorize(color Code, format string, a ...any) {
	stringFmt := fmt.Sprintf(format, a...)
	fmt.Printf("\033[%dm%s\033[%dm", color, stringFmt, Reset)
}
