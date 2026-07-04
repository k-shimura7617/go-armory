package main

import (
	"fmt"
	"os"

	"github.com/k-shimura7617/go-armory/internal/pipfr"
)

func main() {
	if err := pipfr.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
