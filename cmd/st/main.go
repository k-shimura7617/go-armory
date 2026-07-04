package main

import (
	"fmt"
	"os"

	"github.com/k-shimura7617/go-armory/internal/st"
)

func main() {
	if err := st.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
