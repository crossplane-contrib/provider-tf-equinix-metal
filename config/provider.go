package config

import (
	tjconfig "github.com/crossplane-contrib/terrajet/pkg/config"
	tf "github.com/equinix/terraform-provider-metal/metal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/crossplane-contrib/provider-tf-equinixmetal/config/connection"
	"github.com/crossplane-contrib/provider-tf-equinixmetal/config/device"
	"github.com/crossplane-contrib/provider-tf-equinixmetal/config/project"
)

const (
	resourcePrefix = "equinixmetal"
	modulePath     = "github.com/crossplane-contrib/provider-tf-equinixmetal"
)

// GetProvider returns provider configuration
func GetProvider() *tjconfig.Provider {
	resourceMap := tf.Provider().ResourcesMap
	// Comment out the line below instead of the above, if your Terraform
	// provider uses an old version (<v2) of github.com/hashicorp/terraform-plugin-sdk.
	// resourceMap := conversion.GetV2ResourceMap(tf.Provider())

	defaultResourceFn := func(name string, terraformResource *schema.Resource) *tjconfig.Resource {
		r := tjconfig.DefaultResource(name, terraformResource)
		// Add any provider-specific defaulting here. For example:
		r.ExternalName = tjconfig.IdentifierFromProvider

		return r
	}

	pc := tjconfig.NewProvider(resourceMap, resourcePrefix, modulePath,
		tjconfig.WithDefaultResourceFn(defaultResourceFn),
		tjconfig.WithIncludeList([]string{
			".*",
		}),
		tjconfig.WithSkipList([]string{
			"volume",
		}),
	)

	for _, configure := range []func(provider *tjconfig.Provider){
		// add custom config functions
		device.Customize,
		project.Customize,
		connection.Customize,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
