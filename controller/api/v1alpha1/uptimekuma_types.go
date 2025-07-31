// SPDX-License-Identifier: MIT

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// UptimeKumaMonitorSpec defines the desired state of UptimeKumaMonitor
type UptimeKumaMonitorSpec struct {
	// Name of the monitor
	Name string `json:"name"`
	// URL to monitor
	URL string `json:"url"`
	// Interval in seconds
	Interval int `json:"interval"`
}

// UptimeKumaMonitorStatus defines the observed state of UptimeKumaMonitor
type UptimeKumaMonitorStatus struct {
	// LastCheck is the timestamp of the last check
	LastCheck metav1.Time `json:"lastCheck,omitempty"`
	// Status is the current status of the monitor (e.g., up, down)
	Status string `json:"status,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// UptimeKumaMonitor is the Schema for the uptimekumamonitors API
type UptimeKumaMonitor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UptimeKumaMonitorSpec   `json:"spec,omitempty"`
	Status UptimeKumaMonitorStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// UptimeKumaMonitorList contains a list of UptimeKumaMonitor
type UptimeKumaMonitorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UptimeKumaMonitor `json:"items"`
}

func init() {
	SchemeBuilder.Register(&UptimeKumaMonitor{}, &UptimeKumaMonitorList{})
}