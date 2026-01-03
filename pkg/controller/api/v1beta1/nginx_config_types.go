// +k8s:deepcopy-gen=package
package v1beta1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type NginxConfigSpec struct {
	PageContent string `json:"pageContent"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type NginxConfigStatus struct {
	Ready bool `json:"ready"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type NginxConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NginxConfigSpec   `json:"spec,omitempty"`
	Status NginxConfigStatus `json:"status,omitempty"`
}
