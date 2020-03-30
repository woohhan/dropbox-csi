#!/bin/bash

function remove() {
  kubectl delete -f deploy/pod.yaml --ignore-not-found=true --wait=true
  kubectl delete -f deploy/k8s-1.17/csi-dropbox-plugin.yaml --ignore-not-found=true
  kubectl delete -f deploy/k8s-1.17/csi-dropbox-attacher.yaml --ignore-not-found=true
  kubectl delete -f deploy/k8s-1.17/rbac.yaml --ignore-not-found=true
}
function test() {
  remove

  make image-build

  kubectl create secret generic dropbox-csi --from-literal=token=24XCsHIILjAAAAAAAAAAGL_setHJPfETX0GyhVgvYGL-tRZdfgoc_yRAbiFvaN_m
  kubectl apply -f deploy/k8s-1.17/rbac.yaml
  kubectl apply -f deploy/k8s-1.17/csi-dropbox-attacher.yaml
  sed 's|quay.io/woohhan/dropbox-csi:latest|quay.io/woohhan/dropbox-csi:canary|' deploy/k8s-1.17/csi-dropbox-plugin.yaml | kubectl apply -f -
  kubectl apply -f deploy/pod.yaml

  while [[ $(kubectl get pods dropbox-pod -o 'jsonpath={..status.conditions[?(@.type=="Ready")].status}') != "True" ]]; do echo "waiting for pod" && kubectl get pods && sleep 1; done
  echo "Done!"
  kubectl get pods
  kubectl logs csi-dropboxplugin-0 dropbox-csi
}

case "${1:-}" in
l)
  kubectl logs csi-dropboxplugin-0 dropbox-csi
  ;;
wl)
  watch -n 1 kubectl logs csi-dropboxplugin-0 dropbox-csi
  ;;
t)
  test
  ;;
r)
  remove
  ;;
*)
    echo " $0 [command]
Available Commands:
  l    Show logs
  wl   Watch logs
  t    Test
  r    Remove all
" >&2
    ;;
esac

