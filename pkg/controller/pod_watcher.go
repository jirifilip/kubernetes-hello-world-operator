package controller

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func WatchPods(client kubernetes.Interface, ctx context.Context) {
	watcher, err := client.CoreV1().Pods("").Watch(ctx, metav1.ListOptions{})
	Must(err)

	fmt.Println("And now my watch begins...")
	for event := range watcher.ResultChan() {
		fmt.Printf("event type: %s \n", event.Type)

		pod, ok := event.Object.(*corev1.Pod)

		if !ok {
			continue
		}

		fmt.Printf("pod name: %s\n", pod.Name)
	}
}
