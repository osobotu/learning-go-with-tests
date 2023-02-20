package main

import (
	"clockface"
	"os"
	"time"
	// "github.com/osobotu/learning-go-with-tests/math/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
