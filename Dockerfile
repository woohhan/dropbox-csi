#FROM python:3.8-alpine3.11
FROM ubuntu:18.04
LABEL maintainers="Woohyung Han"
LABEL description="Dropbox CSI Driver"

# RUN apk add --no-cache gcc musl-dev libffi-dev openssl-dev
RUN apt-get update && apt-get install -q -y python3.6 python3-pip python3-setuptools libfuse-dev
RUN pip3 install --upgrade --no-cache-dir setuptools wheel dbxfs
COPY ./build/dropbox-csi /dropbox-csi
COPY ./dbxfs /usr/local/bin/dbxfs
ENTRYPOINT ["/dropbox-csi"]
