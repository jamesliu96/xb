package main

import (
	"encoding/hex"
	"io"
	"os"
)

func main() {
	io.Copy(os.Stdout, hex.NewDecoder(os.Stdin))
}
