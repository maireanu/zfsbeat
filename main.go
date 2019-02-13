package main

import (
	"os"

	"github.com/maireanu/zfsbeat/cmd"

	_ "github.com/maireanu/zfsbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
