package clients

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// Clients holds a struct with all the necessary
// interfaces for CRUD events across multiple objects
type Clients struct {
	RestConfig *rest.Config
	Client     kubernetes.Interface
}

// NewClients provides you an instance of Clients
func NewClients() (*Clients, error) {
	// TODO: properly populate each field
	return &Clients{
		RestConfig: nil,
		Client:     nil,
	}, nil
}
