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

type ApiKeyObservation struct {
}

type ApiKeyParameters struct {

	// Description string for the API key
	// +kubebuilder:validation:Required
	Description *string `json:"description" tf:"description,omitempty"`

	// UUID of project which the new API key is scoped to
	// +kubebuilder:validation:Required
	ProjectID *string `json:"projectId" tf:"project_id,omitempty"`

	// Flag indicating whether the API key shoud be read-only
	// +kubebuilder:validation:Required
	ReadOnly *bool `json:"readOnly" tf:"read_only,omitempty"`
}

// ApiKeySpec defines the desired state of ApiKey
type ApiKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApiKeyParameters `json:"forProvider"`
}

// ApiKeyStatus defines the observed state of ApiKey.
type ApiKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApiKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApiKey is the Schema for the ApiKeys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfequinixmetal}
type ApiKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiKeySpec   `json:"spec"`
	Status            ApiKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiKeyList contains a list of ApiKeys
type ApiKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiKey `json:"items"`
}

// Repository type metadata.
var (
	ApiKeyKind             = "ApiKey"
	ApiKeyGroupKind        = schema.GroupKind{Group: Group, Kind: ApiKeyKind}.String()
	ApiKeyKindAPIVersion   = ApiKeyKind + "." + GroupVersion.String()
	ApiKeyGroupVersionKind = GroupVersion.WithKind(ApiKeyKind)
)

func init() {
	SchemeBuilder.Register(&ApiKey{}, &ApiKeyList{})
}
