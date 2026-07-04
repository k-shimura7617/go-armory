// Package spd provides the core logic for the "spd" CLI command.
package spd

import "github.com/k-shimura7617/go-armory/internal"

func Run() error {
	if err := internal.RunCmd("speedtest-cli"); err != nil {
		return err
	}

	return nil
}
