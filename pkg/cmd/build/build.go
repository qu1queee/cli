package build

import (
	"github.com/spf13/cobra"
)

// BuildCmd represents the build command
var BuildCmd = &cobra.Command{
	Use:   "build",
	Short: "creates a Build",
	Long: `

TODO
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
	return nil
}
