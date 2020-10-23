package main

import (
	"github.com/Azure/azure-sdk-for-go/tools/generator/cmd"
	"os"
)

func main() {
	if err := cmd.Command().Execute(); err != nil {
		os.Exit(1)
	}
}
