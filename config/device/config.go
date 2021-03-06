/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package device

import "github.com/crossplane-contrib/terrajet/pkg/config"

// Customize the device group with references to other resources
func Customize(p *config.Provider) {
	p.AddResourceConfigurator("metal_device", func(r *config.Resource) {
		r.ShortGroup = "device"
		r.References["project_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tf-equinix-metal/apis/project/v1alpha1.Project",
		}
	})
}
