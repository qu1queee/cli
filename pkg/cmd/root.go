package cmd

import (
	"github.com/shipwright-io/cli/pkg/cmd/build"
	"github.com/shipwright-io/cli/pkg/cmd/buildrun"
	"github.com/shipwright-io/cli/pkg/cmd/flags"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when used without any subcommands
var rootCmd = &cobra.Command{
	Use:   "shp",
	Short: "Command-line client for Shipwright Build.",
	Long: `Command-line client for Shipwright Build

shp allows you to build container images securely inside your 
Kubernetes cluster.

See the individual commands to get an overview.
`,
}

func init() {
	flags.CommonFlags(rootCmd)
	rootCmd.AddCommand(
		build.BuildCmd,
		buildrun.BuildRunCmd,
	)
}

// NewSHPRootCmd returns the cobra base cmd
func NewSHPRootCmd() *cobra.Command {
	return rootCmd
}
