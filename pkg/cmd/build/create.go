package build

import (
	buildv1alpha1 "github.com/shipwright-io/build/pkg/apis/build/v1alpha1"
	"github.com/shipwright-io/cli/pkg/cmd/flags"
	"github.com/shipwright-io/cli/pkg/cmd/options"
	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Info ...
var Info MetaData

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

		var options options.List
		opts, err := options.NewList()
		if err != nil {
			return err
		}
		b, err := constructObject()
		if err != nil {
			return err
		}
		return createFunc(b, opts)
	},
}

func init() {
	BuildCmd.AddCommand()
	flags.BuildFlags(CreateCmd)
}

func createFunc(build *buildv1alpha1.Build, opts options.List) error {
	_, err := opts.Clients.BuildClient.BuildV1alpha1().Builds(build.Namespace).Create(opts.Context, build, v1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}

func constructObject() (*buildv1alpha1.Build, error) {
	data := Info.NewMetaData()

	// Todo: Make the scope configurable
	strategyKind := buildv1alpha1.ClusterBuildStrategyKind

	buildSpec := buildv1alpha1.BuildSpec{
		Source: buildv1alpha1.GitSource{
			URL: data.Name,
		},
		StrategyRef: &buildv1alpha1.StrategyRef{
			Name: data.Strategy,
			Kind: &strategyKind,
		},
		Output: buildv1alpha1.Image{
			ImageURL: data.RegistryEndpoint,
			SecretRef: &corev1.LocalObjectReference{
				Name: data.SecretName,
			},
		},
	}

	build := &buildv1alpha1.Build{
		ObjectMeta: metav1.ObjectMeta{
			Name:      data.Name,
			Namespace: data.Namespace,
		},
		Spec: buildSpec,
	}

	return build, nil
}
