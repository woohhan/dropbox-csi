#!/bin/bash

function test() {
  make image-build

  kubectl delete -f deploy/pod.yaml --ignore-not-found=true
  kubectl delete -f deploy/k8s-1.17/csi-dropbox-plugin.yaml --ignore-not-found=true
  kubectl delete -f deploy/k8s-1.17/csi-dropbox-attacher.yaml --ignore-not-found=true
  kubectl delete -f deploy/k8s-1.17/rbac.yaml --ignore-not-found=true

  kubectl apply -f deploy/k8s-1.17/rbac.yaml
  kubectl apply -f deploy/k8s-1.17/csi-dropbox-attacher.yaml
  kubectl apply -f deploy/k8s-1.17/csi-dropbox-plugin.yaml
  kubectl apply -f deploy/pod.yaml

  kubectl get pods
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
*)
    echo " $0 [command]
Available Commands:
  l    Show logs
  wl   Watch logs
  t    Test
" >&2
    ;;
esac

