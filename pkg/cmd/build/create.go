package build

import (
	"github.com/spf13/cobra"
)

// CreateCmd represents the create command
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "todo",
	Long: `

todo
	`,

	SilenceUsage:  true,
	SilenceErrors: true,

	RunE: func(cmd *cobra.Command, args []string) error {
		return createFunc()
	},
}

func init() {
	BuildCmd.AddCommand()
}

func createFunc() error {
	return nil
}
