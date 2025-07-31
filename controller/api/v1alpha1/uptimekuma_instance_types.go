// SPDX-License-Identifier: MIT

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// UptimeKumaInstanceSpec defines the desired state of UptimeKumaInstance
type UptimeKumaInstanceSpec struct {
	// URL of the Uptime Kuma instance
	URL string `json:"url"`
	// Username for authentication
	Username string `json:"username,omitempty"`
	// PasswordSecretRef references a Kubernetes Secret for the password
	PasswordSecretRef string `json:"passwordSecretRef,omitempty"`
}

// UptimeKumaInstanceStatus defines the observed state of UptimeKumaInstance
type UptimeKumaInstanceStatus struct {
	// Connected indicates if the controller can reach the instance
	Connected bool `json:"connected,omitempty"`
	// Message provides additional status information
	Message string `json:"message,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// UptimeKumaInstance is the Schema for the uptimekumainstances API
type UptimeKumaInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UptimeKumaInstanceSpec   `json:"spec,omitempty"`
	Status UptimeKumaInstanceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// UptimeKumaInstanceList contains a list of UptimeKumaInstance
type UptimeKumaInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UptimeKumaInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&UptimeKumaInstance{}, &UptimeKumaInstanceList{})
}