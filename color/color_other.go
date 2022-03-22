//go:build windows || darwin

package color

const (
	Reset Code = 0
	Black Code = iota + 29
	Red
	Green
	Yellow
	Blue
	Purple
	Cyan
	Gray
	White
)
