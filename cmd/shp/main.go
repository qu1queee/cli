package main

import (
	"fmt"
	"os"

	"github.com/shipwright-io/cli/pkg/cmd"
)

func main() {
	rootCmd := cmd.NewSHPRootCmd()
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("[ERROR] %#v\n", err)
		os.Exit(1)
	}
}
