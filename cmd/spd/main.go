package main

import (
	"fmt"
	"os"

	"github.com/k-shimura7617/go-armory/internal/spd"
)

func main() {
	if err := spd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
