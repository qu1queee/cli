package buildrun

import (
	"github.com/spf13/cobra"
)

// BuildRunCmd represents the buildrun command
var BuildRunCmd = &cobra.Command{
	Use:   "buildrun",
	Short: "todo",
	Long: `

Todo
Todo
Todo
	`,

	SilenceUsage:  true,
	SilenceErrors: true,

	RunE: func(cmd *cobra.Command, args []string) error {
		// todo: trigger basic validations
		// around the cmd usage
		return CmdFunc()
	},
}

// CmdFunc ...
func CmdFunc() error {
	// todo: add cmd business logic
	return nil
}
