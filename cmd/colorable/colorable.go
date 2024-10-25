package main

import (
	"io"
	"os"

	"github.com/ilius/go-colorable"
)

func main() {
	io.Copy(colorable.NewColorableStdout(), os.Stdin)
}
