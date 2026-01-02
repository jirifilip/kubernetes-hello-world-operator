package main

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)

	if err != nil {
		panic(err)
	}

	fmt.Println("Hello world!")

	client, err := kubernetes.NewForConfig(config)
	must(err)

	ctx := context.Background()

	listedPods, err := client.CoreV1().Pods("kube-system").List(ctx, metav1.ListOptions{})
	must(err)

	watcher, err := client.CoreV1().Pods("").Watch(ctx, metav1.ListOptions{})
	must(err)

	fmt.Println("And now my watch begins...")
	for event := range watcher.ResultChan() {
		fmt.Printf("event type: %s \n", event.Type)

		pod, ok := event.Object.(*corev1.Pod)

		if !ok {
			continue
		}

		fmt.Printf("pod name: %s\n", pod.Name)
	}

	fmt.Println(listedPods)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
