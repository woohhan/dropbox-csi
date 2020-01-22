case "${1:-}" in
make-cp)
  make
  scp -i $(minikube ssh-key) ./build/dropbox-csi docker@$(minikube ip):/home/docker
  ;;
deploy)
  kubectl delete --ignore-not-found=true -f deploy/k8s-1.16/csi-dropbox-plugin.yaml
  kubectl create -f deploy/k8s-1.16/csi-dropbox-plugin.yaml
  ;;
  *)
  echo "usage:" >&2
  echo "  $0 make-cp" >&2
  echo "  $0 deploy" >&2
  echo "  $0 help" >&2
  ;;
esac

