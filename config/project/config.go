package project

import "github.com/crossplane-contrib/terrajet/pkg/config"

// Customize the device group with references to other resources
func Customize(p *config.Provider) {
	p.AddResourceConfigurator("metal_project", func(r *config.Resource) {
		r.Group = "project"
	})
}
