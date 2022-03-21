package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/joaonsantos/gotestcolor/color"
)

func main() {
	args := []string{"test"}
	args = append(args, os.Args[1:]...)

	cmd := exec.Command("go", args...)

	r, w := io.Pipe()

	cmd.Stderr = w
	cmd.Stdout = w
	cmd.Env = os.Environ()

	go parse(r)

	cmd.Run()
}

func parse(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		switch {
		case strings.HasPrefix(line, "--- FAIL") || strings.HasPrefix(line, "FAIL"):
			color.Colorize(color.Red, "%v\n", line)
		case strings.HasPrefix(line, "--- SKIP") || strings.HasPrefix(line, "SKIP"):
			color.Colorize(color.Yellow, "%v\n", line)
		case strings.HasPrefix(line, "PASS") || strings.HasPrefix(line, "ok "):
			color.Colorize(color.Green, "%v\n", line)
		default:
			fmt.Println(line)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
	}
}
