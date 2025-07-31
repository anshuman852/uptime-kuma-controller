// SPDX-License-Identifier: MIT

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NotificationChannelSpec defines the desired state of NotificationChannel
type NotificationChannelSpec struct {
	// Type of the notification channel (e.g., telegram, slack, email)
	Type string `json:"type"`
	// Name of the notification channel
	Name string `json:"name"`
	// Data contains channel-specific configuration (e.g., webhook URL, tokens)
	Data map[string]string `json:"data,omitempty"`
	// Reference to the UptimeKumaInstance this channel belongs to
	UptimeKumaInstanceRef string `json:"uptimeKumaInstanceRef"`
}

// NotificationChannelStatus defines the observed state of NotificationChannel
type NotificationChannelStatus struct {
	// ChannelID is the ID assigned by Uptime Kuma (if available)
	ChannelID string `json:"channelID,omitempty"`
	// Status reflects the current state (e.g., Ready, Error)
	Status string `json:"status,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// NotificationChannel is the Schema for the notificationchannels API
type NotificationChannel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NotificationChannelSpec   `json:"spec,omitempty"`
	Status NotificationChannelStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// NotificationChannelList contains a list of NotificationChannel
type NotificationChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NotificationChannel `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NotificationChannel{}, &NotificationChannelList{})
}