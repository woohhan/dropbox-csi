.DEFAULT_GOAL := help

.PHONY: build image-build clean

build:
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o ./build/dropbox-csi ./cmd/dropbox
image-build:
	make build
	docker build -t quay.io/woohhan/dropbox-csi:canary .
clean:
	go clean ./...
	rm -rf build/
test-cluster-up:
	CHANGE_MINIKUBE_NONE_USER=true sudo minikube start --vm-driver=none --extra-config=kubeadm.ignore-preflight-errors=SystemVerification --extra-config=kubelet.resolv-conf=/run/systemd/resolve/resolv.conf --kubernetes-version=v1.17.4
	sudo chown -R runner:runner /home/runner/
	sleep 10
test-cluster-clean:
	sudo minikube delete
yaml-deploy:
	kubectl create -f deploy/k8s-1.17/rbac.yaml
	cp deploy/k8s-1.17/csi-dropbox-plugin.yaml /tmp/csi-dropbox-plugin.yaml
	sed -i 's\quay.io/woohhan/dropbox-csi:latest\quay.io/woohhan/dropbox-csi:canary\' /tmp/csi-dropbox-plugin.yaml
	kubectl create -f /tmp/csi-dropbox-plugin.yaml
	kubectl create -f deploy/k8s-1.17/csi-dropbox-attacher.yaml
	kubectl create -f deploy/pod.yaml
	sleep 5
	kubectl wait --for=condition=Ready pod/dropbox-pod --timeout 90s
yaml-clean:
	kubectl delete --ignore-not-found=true -f deploy/pod.yaml
	kubectl delete --ignore-not-found=true -f deploy/k8s-1.17/csi-dropbox-attacher.yaml
	kubectl delete --ignore-not-found=true -f deploy/k8s-1.17/csi-dropbox-plugin.yaml
	kubectl delete --ignore-not-found=true -f deploy/k8s-1.17/rbac.yaml
test:
	kubectl exec -it dropbox-pod -- ls /var/www/html
	kubectl exec -it dropbox-pod -- ls /var/www/html | grep e2e_test_file2 > /dev/null
log:
	kubectl logs csi-dropboxplugin-0 dropbox-csi
lt:
	make yaml-clean
	make image-build
	make yaml-deploy
	make test
help:
	@echo "Usage: make [Target ...]"
	@echo "  build"
	@echo "  image-build"
	@echo "  clean"
	@echo "  test-cluster-up"
	@echo "  test-cluster-clean"
	@echo "  yaml-deploy"
	@echo "  yaml-clean"
	@echo "  test"
	@echo "  lt                     Run local test"
