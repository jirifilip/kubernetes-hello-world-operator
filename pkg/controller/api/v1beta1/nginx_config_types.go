package v1beta1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type NginxConfigSpec struct {
	PageContent string `json:"pageContent"`
}

type NginxConfigStatus struct {
	Ready bool `json:"ready"`
}

type NginxConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NginxConfigSpec   `json:"spec,omitempty"`
	Status NginxConfigStatus `json:"status,omitempty"`
}
