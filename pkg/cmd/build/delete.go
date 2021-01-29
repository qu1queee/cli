package build

import (
	"github.com/shipwright-io/cli/pkg/cmd/clients"
	"github.com/shipwright-io/cli/pkg/cmd/flags"
	"github.com/spf13/cobra"
)

// DeleteCmd represents the delete command
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "deletes a Build instance",
	Long: `
This removes an existent Build instance from the desired namespace.
`,

	SilenceUsage:  true,
	SilenceErrors: true,

	RunE: func(cmd *cobra.Command, args []string) error {
		return deleteFunc()
	},
}

func init() {
	BuildCmd.AddCommand()
	flags.CommonFlags(DeleteCmd)
}

func deleteFunc() error {
	// obtain clients first
	clients.NewClients()
	return nil
}
