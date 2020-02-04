.DEFAULT_GOAL := help

build:
	make clean
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o ./build/dropbox-csi ./cmd/dropbox
image-build:
	make clean
	make build
	docker build -t dropbox-csi:canary -f Dockerfile .
clean:
	go clean ./...
	rm -rf build/
test-cluster-up:
	CHANGE_MINIKUBE_NONE_USER=true sudo minikube start --vm-driver=none --extra-config=kubeadm.ignore-preflight-errors=SystemVerification --extra-config=kubelet.resolv-conf=/run/systemd/resolve/resolv.conf
	sudo chown -R runner:runner /home/runner/
	sleep 10
test-cluster-clean:
	sudo minikube delete
yaml-deploy:
	kubectl create -f deploy/k8s-1.16/rbac.yaml
	kubectl create -f deploy/k8s-1.16/csi-dropbox-plugin.yaml
	kubectl create -f deploy/k8s-1.16/csi-dropbox-attacher.yaml
	kubectl create -f deploy/pod.yaml
yaml-clean:
	kubectl delete --ignore-not-found=true -f deploy/pod.yaml
	kubectl delete --ignore-not-found=true -f deploy/k8s-1.16/csi-dropbox-attacher.yaml
	kubectl delete --ignore-not-found=true -f deploy/k8s-1.16/csi-dropbox-plugin.yaml
	kubectl delete --ignore-not-found=true -f deploy/k8s-1.16/rbac.yaml
help:
	@echo "Usage: make [Target ...]"
	@echo "  build"
	@echo "  image-build"
	@echo "  clean"
	@echo "  test-cluster-up"
	@echo "  test-cluster-clean"
	@echo "  yaml-deploy"
	@echo "  yaml-clean"
