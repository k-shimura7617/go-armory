package main

import (
	"fmt"
	"os"

	"github.com/k-shimura7617/go-armory/internal/lsp"
)

func main() {
	if err := lsp.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
