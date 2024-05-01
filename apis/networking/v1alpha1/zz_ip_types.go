// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type IPInitParameters struct {

	// The IP Address of the resource
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The region that the reserved IP is reserved to.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type IPObservation struct {

	// The ID of Droplet that the reserved IP will be assigned to.
	DropletID *float64 `json:"dropletId,omitempty" tf:"droplet_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The IP Address of the resource
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The region that the reserved IP is reserved to.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The uniform resource name of the reserved ip
	// the uniform resource name for the reserved ip
	Urn *string `json:"urn,omitempty" tf:"urn,omitempty"`
}

type IPParameters struct {

	// The ID of Droplet that the reserved IP will be assigned to.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-digitalocean/apis/compute/v1alpha1.Droplet
	// +kubebuilder:validation:Optional
	DropletID *float64 `json:"dropletId,omitempty" tf:"droplet_id,omitempty"`

	// Reference to a Droplet in compute to populate dropletId.
	// +kubebuilder:validation:Optional
	DropletIDRef *v1.Reference `json:"dropletIdRef,omitempty" tf:"-"`

	// Selector for a Droplet in compute to populate dropletId.
	// +kubebuilder:validation:Optional
	DropletIDSelector *v1.Selector `json:"dropletIdSelector,omitempty" tf:"-"`

	// The IP Address of the resource
	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The region that the reserved IP is reserved to.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// IPSpec defines the desired state of IP
type IPSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IPParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider IPInitParameters `json:"initProvider,omitempty"`
}

// IPStatus defines the observed state of IP.
type IPStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IPObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IP is the Schema for the IPs API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,do}
type IP struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.region) || (has(self.initProvider) && has(self.initProvider.region))",message="spec.forProvider.region is a required parameter"
	Spec   IPSpec   `json:"spec"`
	Status IPStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IPList contains a list of IPs
type IPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IP `json:"items"`
}

// Repository type metadata.
var (
	IP_Kind             = "IP"
	IP_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IP_Kind}.String()
	IP_KindAPIVersion   = IP_Kind + "." + CRDGroupVersion.String()
	IP_GroupVersionKind = CRDGroupVersion.WithKind(IP_Kind)
)

func init() {
	SchemeBuilder.Register(&IP{}, &IPList{})
}
