/*
Copyright 2021.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NetworkSpec defines the desired state of Network
type NetworkSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Disable ingress traffic inside network
	// +optional
	DisableInsideIngressTraffic bool `json:"disableInsideIngressTraffic"`

	// Disable egress traffic inside network
	// +optional
	DisableInsideEgressTraffic bool `json:"disableInsideEgressTraffic"`

	// NetworkIngressPorts, specifies ports to which this network can receive connection
	// +optional
	NetworkIngressPorts []Ports `json:"networkIngressPorts"`

	// NetworkEgressPorts, specifies ports from which this network can create connection
	// +optional
	NetworkEgressPorts []Ports `json:"networkEgressPorts"`
}

// NetworkStatus defines the observed state of Network
type NetworkStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Cluster

// Network is the Schema for the networks API
type Network struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkSpec   `json:"spec,omitempty"`
	Status NetworkStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// NetworkList contains a list of Network
type NetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Network `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Network{}, &NetworkList{})
}

func (in Network) NetworkPolicyNameIsolation() string {
	return in.Name + "-network-policy-isolation"
}

func (in Network) NetworkPolicyNameInternet() string {
	return in.Name + "-network-policy-internet"
}

func (in Network) NetworkPolicyNameConnection() string {
	return in.Name + "-network-policy-con"
}
