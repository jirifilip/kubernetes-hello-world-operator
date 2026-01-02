package main

import (
	"context"
	"fmt"

	"github.com/jirifilip/kubernetes-operator-hello-world/pkg/controller"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)

	if err != nil {
		panic(err)
	}

	fmt.Println("Hello world!")

	typedClient, err := kubernetes.NewForConfig(config)
	controller.Must(err)

	ctx := context.Background()

	controller.WatchPods(typedClient, ctx)

}
