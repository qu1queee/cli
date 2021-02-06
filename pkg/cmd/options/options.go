package options

import (
	"context"

	"github.com/shipwright-io/cli/pkg/cmd/clients"
)

// List ...
type List struct {
	Context context.Context
	Clients *clients.Clients
}

// NewList ...
func (l List) NewList() (List, error) {

	// Get the clients
	c, err := clients.NewClients()
	if err != nil {
		return List{}, err
	}

	return List{
		Context: context.TODO(),
		Clients: c,
	}, nil
}
