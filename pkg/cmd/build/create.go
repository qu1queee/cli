package build

import (
	"context"
	"time"

	"github.com/shipwright-io/cli/pkg/cmd/clients"
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

		return createFunc(c)
	},
}

func init() {
	BuildCmd.AddCommand()
	flags.CommonFlags(CreateCmd)
}

func createFunc(c *clients.Clients) error {
	return nil
}
