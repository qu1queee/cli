package clients

import (
	"context"
	"os"
	"path/filepath"

	buildclient "github.com/shipwright-io/build/pkg/client/build/clientset/versioned"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// Clients holds a struct with all the necessary
// interfaces for CRUD events across multiple objects
type Clients struct {
	Context     context.Context
	RestConfig  *restclient.Config
	BuildClient buildclient.Interface
	Client      kubernetes.Interface
}

// NewClients provides you an instance of Clients
func NewClients(ctx context.Context) (*Clients, error) {

	c, err := getRestClient()
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(c)
	if err != nil {
		return nil, err
	}

	buildset, err := buildclient.NewForConfig(c)
	if err != nil {
		return nil, err
	}

	return &Clients{
		Context:     ctx,
		RestConfig:  c,
		BuildClient: buildset,
		Client:      clientset,
	}, nil
}

func getRestClient() (*restclient.Config, error) {
	kubeConfigPath := filepath.Join(
		os.Getenv("HOME"), ".kube", "config",
	)
	c, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		return nil, err
	}
	return c, nil

}
