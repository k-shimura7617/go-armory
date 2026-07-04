// Package lg provides the core logic for the "lg" CLI command.
package lg

import (
	"fmt"
	"os"

	"github.com/k-shimura7617/go-armory/internal"
)

func Run() {
	if err := internal.RunCmd("lazygit"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
