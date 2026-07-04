package main

import (
	"fmt"
	"os"

	"github.com/k-shimura7617/go-armory/internal/cl"
)

func main() {
	if err := cl.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
