package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)

	if err != nil {
		panic(err)
	}

	fmt.Println(config)

	fmt.Println("Hello world!")

	client, err := kubernetes.NewForConfig(config)

	ctx := context.Background()

	listedPods, err := client.CoreV1().Pods("kube-system").List(ctx, metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println(listedPods)
}
