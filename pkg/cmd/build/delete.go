package build

import (
	"github.com/spf13/cobra"
)

// DeleteCmd represents the delete command
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete´s a Build",
	Long: `

todo
	`,

	SilenceUsage:  true,
	SilenceErrors: true,

	RunE: func(cmd *cobra.Command, args []string) error {
		return deleteFunc()
	},
}

func init() {
	BuildCmd.AddCommand()
}

func deleteFunc() error {
	return nil
}
