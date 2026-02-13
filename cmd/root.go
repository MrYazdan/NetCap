package cmd

import (
	"github.com/spf13/cobra"
)

// Version is the current version of NetCap.
// It can be set during build using ldflags.
var Version = "dev"

// BuildTime is the build timestamp.
// It can be set during build using ldflags.
var BuildTime = "unknown"

// GlobalFlags holds global flags that apply to all commands.
type GlobalFlags struct {
	// Config is the path to the config file
	Config string

	// Output is the output file or directory
	Output string

	// Format is the output format
	Format string

	// Verbose enables verbose logging
	Verbose bool

	// Quiet suppresses all output except errors
	Quiet bool
}

// globalFlags holds the parsed global flags.
var globalFlags = &GlobalFlags{}

// Execute runs the root command.
// It returns an error if command execution fails.
func Execute() error {
	// Create root command
	rootCmd := &cobra.Command{
		Use:   "netcap",
		Short: "NetCap - Powerful network interception and logging tool",
		Long: `NetCap is a network interception and logging tool that wraps 
command execution and captures all HTTP/HTTPS traffic passing 
through the system.

Examples:
  # Run a command with network interception
  netcap python main.py

  # Run as a proxy server
  netcap proxy --port 8080

  # List recorded sessions
  netcap list

  # Replay a recorded request
  netcap replay <session-id>`,
		Version: Version,
	}

	// Add global flags
	rootCmd.PersistentFlags().StringVarP(&globalFlags.Config, "config", "c", "", "Config file path (default: ~/.netcap.yaml)")
	rootCmd.PersistentFlags().StringVarP(&globalFlags.Output, "output", "o", "", "Output file or directory (default: ./netcap-output)")
	rootCmd.PersistentFlags().StringVarP(&globalFlags.Format, "format", "f", "pretty", "Output format: json, csv, pretty, har, raw")
	rootCmd.PersistentFlags().BoolVarP(&globalFlags.Verbose, "verbose", "v", false, "Enable verbose logging")
	rootCmd.PersistentFlags().BoolVarP(&globalFlags.Quiet, "quiet", "q", false, "Suppress all output except errors")

	// Add commands
	rootCmd.AddCommand(
		// TODO: implement sub-commands :
		//   - run
		//   - proxy
		//   - list
		//   - replay
		newVersionCommand(),
	)

	// Execute
	return rootCmd.Execute()
}

// GetGlobalFlags returns the global flags.
func GetGlobalFlags() *GlobalFlags {
	return globalFlags
}
