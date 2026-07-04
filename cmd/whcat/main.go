package main

import (
	"fmt"
	"os"

	"github.com/k-shimura7617/go-armory/internal/whcat"
)

func main() {
	if err := whcat.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
