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

package controller

import (
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/crossplane-runtime/pkg/logging"

	tjconfig "github.com/crossplane-contrib/terrajet/pkg/config"
	"github.com/crossplane-contrib/terrajet/pkg/terraform"

	session "github.com/crossplane-contrib/provider-tf-equinix-metal/internal/controller/bgp/session"
	connection "github.com/crossplane-contrib/provider-tf-equinix-metal/internal/controller/connection/connection"
	device "github.com/crossplane-contrib/provider-tf-equinix-metal/internal/controller/device/device"
	networktype "github.com/crossplane-contrib/provider-tf-equinix-metal/internal/controller/device/networktype"
	attachment "github.com/crossplane-contrib/provider-tf-equinix-metal/internal/controller/ip/attachment"
	gateway "github.com/crossplane-contrib/provider-tf-equinix-metal/internal/controller/metal/gateway"
	organization "github.com/crossplane-contrib/provider-tf-equinix-metal/internal/controller/metal/organization"
	port "github.com/crossplane-contrib/provider-tf-equinix-metal/internal/controller/metal/port"
	vlan "github.com/crossplane-contrib/provider-tf-equinix-metal/internal/controller/metal/vlan"
	vlanattachment "github.com/crossplane-contrib/provider-tf-equinix-metal/internal/controller/port/vlanattachment"
	apikey "github.com/crossplane-contrib/provider-tf-equinix-metal/internal/controller/project/apikey"
	project "github.com/crossplane-contrib/provider-tf-equinix-metal/internal/controller/project/project"
	sshkey "github.com/crossplane-contrib/provider-tf-equinix-metal/internal/controller/project/sshkey"
	providerconfig "github.com/crossplane-contrib/provider-tf-equinix-metal/internal/controller/providerconfig"
	ipblock "github.com/crossplane-contrib/provider-tf-equinix-metal/internal/controller/reserved/ipblock"
	marketrequest "github.com/crossplane-contrib/provider-tf-equinix-metal/internal/controller/spot/marketrequest"
	key "github.com/crossplane-contrib/provider-tf-equinix-metal/internal/controller/ssh/key"
	apikeyuser "github.com/crossplane-contrib/provider-tf-equinix-metal/internal/controller/user/apikey"
	circuit "github.com/crossplane-contrib/provider-tf-equinix-metal/internal/controller/virtual/circuit"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, l logging.Logger, wl workqueue.RateLimiter, ps terraform.SetupFn, ws *terraform.WorkspaceStore, cfg *tjconfig.Provider, concurrency int) error {
	for _, setup := range []func(ctrl.Manager, logging.Logger, workqueue.RateLimiter, terraform.SetupFn, *terraform.WorkspaceStore, *tjconfig.Provider, int) error{
		session.Setup,
		connection.Setup,
		device.Setup,
		networktype.Setup,
		attachment.Setup,
		gateway.Setup,
		organization.Setup,
		port.Setup,
		vlan.Setup,
		vlanattachment.Setup,
		apikey.Setup,
		project.Setup,
		sshkey.Setup,
		providerconfig.Setup,
		ipblock.Setup,
		marketrequest.Setup,
		key.Setup,
		apikeyuser.Setup,
		circuit.Setup,
	} {
		if err := setup(mgr, l, wl, ps, ws, cfg, concurrency); err != nil {
			return err
		}
	}
	return nil
}
