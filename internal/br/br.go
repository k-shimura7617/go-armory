// Package br provides the core logic for the "br" CLI command.
package br

import (
	"github.com/k-shimura7617/go-armory/internal"
)

func Run() error {
	if err := internal.RunCmd("git", "branch", "-a"); err != nil {
		return err
	}

	return nil
}
