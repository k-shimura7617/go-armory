package main

import (
	"fmt"
	"os"

	"github.com/k-shimura7617/go-armory/internal/vid"
)

func main() {
	if err := vid.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
