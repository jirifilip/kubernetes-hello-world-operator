package main

import (
	"context"
	"fmt"

	"github.com/jirifilip/kubernetes-operator-hello-world/pkg/controller"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/jirifilip/kubernetes-operator-hello-world/pkg/generated/internalclientset"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)

	if err != nil {
		panic(err)
	}

	fmt.Println("Hello world!")

	dynamicClient, err := dynamic.NewForConfig(config)
	controller.Must(err)

	cs, err := internalclientset.NewForConfig(config)
	fmt.Println(cs)

	ctx := context.Background()

	resource := schema.GroupVersionResource{
		Group:    "jirifilip.github.com",
		Version:  "v1beta1",
		Resource: "nginxconfigs",
	}
	result, err := dynamicClient.Resource(resource).Namespace("").List(ctx, metav1.ListOptions{})
	controller.Must(err)

	for _, podObject := range result.Items {
		spec := podObject.Object["spec"].(map[string]interface{})

		fmt.Println(spec["pageContent"])
	}

	controller.WatchNginxConfigs(dynamicClient, ctx)
}
