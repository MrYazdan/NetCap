// Package main is the entry point for the netcap CLI application.
//
// It bootstraps the command layer and delegates execution to the root command defined in the cmd package.
// Any execution error is treated as a fatal condition and results in a non-zero exit status.
package main

import (
	"netcap/cmd"
	"os"
)

func main() {
	// Execute initializes and runs the root command.
	// If an error occurs, exit with a non-zero status code
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
