FROM quay.io/woohhan/dbxfs:v1.0.0
LABEL maintainers="Woohyung Han"
LABEL description="Dropbox CSI Driver"

COPY ./build/dropbox-csi /dropbox-csi
ENTRYPOINT ["/dropbox-csi"]
