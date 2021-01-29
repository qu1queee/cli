package build

import (
	"github.com/shipwright-io/cli/pkg/cmd/flags"
	"github.com/spf13/cobra"
)

// CreateCmd represents the create command
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "creates a Build instance",
	Long: `
This generates a Build instance in the desired namespace.
`,

	SilenceUsage:  true,
	SilenceErrors: true,

	RunE: func(cmd *cobra.Command, args []string) error {
		return createFunc()
	},
}

func init() {
	BuildCmd.AddCommand()
	flags.CommonFlags(CreateCmd)
}

func createFunc() error {
	// obtain clients first
	// clients.NewClients()
	return nil
}
