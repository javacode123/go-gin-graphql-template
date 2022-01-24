package main

import (
	"os"

	"github.com/java_code123/go-gin-graphql-template/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
