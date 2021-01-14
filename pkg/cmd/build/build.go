package build

import "github.com/spf13/cobra"

// BuildCmd represents the build command
var BuildCmd = &cobra.Command{
	Use:   "build",
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
