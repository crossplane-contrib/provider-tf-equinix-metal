package device

import "github.com/crossplane-contrib/terrajet/pkg/config"

// Customize the device group with references to other resources
func Customize(p *config.Provider) {
	p.AddResourceConfigurator("metal_device", func(r *config.Resource) {
		r.Group = "device"
		r.References["project_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tf-equinix-metal/apis/project/v1alpha1.Project",
		}
	})
}
