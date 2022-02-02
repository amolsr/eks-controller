// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FargateProfileSpec defines the desired state of FargateProfile.
//
// An object representing an Fargate profile.
type FargateProfileSpec struct {
	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientRequestToken *string `json:"clientRequestToken,omitempty"`
	// The name of the Amazon EKS cluster to apply the Fargate profile to.
	ClusterName *string                                  `json:"clusterName,omitempty"`
	ClusterRef  *ackv1alpha1.AWSResourceReferenceWrapper `json:"clusterRef,omitempty"`
	// The name of the Fargate profile.
	// +kubebuilder:validation:Required
	Name *string `json:"name"`
	// The Amazon Resource Name (ARN) of the pod execution role to use for pods
	// that match the selectors in the Fargate profile. The pod execution role allows
	// Fargate infrastructure to register with your cluster as a node, and it provides
	// read access to Amazon ECR image repositories. For more information, see Pod
	// Execution Role (https://docs.aws.amazon.com/eks/latest/userguide/pod-execution-role.html)
	// in the Amazon EKS User Guide.
	// +kubebuilder:validation:Required
	PodExecutionRoleARN *string `json:"podExecutionRoleARN"`
	// The selectors to match for pods to use this Fargate profile. Each selector
	// must have an associated namespace. Optionally, you can also specify labels
	// for a namespace. You may specify up to five selectors in a Fargate profile.
	Selectors []*FargateProfileSelector `json:"selectors,omitempty"`
	// The IDs of subnets to launch your pods into. At this time, pods running on
	// Fargate are not assigned public IP addresses, so only private subnets (with
	// no direct route to an Internet Gateway) are accepted for this parameter.
	Subnets []*string `json:"subnets,omitempty"`
	// The metadata to apply to the Fargate profile to assist with categorization
	// and organization. Each tag consists of a key and an optional value, both
	// of which you define. Fargate profile tags do not propagate to any other resources
	// associated with the Fargate profile, such as the pods that are scheduled
	// with it.
	Tags map[string]*string `json:"tags,omitempty"`
}

// FargateProfileStatus defines the observed state of FargateProfile
type FargateProfileStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// The Unix epoch timestamp in seconds for when the Fargate profile was created.
	// +kubebuilder:validation:Optional
	CreatedAt *metav1.Time `json:"createdAt,omitempty"`
	// The current status of the Fargate profile.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty"`
}

// FargateProfile is the Schema for the FargateProfiles API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type FargateProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FargateProfileSpec   `json:"spec,omitempty"`
	Status            FargateProfileStatus `json:"status,omitempty"`
}

// FargateProfileList contains a list of FargateProfile
// +kubebuilder:object:root=true
type FargateProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FargateProfile `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FargateProfile{}, &FargateProfileList{})
}
