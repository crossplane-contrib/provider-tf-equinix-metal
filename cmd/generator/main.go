package main

import (
	"github.com/crossplane-contrib/terrajet/pkg/pipeline"

	"github.com/crossplane-contrib/provider-tf-equinixmetal/config"
)

func main() {
	pipeline.Run(config.GetProvider())
}
