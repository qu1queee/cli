module github.com/shipwright-io/cli

go 1.14

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/mitchellh/go-homedir v1.1.0
	github.com/shipwright-io/build v0.2.0
	github.com/spf13/cobra v1.1.1
	k8s.io/apimachinery v0.20.2
	k8s.io/client-go v12.0.0+incompatible
)

replace k8s.io/client-go => k8s.io/client-go v0.18.10
