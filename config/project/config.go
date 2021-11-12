package project

import "github.com/crossplane-contrib/terrajet/pkg/config"

func Customize(p *config.Provider) {
	p.AddResourceConfigurator("metal_project", func(r *config.Resource) {
		r.Group = "project"
	})
}
