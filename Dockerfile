FROM ubuntu:18.04
LABEL maintainers="Woohyung Han"
LABEL description="Dropbox CSI Driver"

RUN apt-get update && apt-get install -q -y python3.6 python3-pip python3-setuptools libfuse-dev
RUN pip3 install --upgrade --no-cache-dir setuptools wheel dbxfs

COPY ./build/dropbox-csi /dropbox-csi
ENTRYPOINT ["/dropbox-csi"]
