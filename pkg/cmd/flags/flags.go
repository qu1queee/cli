package flags

import (
	"fmt"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// CommonFlags provides a set of flags widely use across commands
func CommonFlags(cmd *cobra.Command) {
	home, err := homedir.Dir()
	if err != nil {
		panic(fmt.Errorf("unable to get home directory: %v", err))
	}

	cmd.PersistentFlags().String("kubeconfig", filepath.Join(home, ".kube", "config"), "Kubernetes configuration file")
	cmd.PersistentFlags().String("namespace", "default", "desired namespace")
}
