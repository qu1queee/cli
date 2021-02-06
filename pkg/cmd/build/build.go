package build

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// BuildCmd represents the build command
var BuildCmd = &cobra.Command{
	Use:   "build",
	Short: "how to build container images",
	Long: `
Builds are used for describing how to build
your container images and where to store them.

Creating a Build instance doesn´t trigger the 
building mechanism.
`,

	SilenceUsage:  true,
	SilenceErrors: true,

	RunE: func(cmd *cobra.Command, args []string) error {
		return buildFunc()
	},
}

func init() {
	BuildCmd.AddCommand(
		CreateCmd,
		DeleteCmd,
	)
}

func buildFunc() error {
	// authenticate to cluster
	return nil
}

// MetaData ...
type MetaData struct {
	Name             string
	Namespace        string
	SourceURL        string
	RegistryEndpoint string
	Strategy         string
	SecretName       string
}

// NewMetaData ...
func (o *MetaData) NewMetaData() MetaData {
	return MetaData{
		Name:             viper.GetString("build-name"),
		Namespace:        viper.GetString("namespace"),
		SourceURL:        viper.GetString("source-url"),
		RegistryEndpoint: viper.GetString("registry-endpoint"),
		Strategy:         viper.GetString("strategy"),
		SecretName:       viper.GetString("registry-secret"),
	}
}
