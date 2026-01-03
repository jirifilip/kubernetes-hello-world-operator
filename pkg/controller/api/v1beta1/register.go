package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	SchemeBuilder      = runtime.NewSchemeBuilder(addKnownTypes)
	AddToSheme         = SchemeBuilder.AddToScheme
	SchemeGroupVersion = schema.GroupVersion{
		Group:   "jirifilip.github.com",
		Version: "v1beta1",
	}
)

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&NginxConfig{},
	)
	return nil
}
