//go:build linux || openbsd

package color

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
