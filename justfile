set shell := ["powershell.exe", "-c"]

kube-apply:
    kubectl apply -f ./manifests

kube-delete:
    kubectl delete -f ./manifests

codegen:
    deepcopy-gen \
    -v 5 \
    --bounding-dirs github.com/jirifilip/kubernetes-operator-hello-world/pkg/controller/api/v1beta1 \

    
    client-gen -v 5 \
    --input-base github.com/jirifilip/kubernetes-operator-hello-world \
    --input pkg/controller/api/v1beta1 \
    --output-pkg github.com/jirifilip/kubernetes-operator-hello-world/pkg/generated \
    --output-dir pkg/generated

run:
    go run main.go
