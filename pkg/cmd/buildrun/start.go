package buildrun

import (
	"github.com/spf13/cobra"
)

// StartCmd represents the start command
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "todo",
	Long: `

Todo
Todo
Todo
	`,

	SilenceUsage:  true,
	SilenceErrors: true,

	RunE: func(cmd *cobra.Command, args []string) error {
		return startFunc()
	},
}

func startFunc() error {
	return nil
}
