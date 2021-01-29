package build

import (
	"github.com/spf13/cobra"
)

// BuildCmd represents the build command
var BuildCmd = &cobra.Command{
	Use:   "build",
	Short: "how to build container images",
	Long: `
Builds are used for describing how to build
your container images and where to store them.

Creating a Build instance doesn´t trigger the 
building mechanism.
`,

	SilenceUsage:  true,
	SilenceErrors: true,

	RunE: func(cmd *cobra.Command, args []string) error {
		return buildFunc()
	},
}

func init() {
	BuildCmd.AddCommand(
		CreateCmd,
		DeleteCmd,
	)
}

func buildFunc() error {
	// authenticate to cluster
	return nil
}
