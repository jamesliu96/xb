package main

import (
	"encoding/hex"
	"io"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		if _, err := io.Copy(os.Stdout, hex.NewDecoder(os.Stdin)); err != nil {
			panic(err)
		}
		return
	}
	if _, err := io.Copy(hex.NewEncoder(os.Stdout), os.Stdin); err != nil {
		panic(err)
	}
}
