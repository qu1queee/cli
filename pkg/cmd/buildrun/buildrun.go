package buildrun

import (
	"github.com/spf13/cobra"
)

// BuildRunCmd represents the buildrun command
var BuildRunCmd = &cobra.Command{
	Use:   "buildrun",
	Short: "starts a BuildRun",
	Long: `

Todo
Todo
Todo
	`,

	SilenceUsage:  true,
	SilenceErrors: true,

	RunE: func(cmd *cobra.Command, args []string) error {
		return buildRunFunc()
	},
}

func init() {
	BuildRunCmd.AddCommand(
		StartCmd,
	)
}

func buildRunFunc() error {
	return nil
}
