package connection

import "github.com/crossplane-contrib/terrajet/pkg/config"

// Customize the connection group with references to other resources
func Customize(p *config.Provider) {
	p.AddResourceConfigurator("metal_connection", func(r *config.Resource) {
		r.Group = "connection"

		r.References["organization_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tf-equinixmetal/apis/metal/v1alpha1.Organization",
		}

		r.References["project_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tf-equinixmetal/apis/project/v1alpha1.Project",
		}
	})
}
