package build

import (
	"github.com/shipwright-io/cli/pkg/cmd/options"
	"github.com/spf13/cobra"
	corev1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

		var options options.List
		opts, err := options.NewList()
		if err != nil {
			return err
		}
		return deleteFunc(opts)
	},
}

func init() {
	BuildCmd.AddCommand()
	// flags.CommonFlags(DeleteCmd)
}

func deleteFunc(opts options.List) error {
	data := Info.NewMetaData()
	err := opts.Clients.BuildClient.BuildV1alpha1().Builds(data.Namespace).Delete(opts.Context, data.Name, corev1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}
