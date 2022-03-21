// Package color provides definitions to colorize terminal output.
package color

import "fmt"

type Code int

func Colorize(color Code, format string, a ...any) {
	stringFmt := fmt.Sprintf(format, a...)
	fmt.Printf("\033[%dm%s\033[%dm", color, stringFmt, Reset)
}
