// Package gga provides the core logic for the "gga" CLI command.
package gga

import "github.com/k-shimura7617/go-armory/internal"

func Run() error {
	if err := internal.RunCmd("git", "tree"); err != nil {
		return err
	}

	return nil
}
