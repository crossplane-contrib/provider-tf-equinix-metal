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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ConnectionObservation struct {
	Ports []PortsObservation `json:"ports,omitempty" tf:"ports,omitempty"`

	Speed *int64 `json:"speed,omitempty" tf:"speed,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	Token *string `json:"token,omitempty" tf:"token,omitempty"`
}

type ConnectionParameters struct {

	// Description of the connection resource
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Facility where the connection will be created
	// +kubebuilder:validation:Optional
	Facility *string `json:"facility,omitempty" tf:"facility,omitempty"`

	// Metro where to the connection will be created
	// +kubebuilder:validation:Optional
	Metro *string `json:"metro,omitempty" tf:"metro,omitempty"`

	// Mode for connections in IBX facilities with the dedicated type - standard or tunnel
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// Name of the connection resource
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// ID of the organization responsible for the connection
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tf-equinix-metal/apis/metal/v1alpha1.Organization
	// +kubebuilder:validation:Optional
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// +kubebuilder:validation:Optional
	OrganizationIDRef *v1.Reference `json:"organizationIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	OrganizationIDSelector *v1.Selector `json:"organizationIdSelector,omitempty" tf:"-"`

	// ID of the project where the connection is scoped to, only used for type == "shared"
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tf-equinix-metal/apis/project/v1alpha1.Project
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// Connection redundancy - redundant or primary
	// +kubebuilder:validation:Required
	Redundancy *string `json:"redundancy" tf:"redundancy,omitempty"`

	// Tags attached to the connection
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Connection type - dedicated or shared
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type PortsObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	LinkStatus *string `json:"linkStatus,omitempty" tf:"link_status,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	Speed *int64 `json:"speed,omitempty" tf:"speed,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	VirtualCircuitIds []string `json:"virtualCircuitIds,omitempty" tf:"virtual_circuit_ids,omitempty"`
}

type PortsParameters struct {
}

// ConnectionSpec defines the desired state of Connection
type ConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConnectionParameters `json:"forProvider"`
}

// ConnectionStatus defines the observed state of Connection.
type ConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Connection is the Schema for the Connections API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,equinixmetaljet}
type Connection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConnectionSpec   `json:"spec"`
	Status            ConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectionList contains a list of Connections
type ConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Connection `json:"items"`
}

// Repository type metadata.
var (
	Connection_Kind             = "Connection"
	Connection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Connection_Kind}.String()
	Connection_KindAPIVersion   = Connection_Kind + "." + CRDGroupVersion.String()
	Connection_GroupVersionKind = CRDGroupVersion.WithKind(Connection_Kind)
)

func init() {
	SchemeBuilder.Register(&Connection{}, &ConnectionList{})
}
