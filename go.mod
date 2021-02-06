module github.com/shipwright-io/cli

go 1.14

require (
	github.com/mitchellh/go-homedir v1.1.0
	github.com/shipwright-io/build v0.2.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.0
	k8s.io/api v0.18.10
	k8s.io/apimachinery v0.20.2
	k8s.io/client-go v12.0.0+incompatible
)

replace k8s.io/client-go => k8s.io/client-go v0.18.10
