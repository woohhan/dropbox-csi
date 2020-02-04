FROM ubuntu:18.04
LABEL maintainers="Woohyung Han"
LABEL description="Dropbox CSI Driver"

RUN apt-get update && \
    apt-get install -y python3.6 python3-pip python3-setuptools libfuse-dev && \
    pip3 install --upgrade --no-cache-dir setuptools wheel dbxfs && \
    apt-get -y autoremove && \
    apt-get clean
