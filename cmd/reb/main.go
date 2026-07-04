package main

import (
	"fmt"
	"os"

	"github.com/k-shimura7617/go-armory/internal/reb"
)

func main() {
	if err := reb.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
