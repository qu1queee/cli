package flags

import (
	"fmt"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// CommonFlags provides a set of persistent flags for root and subcmds
func CommonFlags(cmd *cobra.Command) {
	home, err := homedir.Dir()
	if err != nil {
		panic(fmt.Errorf("unable to get home directory: %v", err))
	}
	cmd.PersistentFlags().StringP("namespace", "n", "default", "desired namespace")
	cmd.PersistentFlags().StringP("kubeconfig", "k", filepath.Join(home, ".kube", "config"), "Kubernetes configuration file")

	viper.BindPFlag("namespace", cmd.PersistentFlags().Lookup("namespace"))
	viper.BindPFlag("kubeconfig", cmd.PersistentFlags().Lookup("kubeconfig"))
}

// BuildFlags provides a set of flags for the Build subcommands
func BuildFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("build-name", "b", "", "build name")
	cmd.Flags().StringP("source-url", "u", "", "url to the source code")
	cmd.Flags().StringP("registry-endpoint", "e", "", "registry endpoint")
	cmd.Flags().StringP("registry-secret", "r", "", "registry secret for pushing")
	cmd.Flags().StringP("strategy", "s", "", "strategy of choice")

	viper.BindPFlag("build-name", cmd.Flags().Lookup("build-name"))
	viper.BindPFlag("source-url", cmd.Flags().Lookup("source-url"))
	viper.BindPFlag("registry-endpoint", cmd.Flags().Lookup("registry-endpoint"))
	viper.BindPFlag("strategy", cmd.Flags().Lookup("strategy"))
	viper.BindPFlag("registry-secret", cmd.Flags().Lookup("registry-secret"))

}
