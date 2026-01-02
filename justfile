set shell := ["powershell.exe", "-c"]

kube-apply:
    kubectl apply -f manifests

kube-delete:
    kubectl delete -f manifests

run:
    go run main.go