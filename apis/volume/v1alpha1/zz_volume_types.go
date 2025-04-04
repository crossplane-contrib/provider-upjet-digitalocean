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

type VolumeInitParameters struct {

	// A free-form text field up to a limit of 1024 bytes to describe a block storage volume.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Filesystem type (xfs or ext4) for the block storage volume.
	FilesystemType *string `json:"filesystemType,omitempty" tf:"filesystem_type,omitempty"`

	// Initial filesystem label for the block storage volume.
	InitialFilesystemLabel *string `json:"initialFilesystemLabel,omitempty" tf:"initial_filesystem_label,omitempty"`

	// Initial filesystem type (xfs or ext4) for the block storage volume.
	InitialFilesystemType *string `json:"initialFilesystemType,omitempty" tf:"initial_filesystem_type,omitempty"`

	// A name for the block storage volume. Must be lowercase and be composed only of numbers, letters and "-", up to a limit of 64 characters. The name must begin with a letter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region that the block storage volume will be created in.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The size of the block storage volume in GiB. If updated, can only be expanded.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The ID of an existing volume snapshot from which the new volume will be created. If supplied, the region and size will be limited on creation to that of the referenced snapshot
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-digitalocean/apis/volume/v1alpha1.Snapshot
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// Reference to a Snapshot in volume to populate snapshotId.
	// +kubebuilder:validation:Optional
	SnapshotIDRef *v1.Reference `json:"snapshotIdRef,omitempty" tf:"-"`

	// Selector for a Snapshot in volume to populate snapshotId.
	// +kubebuilder:validation:Optional
	SnapshotIDSelector *v1.Selector `json:"snapshotIdSelector,omitempty" tf:"-"`

	// A list of the tags to be applied to this Volume.
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type VolumeObservation struct {

	// A free-form text field up to a limit of 1024 bytes to describe a block storage volume.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A list of associated droplet ids.
	// +listType=set
	DropletIds []*float64 `json:"dropletIds,omitempty" tf:"droplet_ids,omitempty"`

	// Filesystem label for the block storage volume.
	FilesystemLabel *string `json:"filesystemLabel,omitempty" tf:"filesystem_label,omitempty"`

	// Filesystem type (xfs or ext4) for the block storage volume.
	FilesystemType *string `json:"filesystemType,omitempty" tf:"filesystem_type,omitempty"`

	// The unique identifier for the volume.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Initial filesystem label for the block storage volume.
	InitialFilesystemLabel *string `json:"initialFilesystemLabel,omitempty" tf:"initial_filesystem_label,omitempty"`

	// Initial filesystem type (xfs or ext4) for the block storage volume.
	InitialFilesystemType *string `json:"initialFilesystemType,omitempty" tf:"initial_filesystem_type,omitempty"`

	// A name for the block storage volume. Must be lowercase and be composed only of numbers, letters and "-", up to a limit of 64 characters. The name must begin with a letter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region that the block storage volume will be created in.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The size of the block storage volume in GiB. If updated, can only be expanded.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The ID of an existing volume snapshot from which the new volume will be created. If supplied, the region and size will be limited on creation to that of the referenced snapshot
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// A list of the tags to be applied to this Volume.
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The uniform resource name for the volume.
	// the uniform resource name for the volume.
	Urn *string `json:"urn,omitempty" tf:"urn,omitempty"`
}

type VolumeParameters struct {

	// A free-form text field up to a limit of 1024 bytes to describe a block storage volume.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Filesystem type (xfs or ext4) for the block storage volume.
	// +kubebuilder:validation:Optional
	FilesystemType *string `json:"filesystemType,omitempty" tf:"filesystem_type,omitempty"`

	// Initial filesystem label for the block storage volume.
	// +kubebuilder:validation:Optional
	InitialFilesystemLabel *string `json:"initialFilesystemLabel,omitempty" tf:"initial_filesystem_label,omitempty"`

	// Initial filesystem type (xfs or ext4) for the block storage volume.
	// +kubebuilder:validation:Optional
	InitialFilesystemType *string `json:"initialFilesystemType,omitempty" tf:"initial_filesystem_type,omitempty"`

	// A name for the block storage volume. Must be lowercase and be composed only of numbers, letters and "-", up to a limit of 64 characters. The name must begin with a letter.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region that the block storage volume will be created in.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The size of the block storage volume in GiB. If updated, can only be expanded.
	// +kubebuilder:validation:Optional
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The ID of an existing volume snapshot from which the new volume will be created. If supplied, the region and size will be limited on creation to that of the referenced snapshot
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-digitalocean/apis/volume/v1alpha1.Snapshot
	// +kubebuilder:validation:Optional
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// Reference to a Snapshot in volume to populate snapshotId.
	// +kubebuilder:validation:Optional
	SnapshotIDRef *v1.Reference `json:"snapshotIdRef,omitempty" tf:"-"`

	// Selector for a Snapshot in volume to populate snapshotId.
	// +kubebuilder:validation:Optional
	SnapshotIDSelector *v1.Selector `json:"snapshotIdSelector,omitempty" tf:"-"`

	// A list of the tags to be applied to this Volume.
	// +kubebuilder:validation:Optional
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// VolumeSpec defines the desired state of Volume
type VolumeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VolumeParameters `json:"forProvider"`
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
	InitProvider VolumeInitParameters `json:"initProvider,omitempty"`
}

// VolumeStatus defines the observed state of Volume.
type VolumeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VolumeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Volume is the Schema for the Volumes API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,do}
type Volume struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.region) || (has(self.initProvider) && has(self.initProvider.region))",message="spec.forProvider.region is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.size) || (has(self.initProvider) && has(self.initProvider.size))",message="spec.forProvider.size is a required parameter"
	Spec   VolumeSpec   `json:"spec"`
	Status VolumeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VolumeList contains a list of Volumes
type VolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Volume `json:"items"`
}

// Repository type metadata.
var (
	Volume_Kind             = "Volume"
	Volume_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Volume_Kind}.String()
	Volume_KindAPIVersion   = Volume_Kind + "." + CRDGroupVersion.String()
	Volume_GroupVersionKind = CRDGroupVersion.WithKind(Volume_Kind)
)

func init() {
	SchemeBuilder.Register(&Volume{}, &VolumeList{})
}
