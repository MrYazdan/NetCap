package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// newVersionCommand creates and returns the "version" command for NetCap.
// This command displays version-related information including the current
// version number and build timestamp.
//
// Returns:
//   - *cobra.Command: A configured Cobra command that displays version information
func newVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Show version information",
		Long: `Show the version information for NetCap.

Displays the current version, build time, and other 
version-related information.`,
		Run: func(cmd *cobra.Command, args []string) {
			PrintVersion()
		},
	}

	return cmd
}

// PrintVersion prints version information.
func PrintVersion() {
	// Print the application version to stdout
	fmt.Printf("NetCap version %s\n", Version)
	// Print the build timestamp to stdout
	fmt.Printf("Build time: %s\n", BuildTime)
}
