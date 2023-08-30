/*
Copyright 2023.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AROClusterSpec defines the desired state of AROCluster
type AROClusterSpec struct {
	// ControlPlaneEndpoint represents the endpoint used to communicate with the control plane.
	// +optional
	ControlPlaneEndpoint string `json:"controlPlaneEndpoint"`
}

// AROClusterStatus defines the observed state of AROCluster
type AROClusterStatus struct {
	// Ready is when the AROControlPlane has a API server URL.
	// +optional
	Ready bool `json:"ready,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// AROCluster is the Schema for the aroclusters API
type AROCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AROClusterSpec   `json:"spec,omitempty"`
	Status AROClusterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AROClusterList contains a list of AROCluster
type AROClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AROCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AROCluster{}, &AROClusterList{})
}
