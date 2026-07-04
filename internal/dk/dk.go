// Package dk provides the core logic for the "dk" CLI command.
package dk

import (
	"fmt"
	"os"

	"github.com/k-shimura7617/go-armory/internal"
)

func Run() {
	if err := internal.RunCmd("docker", "compose", "ps", "-a"); err != nil {
		if err := internal.RunCmd("docker", "ps"); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
