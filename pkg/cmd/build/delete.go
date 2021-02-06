package build

import (
	"context"
	"time"

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

		// Ensure all subCommands get a ctx timeout
		// TODO: move this block out of here
		ctx := context.Background()
		subCommandsCTX, cancelFunc := context.WithTimeout(ctx, 5*time.Second)
		defer cancelFunc()

		// Get the clients
		c, err := clients.NewClients(subCommandsCTX)
		if err != nil {
			return err
		}

		return deleteFunc(c)
	},
}

func init() {
	BuildCmd.AddCommand()
	flags.CommonFlags(DeleteCmd)
}

func deleteFunc(c *clients.Clients) error {
	return nil
}
