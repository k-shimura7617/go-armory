// Package dkdown provides the core logic for the "dkdown" CLI command.
package dkdown

import (
	"fmt"
	"os"

	"github.com/k-shimura7617/go-armory/internal"
)

func Run() {
	if err := internal.RunCmd("docker", "compose", "down"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
