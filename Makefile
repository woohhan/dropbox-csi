build_binary:
	go build -o build/dropbox-csi ./cmd/dropbox

clean:
	go clean -r -x
	rm -rf build/