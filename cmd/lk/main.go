package main

import (
	"fmt"
	"os"

	"github.com/k-shimura7617/go-armory/internal/lk"
)

func main() {
	if err := lk.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
