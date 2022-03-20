# gotest

Bring some color to `go test`.



Inspired by <https://github.com/rakyll/gotest> and <https://github.com/kyoh86/richgo>.
Please check these out for a more complete solution.

## Installation

For Go 1.16+:

```
$ go install github.com/joaonsantos/gotestcolor@latest
```

# Usage

Use `gotestcolor` just like you would `go test`.

You can also replace `go test` with `gotestcolor` in your local makefiles.

Example:

```
$ gotestcolor -v -count=1 ./...
```
![gotestcolor output example](https://raw.githubusercontent.com/joaonsantos/gotestcolor/assets/gotestcolor-demo.png)
