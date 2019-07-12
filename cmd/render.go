package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"helm.sh/helm/cmd/helm/helm"
)

func main() {
	cmd := helm.newRootCmd(os.Args[1:])
	if err := cmd.Execute(); err != nil {
		switch e := err.(type) {
		case pluginError:
			os.Exit(e.code)
		default:
			os.Exit(1)
		}
	}
}