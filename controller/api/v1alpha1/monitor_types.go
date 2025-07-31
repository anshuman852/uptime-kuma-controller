// SPDX-License-Identifier: MIT

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MonitorSpec defines the desired state of Monitor
type MonitorSpec struct {
	// URL to monitor
	URL string `json:"url"`
	// Type of monitor (e.g., http, tcp)
	Type string `json:"type"`
	// Interval in seconds
	Interval int `json:"interval"`
	// Reference to UptimeKumaInstance
	UptimeKumaInstanceRef string `json:"uptimeKumaInstanceRef"`
}

// MonitorStatus defines the observed state of Monitor
type MonitorStatus struct {
	// MonitorID in Uptime Kuma (if provisioned)
	MonitorID *string `json:"monitorID,omitempty"`
	// Status of the monitor (e.g., Pending, Ready, Error)
	Status string `json:"status,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Monitor is the Schema for the monitors API
type Monitor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MonitorSpec   `json:"spec,omitempty"`
	Status MonitorStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MonitorList contains a list of Monitor
type MonitorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Monitor `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Monitor{}, &MonitorList{})
}